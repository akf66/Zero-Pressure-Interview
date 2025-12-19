namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

// ==================== 用户注册 ====================
struct RegisterRequest {
    1: string email (api.body="email", api.vd="len($) > 0 && len($) < 100"),
    2: string password (api.body="password", api.vd="len($) >= 6 && len($) <= 20"),
    3: string nickname (api.body="nickname", api.vd="len($) < 50"),
}

struct RegisterResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: i64 user_id (api.body="user_id"),
}

// ==================== 用户登录 ====================
struct LoginRequest {
    1: string email (api.body="email", api.vd="len($) > 0"),
    2: string password (api.body="password", api.vd="len($) > 0"),
}

struct LoginResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: string token (api.body="token"),
    4: user.UserEntity user_entity (api.body="user"),
}

// ==================== 获取用户信息 ====================
struct GetProfileRequest {}

struct GetProfileResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: user.UserEntity user_entity (api.body="user"),
}

// ==================== 更新用户信息 ====================
struct UpdateProfileRequest {
    1: string nickname (api.body="nickname", api.vd="len($) < 50"),
    2: string avatar_url (api.body="avatar_url", api.vd="len($) < 500"),
}

// ==================== 修改密码 ====================
struct ChangePasswordRequest {
    1: string old_password (api.body="old_password", api.vd="len($) > 0"),
    2: string new_password (api.body="new_password", api.vd="len($) >= 6 && len($) <= 20"),
}

// ==================== 上传简历 ====================
struct UploadResumeRequest {
    1: string file_url (api.body="file_url", api.vd="len($) > 0"),
    2: string parsed_content (api.body="parsed_content"),
}

struct UploadResumeResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: i64 resume_id (api.body="resume_id"),
}

// ==================== 获取简历 ====================
struct GetResumeRequest {}

struct GetResumeResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: user.ResumeEntity resume_entity (api.body="resume"),
}

// ==================== 用户服务接口 ====================
service UserService {
    RegisterResponse Register(1: RegisterRequest req) (api.post="/api/v1/user/register"),
    LoginResponse Login(1: LoginRequest req) (api.post="/api/v1/user/login"),
    GetProfileResponse GetProfile(1: GetProfileRequest req) (api.get="/api/v1/user/profile"),
    common.NilResponse UpdateProfile(1: UpdateProfileRequest req) (api.put="/api/v1/user/profile"),
    common.NilResponse ChangePassword(1: ChangePasswordRequest req) (api.post="/api/v1/user/password"),
    UploadResumeResponse UploadResume(1: UploadResumeRequest req) (api.post="/api/v1/user/resume"),
    GetResumeResponse GetResume(1: GetResumeRequest req) (api.get="/api/v1/user/resume"),
}