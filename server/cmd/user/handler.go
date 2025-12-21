package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"zpi/server/shared/dal/sqlentity"
	"zpi/server/shared/dal/sqlfunc"
	"zpi/server/shared/errno"
	base "zpi/server/shared/kitex_gen/base"
	user "zpi/server/shared/kitex_gen/user"

	"github.com/hertz-contrib/paseto"
	"gorm.io/gorm"
)

// VerifyCodeManager 验证码管理器接口
type VerifyCodeManager interface {
	GenerateCode() (string, error)
	StoreCode(ctx context.Context, codeType, target, purpose, code string) error
	VerifyCode(ctx context.Context, codeType, target, purpose, code string) (bool, error)
	CheckRateLimit(ctx context.Context, target string) (bool, error)
	IncrementRateLimit(ctx context.Context, target string) error
}

// EmailSender 邮件发送器接口
type EmailSender interface {
	SendVerifyCode(to, code, purpose string) error
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	EncryptManager
	TokenGenerator
	VerifyCodeManager
	EmailSender
	*UserManager
}

type UserManager struct {
	Query *sqlfunc.Query
}

type EncryptManager interface {
	EncryptPassword(password string) string
}

type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

// HealthCheck implements the UserServiceImpl interface.
func (s *UserServiceImpl) HealthCheck(ctx context.Context) (resp *base.HealthCheckResponse, err error) {
	resp = &base.HealthCheckResponse{
		Status:  "ok",
		Version: "1.0.0",
		Uptime:  0,
		Details: map[string]string{
			"service": "user",
			"status":  "healthy",
		},
	}
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = &user.RegisterResponse{
		BaseResp: &base.BaseResponse{}}

	// 参数验证
	if req.Email == nil && req.Phone == nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "邮箱或手机号必须提供一个"))
		return
	}

	if req.Password == "" || req.Nickname == "" || req.VerifyCode == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "密码、昵称和验证码不能为空"))
		return
	}

	// 验证验证码
	var target string
	var codeType string
	if req.Email != nil && *req.Email != "" {
		target = *req.Email
		codeType = "email"
	} else {
		target = *req.Phone
		codeType = "phone"
	}

	valid, err := s.VerifyCode(ctx, codeType, target, "1", req.VerifyCode) // purpose=1 表示注册
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "验证码验证失败"))
		return
	}
	if !valid {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_VerifyCodeError, "验证码错误或已过期"))
		return
	}

	u := s.Query.User

	// 检查邮箱是否已存在
	if req.Email != nil && *req.Email != "" {
		existUser, _ := u.WithContext(ctx).Where(u.Email.Eq(*req.Email)).First()
		if existUser != nil {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_EmailAlreadyExist, "邮箱已被注册"))
			return
		}
	}

	// 检查手机号是否已存在
	if req.Phone != nil && *req.Phone != "" {
		existUser, _ := u.WithContext(ctx).Where(u.Phone.Eq(*req.Phone)).First()
		if existUser != nil {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_PhoneAlreadyExist, "手机号已被注册"))
			return
		}
	}

	// 检查用户名是否已存在
	if req.Username != nil && *req.Username != "" {
		existUser, _ := u.WithContext(ctx).Where(u.Username.Eq(*req.Username)).First()
		if existUser != nil {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UsernameAlreadyExist, "用户名已被使用"))
			return
		}
	}

	// 加密密码
	encryptedPassword := s.EncryptPassword(req.Password)

	// 创建用户
	newUser := &sqlentity.User{
		Username: func() string {
			if req.Username != nil {
				return *req.Username
			}
			// 如果没有提供用户名，使用邮箱或手机号作为用户名
			if req.Email != nil {
				return *req.Email
			}
			return *req.Phone
		}(),
		Email: func() string {
			if req.Email != nil {
				return *req.Email
			}
			return ""
		}(),
		Phone: func() *string {
			return req.Phone
		}(),
		Password: encryptedPassword,
		Nickname: &req.Nickname,
		Status:   1, // 默认激活状态
		Role:     "user",
	}

	if err = u.WithContext(ctx).Create(newUser); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "创建用户失败"))
		return
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.UserId = int64(newUser.ID)
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = &user.LoginResponse{
		BaseResp: &base.BaseResponse{},
	}

	// 参数验证
	if req.Account == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "账号不能为空"))
		return
	}

	if req.Password == nil && req.VerifyCode == nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "密码或验证码必须提供一个"))
		return
	}

	u := s.Query.User

	// 查找用户（支持邮箱/用户名/手机号登录）
	var foundUser *sqlentity.User
	foundUser, err = u.WithContext(ctx).Where(
		u.Email.Eq(req.Account),
	).Or(
		u.Username.Eq(req.Account),
	).Or(
		u.Phone.Eq(req.Account),
	).First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserNotFound, "用户不存在"))
		} else {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "查询用户失败"))
		}
		return
	}

	// 检查用户状态
	if foundUser.Status != 1 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_AuthorizeFail, "账号已被禁用"))
		return
	}

	// 验证密码或验证码
	if req.Password != nil && *req.Password != "" {
		// 密码登录
		encryptedPassword := s.EncryptPassword(*req.Password)
		if encryptedPassword != foundUser.Password {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_PasswordError, "密码错误"))
			return
		}
	} else if req.VerifyCode != nil && *req.VerifyCode != "" {
		// 验证码登录
		var target string
		var codeType string

		// 判断账号类型
		if foundUser.Email != "" && foundUser.Email == req.Account {
			target = foundUser.Email
			codeType = "email"
		} else if foundUser.Phone != nil && *foundUser.Phone == req.Account {
			target = *foundUser.Phone
			codeType = "phone"
		} else {
			// 使用用户名登录时，优先使用邮箱
			if foundUser.Email != "" {
				target = foundUser.Email
				codeType = "email"
			} else if foundUser.Phone != nil {
				target = *foundUser.Phone
				codeType = "phone"
			} else {
				resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "该账号未绑定邮箱或手机号"))
				return
			}
		}

		valid, verifyErr := s.VerifyCode(ctx, codeType, target, "2", *req.VerifyCode) // purpose=2 表示登录
		if verifyErr != nil {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "验证码验证失败"))
			return
		}
		if !valid {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_VerifyCodeError, "验证码错误或已过期"))
			return
		}
	}

	// 更新最后登录时间
	now := time.Now()
	_, _ = u.WithContext(ctx).Where(u.ID.Eq(foundUser.ID)).Update(u.LastLoginAt, &now)

	// 生成 Token
	claims := &paseto.StandardClaims{
		ID:        string(rune(foundUser.ID)),
		Subject:   foundUser.Username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour), // 24小时过期
	}

	token, err := s.CreateToken(claims)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "生成Token失败"))
		return
	}

	// 构造返回的用户信息
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Token = token
	resp.UserEntity = &base.UserEntity{
		Id: int64(foundUser.ID),
		User: &base.User{
			Email:    foundUser.Email,
			Nickname: *foundUser.Nickname,
			AvatarUrl: func() string {
				if foundUser.Avatar != nil {
					return *foundUser.Avatar
				}
				return ""
			}(),
			CreatedAt: foundUser.CreatedAt.Unix(),
			Username:  &foundUser.Username,
			Phone:     foundUser.Phone,
		},
	}

	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	resp = &user.GetUserResponse{
		BaseResp: &base.BaseResponse{},
	}

	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID无效"))
		return
	}

	u := s.Query.User
	foundUser, err := u.WithContext(ctx).Where(u.ID.Eq(uint64(req.UserId))).First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserNotFound, "用户不存在"))
		} else {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "查询用户失败"))
		}
		return
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.UserEntity = &base.UserEntity{
		Id: int64(foundUser.ID),
		User: &base.User{
			Email:    foundUser.Email,
			Nickname: *foundUser.Nickname,
			AvatarUrl: func() string {
				if foundUser.Avatar != nil {
					return *foundUser.Avatar
				}
				return ""
			}(),
			CreatedAt: foundUser.CreatedAt.Unix(),
			Username:  &foundUser.Username,
			Phone:     foundUser.Phone,
		},
	}

	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	resp = &user.UpdateUserResponse{
		BaseResp: &base.BaseResponse{},
	}

	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID无效"))
		return
	}

	u := s.Query.User

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Nickname != nil && *req.Nickname != "" {
		updates["nickname"] = *req.Nickname
	}
	if req.AvatarUrl != nil && *req.AvatarUrl != "" {
		updates["avatar"] = *req.AvatarUrl
	}

	if len(updates) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "没有需要更新的字段"))
		return
	}

	_, err = u.WithContext(ctx).Where(u.ID.Eq(uint64(req.UserId))).Updates(updates)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "更新用户信息失败"))
		return
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return
}

// ChangePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangePassword(ctx context.Context, req *user.ChangePasswordRequest) (resp *user.ChangePasswordResponse, err error) {
	resp = &user.ChangePasswordResponse{
		BaseResp: &base.BaseResponse{},
	}

	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID无效"))
		return
	}

	if req.OldPassword == "" || req.NewPassword_ == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "旧密码和新密码不能为空"))
		return
	}

	u := s.Query.User

	// 查找用户
	foundUser, err := u.WithContext(ctx).Where(u.ID.Eq(uint64(req.UserId))).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserNotFound, "用户不存在"))
		} else {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "查询用户失败"))
		}
		return
	}

	// 验证旧密码
	encryptedOldPassword := s.EncryptPassword(req.OldPassword)
	if encryptedOldPassword != foundUser.Password {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_PasswordError, "旧密码错误"))
		return
	}

	// 更新密码
	encryptedNewPassword := s.EncryptPassword(req.NewPassword_)
	_, err = u.WithContext(ctx).Where(u.ID.Eq(uint64(req.UserId))).Update(u.Password, encryptedNewPassword)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "修改密码失败"))
		return
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return
}

// ResetPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest) (resp *user.ResetPasswordResponse, err error) {
	resp = &user.ResetPasswordResponse{
		BaseResp: &base.BaseResponse{},
	}

	// 参数验证
	if req.Account == "" || req.VerifyCode == "" || req.NewPassword_ == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "账号、验证码和新密码不能为空"))
		return
	}

	u := s.Query.User

	// 查找用户
	var foundUser *sqlentity.User
	foundUser, err = u.WithContext(ctx).Where(
		u.Email.Eq(req.Account),
	).Or(
		u.Phone.Eq(req.Account),
	).First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserNotFound, "用户不存在"))
		} else {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "查询用户失败"))
		}
		return
	}

	// 验证验证码
	var codeType string
	if foundUser.Email == req.Account {
		codeType = "email"
	} else {
		codeType = "phone"
	}

	valid, err := s.VerifyCode(ctx, codeType, req.Account, "3", req.VerifyCode) // purpose=3 表示重置密码
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "验证码验证失败"))
		return
	}
	if !valid {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_VerifyCodeError, "验证码错误或已过期"))
		return
	}

	// 更新密码
	encryptedPassword := s.EncryptPassword(req.NewPassword_)
	_, err = u.WithContext(ctx).Where(u.ID.Eq(foundUser.ID)).Update(u.Password, encryptedPassword)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "重置密码失败"))
		return
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return
}

