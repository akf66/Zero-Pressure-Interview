namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

// ==================== 发送验证码 ====================
struct SendVerifyCodeRequest {
    1: user.VerifyCodeType code_type      // 验证码类型（邮箱/短信）
    2: user.VerifyCodePurpose purpose     // 验证码用途
    3: string target       // 目标（邮箱地址或手机号）
}

struct SendVerifyCodeResponse {
    1: common.BaseResponse base_resp
}

// ==================== 用户注册 ====================
struct RegisterRequest {
    1: optional string email              // 邮箱（邮箱注册时必填）
    2: optional string phone              // 手机号（手机号注册时必填）
    3: string password                    // 密码
    4: string nickname                    // 昵称
    5: string verify_code                 // 验证码（必填）
    6: optional string username// 用户名（可选）
}

struct RegisterResponse {
    1: common.BaseResponse base_resp
    2: i64 user_id
}

// ==================== 用户登录 ====================
struct LoginRequest {
    1: string account                     // 账号（邮箱/用户名/手机号）
    2: optional string password           // 密码（密码登录时必填）
    3: optional string verify_code        // 验证码（验证码登录时必填）
}

struct LoginResponse {
    1: common.BaseResponse base_resp
    2: string token
    3: user.UserEntity user_entity
}

// ==================== 退出登录 ====================
struct LogoutRequest {
    1: i64 user_id
    2: string token
}

struct LogoutResponse {
    1: common.BaseResponse base_resp
}

// ==================== 注销账号 ====================
struct DeleteAccountRequest {
    1: i64 user_id
    2: string password      // 需要验证密码
}

struct DeleteAccountResponse {
    1: common.BaseResponse base_resp
}

// ==================== 获取用户信息 ====================
struct GetUserRequest {
    1: i64 user_id
}

struct GetUserResponse {
    1: common.BaseResponse base_resp
    2: user.UserEntity user_entity
}

// ==================== 更新用户信息 ====================
struct UpdateUserRequest {
    1: i64 user_id
    2: optional string nickname
    3: optional string avatar_url
}

struct UpdateUserResponse {
    1: common.BaseResponse base_resp
}

// ==================== 修改密码 ====================
struct ChangePasswordRequest {
    1: i64 user_id
    2: string old_password
    3: string new_password
}

struct ChangePasswordResponse {
    1: common.BaseResponse base_resp
}

// ==================== 重置密码 ====================
struct ResetPasswordRequest {
    1: string account                     // 账号（邮箱/手机号）
    2: string verify_code                 // 验证码
    3: string new_password                // 新密码
}

struct ResetPasswordResponse {
    1: common.BaseResponse base_resp
}

// ==================== 上传简历 ====================
struct UploadResumeRequest {
    1: i64 user_id
    2: string file_url
    3: string parsed_content
}

struct UploadResumeResponse {
    1: common.BaseResponse base_resp
    2: i64 resume_id
}

// ==================== 获取简历 ====================
struct GetResumeRequest {
    1: i64 user_id
}

struct GetResumeResponse {
    1: common.BaseResponse base_resp
    2: user.ResumeEntity resume_entity
}

// ==================== 删除简历 ====================
struct DeleteResumeRequest {
    1: i64 user_id
    2: i64 resume_id
}

struct DeleteResumeResponse {
    1: common.BaseResponse base_resp
}

// ==================== 用户服务接口 ====================
service UserService {
    // 健康检查
    common.HealthCheckResponse HealthCheck()
    
    // 验证码
    SendVerifyCodeResponse SendVerifyCode(1: SendVerifyCodeRequest req)
    
    // 用户认证
    RegisterResponse Register(1: RegisterRequest req)
    LoginResponse Login(1: LoginRequest req)
    LogoutResponse Logout(1: LogoutRequest req)
    DeleteAccountResponse DeleteAccount(1: DeleteAccountRequest req)
    
    // 用户信息管理
    GetUserResponse GetUser(1: GetUserRequest req)
    UpdateUserResponse UpdateUser(1: UpdateUserRequest req)
    ChangePasswordResponse ChangePassword(1: ChangePasswordRequest req)
    ResetPasswordResponse ResetPassword(1: ResetPasswordRequest req)
    
    // 简历管理
    UploadResumeResponse UploadResume(1: UploadResumeRequest req)
    GetResumeResponse GetResume(1: GetResumeRequest req)
    DeleteResumeResponse DeleteResume(1: DeleteResumeRequest req)
}