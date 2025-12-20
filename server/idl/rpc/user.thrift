namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

// ==================== 用户注册 ====================
struct RegisterRequest {
    1: string email
    2: string password
    3: string nickname
}

struct RegisterResponse {
    1: common.BaseResponse base_resp
    2: i64 user_id
}

// ==================== 用户登录 ====================
struct LoginRequest {
    1: string email
    2: string password
}

struct LoginResponse {
    1: common.BaseResponse base_resp
    2: string token
    3: user.UserEntity user_entity
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
    2: string nickname
    3: string avatar_url
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
    
    // 用户认证
    RegisterResponse Register(1: RegisterRequest req)
    LoginResponse Login(1: LoginRequest req)
    
    // 用户信息管理
    GetUserResponse GetUser(1: GetUserRequest req)
    UpdateUserResponse UpdateUser(1: UpdateUserRequest req)
    ChangePasswordResponse ChangePassword(1: ChangePasswordRequest req)
    
    // 简历管理
    UploadResumeResponse UploadResume(1: UploadResumeRequest req)
    GetResumeResponse GetResume(1: GetResumeRequest req)
    DeleteResumeResponse DeleteResume(1: DeleteResumeRequest req)
}