// SendVerifyCode implements the UserServiceImpl interface.
func (s *UserServiceImpl) SendVerifyCode(ctx context.Context, req *user.SendVerifyCodeRequest) (resp *user.SendVerifyCodeResponse, err error) {
	resp = &user.SendVerifyCodeResponse{
		BaseResp: &base.BaseResponse{},
	}

	// 参数验证
	if req.Target == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "目标地址不能为空"))
		return
	}

	// 检查发送频率限制
	allowed, err := s.CheckRateLimit(ctx, req.Target)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "检查发送频率失败"))
		return
	}
	if !allowed {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_TooManyVerifyCodeRequest, "发送过于频繁，请1小时后再试"))
		return
	}

	// 生成验证码
	code, err := s.GenerateCode()
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "生成验证码失败"))
		return
	}

	// 确定验证码类型
	var codeType string
	if req.CodeType == base.VerifyCodeType_EMAIL {
		codeType = "email"
	} else {
		codeType = "phone"
	}

	// 存储验证码到 Redis
	purposeStr := fmt.Sprintf("%d", req.Purpose)
	if err = s.StoreCode(ctx, codeType, req.Target, purposeStr, code); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "存储验证码失败"))
		return
	}

	// 发送验证码
	if codeType == "email" {
		if err = s.EmailSender.SendVerifyCode(req.Target, code, purposeStr); err != nil {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "发送邮件失败"))
			return
		}
	} else {
		// TODO: 实现短信发送
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ServiceErr, "短信发送功能暂未实现"))
		return
	}

	// 增加发送计数
	if err = s.IncrementRateLimit(ctx, req.Target); err != nil {
		// 这里失败不影响主流程，只记录日志
		fmt.Printf("Failed to increment rate limit: %v\n", err)
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return
}

// UploadResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadResume(ctx context.Context, req *user.UploadResumeRequest) (resp *user.UploadResumeResponse, err error) {
	resp = &user.UploadResumeResponse{
		BaseResp: &base.BaseResponse{},
	}

	// 参数验证
	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID无效"))
		return
	}

	if req.FileUrl == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "文件URL不能为空"))
		return
	}

	r := s.Query.Resume

	// 检查用户是否已有简历，如果有则更新
	existingResume, _ := r.WithContext(ctx).Where(r.UserID.Eq(uint64(req.UserId))).First()

	if existingResume != nil {
		// 更新现有简历
		updates := map[string]interface{}{
			"file_url":       req.FileUrl,
			"parsed_content": req.ParsedContent,
			"parse_status":   2, // 2-解析成功
		}

		_, err = r.WithContext(ctx).Where(r.ID.Eq(existingResume.ID)).Updates(updates)
		if err != nil {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "更新简历失败"))
			return
		}
		resp.BaseResp = errno.BuildBaseResp(errno.Success)
		resp.ResumeId = int64(existingResume.ID)
		return
	}

	// 创建新简历
	newResume := &sqlentity.Resume{
		UserID:        uint64(req.UserId),
		FileURL:       req.FileUrl,
		ParsedContent: &req.ParsedContent,
		ParseStatus:   2, // 2-解析成功
	}

	if err = r.WithContext(ctx).Create(newResume); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "创建简历失败"))
		return
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.ResumeId = int64(newResume.ID)
	return
}

// GetResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetResume(ctx context.Context, req *user.GetResumeRequest) (resp *user.GetResumeResponse, err error) {
	resp = &user.GetResumeResponse{
		BaseResp: &base.BaseResponse{},
	}

	// 参数验证
	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID无效"))
		return
	}

	r := s.Query.Resume

	// 查找用户的简历
	foundResume, err := r.WithContext(ctx).Where(r.UserID.Eq(uint64(req.UserId))).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_RecordNotFound, "简历不存在"))
		} else {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "查询简历失败"))
		}
		return
	}

	// 构造返回数据
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.ResumeEntity = &base.ResumeEntity{
		Id: int64(foundResume.ID),
		Resume: &base.Resume{
			UserId:  int64(foundResume.UserID),
			FileUrl: foundResume.FileURL,
			ParsedContent: func() string {
				if foundResume.ParsedContent != nil {
					return *foundResume.ParsedContent
				}
				return ""
			}(),
			CreatedAt: foundResume.CreatedAt.Unix(),
		},
	}

	return
}

// DeleteResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteResume(ctx context.Context, req *user.DeleteResumeRequest) (resp *user.DeleteResumeResponse, err error) {
	resp = &user.DeleteResumeResponse{
		BaseResp: &base.BaseResponse{},
	}

	// 参数验证
	if req.UserId <= 0 || req.ResumeId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID或简历ID无效"))
		return
	}

	r := s.Query.Resume

	// 验证简历是否属于该用户
	foundResume, err := r.WithContext(ctx).Where(
		r.ID.Eq(uint64(req.ResumeId)),
		r.UserID.Eq(uint64(req.UserId)),
	).First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_RecordNotFound, "简历不存在或无权限"))
		} else {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "查询简历失败"))
		}
		return
	}

	// 软删除简历
	_, err = r.WithContext(ctx).Where(r.ID.Eq(foundResume.ID)).Delete()
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "删除简历失败"))
		return
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LogoutRequest) (resp *user.LogoutResponse, err error) {
	resp = &user.LogoutResponse{
		BaseResp: &base.BaseResponse{},
	}

	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID无效"))
		return
	}

	// TODO: Phase 4 实现 Token 黑名单（需要 Redis）
	// 当前简单实现：直接返回成功，客户端删除 Token
	// 完整实现需要：
	// 1. 将 Token 加入 Redis 黑名单
	// 2. 设置过期时间为 Token 的剩余有效期
	// 3. 在认证中间件中检查黑名单

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return
}

// DeleteAccount implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteAccount(ctx context.Context, req *user.DeleteAccountRequest) (resp *user.DeleteAccountResponse, err error) {
	resp = &user.DeleteAccountResponse{
		BaseResp: &base.BaseResponse{},
	}

	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "用户ID无效"))
		return
	}

	if req.Password == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_ParamsErr, "密码不能为空"))
		return
	}

	u := s.Query.User

	// 查找用户
	foundUser, err := u.WithContext(ctx).Where(u.ID.Eq(uint64(req.UserId))).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserNotFound, "用户不存在"))
		} else {
			resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "查询用户失败"))
		}
		return
	}

	// 验证密码
	encryptedPassword := s.EncryptPassword(req.Password)
	if encryptedPassword != foundUser.Password {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_PasswordError, "密码错误"))
		return
	}

	// 软删除用户（将状态设置为已删除）
	// 注意：这里使用软删除而不是物理删除，保留数据用于审计
	_, err = u.WithContext(ctx).Where(u.ID.Eq(uint64(req.UserId))).Update(u.Status, 0)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.NewErrNo(base.ErrCode_UserSrvErr, "注销账号失败"))
		return
	}

	// TODO: Phase 5 实现
	// 1. 删除用户的简历数据
	// 2. 删除用户的面试记录
	// 3. 将当前 Token 加入黑名单

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return
}
