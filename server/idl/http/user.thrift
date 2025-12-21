namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

// ==================== 发送验证码 ====================
struct SendVerifyCodeRequest {
    1: i32 code_type (api.body="code_type", api.vd="$ > 0")      // 验证码类型（1-邮箱，2-短信）
    2: i32 purpose (api.body="purpose", api.vd="$ > 0")          // 验证码用途
    3: string target (api.body="target", api.vd="len($) > 0")    // 目标（邮箱地址或手机号）
}

// ==================== 用户注册 ====================
struct RegisterRequest {
    1: optional string email (api.body="email", api.vd="len($) < 100")
    2: optional string phone (api.body="phone", api.vd="len($) < 20")
    3: string password (api.body="password", api.vd="len($) >= 6 && len($) <= 20")
    4: string nickname (api.body="nickname", api.vd="len($) < 50")
    5: string verify_code (api.body="verify_code", api.vd="len($) > 0")
    6: optional string username (api.body="username", api.vd="len($) < 50")
}

struct RegisterResponse {
    1: i64 status_code (api.body="code")
    2: string status_msg (api.body="message")
    3: i64 user_id (api.body="user_id")
}

// ==================== 用户登录 ====================
struct LoginRequest {
    1: string account (api.body="account", api.vd="len($) > 0")           // 账号（邮箱/用户名/手机号）
    2: optional string password (api.body="password")// 密码（密码登录时必填）
    3: optional string verify_code (api.body="verify_code")                // 验证码（验证码登录时必填）
}

struct LoginResponse {
    1: i64 status_code (api.body="code")
    2: string status_msg (api.body="message")
    3: string token (api.body="token")
    4: user.UserEntity user_entity (api.body="user")
}

// ==================== 退出登录 ====================
struct LogoutRequest {}

// ==================== 注销账号 ====================
struct DeleteAccountRequest {
    1: string password (api.body="password", api.vd="len($) > 0")      // 需要验证密码
}

// ==================== 获取用户信息 ====================
struct GetProfileRequest {}

struct GetProfileResponse {
    1: i64 status_code (api.body="code")
    2: string status_msg (api.body="message")
    3: user.UserEntity user_entity (api.body="user")
}

// ==================== 更新用户信息 ====================
struct UpdateProfileRequest {
    1: optional string nickname (api.body="nickname", api.vd="len($) < 50")
    2: optional string avatar_url (api.body="avatar_url", api.vd="len($) < 500")
}

// ==================== 修改密码 ====================
struct ChangePasswordRequest {
    1: string old_password (api.body="old_password", api.vd="len($) > 0")
    2: string new_password (api.body="new_password", api.vd="len($) >= 6 && len($) <= 20")
}

// ==================== 重置密码 ====================
struct ResetPasswordRequest {
    1: string account (api.body="account", api.vd="len($) > 0")           // 账号（邮箱/手机号）
    2: string verify_code (api.body="verify_code", api.vd="len($) > 0")   // 验证码
    3: string new_password (api.body="new_password", api.vd="len($) >= 6 && len($) <= 20")
}

// ==================== 上传简历 ====================
struct UploadResumeRequest {
    1: string file_url (api.body="file_url", api.vd="len($) > 0")
    2: string parsed_content (api.body="parsed_content")
}

struct UploadResumeResponse {
    1: i64 status_code (api.body="code")
    2: string status_msg (api.body="message")
    3: i64 resume_id (api.body="resume_id")
}

// ==================== 获取简历 ====================
struct GetResumeRequest {}

struct GetResumeResponse {
    1: i64 status_code (api.body="code")
    2: string status_msg (api.body="message")
    3: user.ResumeEntity resume_entity (api.body="resume")
}

// ==================== 用户服务接口 ====================
service UserService {
    // 验证码
    common.NilResponse SendVerifyCode(1: SendVerifyCodeRequest req) (api.post="/api/v1/user/verify-code")
    
    // 用户认证
    RegisterResponse Register(1: RegisterRequest req) (api.post="/api/v1/user/register")
    LoginResponse Login(1: LoginRequest req) (api.post="/api/v1/user/login")
    common.NilResponse Logout(1: LogoutRequest req) (api.post="/api/v1/user/logout")
    common.NilResponse DeleteAccount(1: DeleteAccountRequest req) (api.delete="/api/v1/user/account")
    
    // 用户信息管理
    GetProfileResponse GetProfile(1: GetProfileRequest req) (api.get="/api/v1/user/profile")
    common.NilResponse UpdateProfile(1: UpdateProfileRequest req) (api.put="/api/v1/user/profile")
    common.NilResponse ChangePassword(1: ChangePasswordRequest req) (api.post="/api/v1/user/password")
    common.NilResponse ResetPassword(1: ResetPasswordRequest req) (api.post="/api/v1/user/password/reset")
    
    // 简历管理
    UploadResumeResponse UploadResume(1: UploadResumeRequest req) (api.post="/api/v1/user/resume")
    GetResumeResponse GetResume(1: GetResumeRequest req) (api.get="/api/v1/user/resume")
}