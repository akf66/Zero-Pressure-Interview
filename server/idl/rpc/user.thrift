namespace go user

include "../base/base.thrift"

// ==================== 用户注册 ====================
struct RegisterReq {
    1: required string email
    2: required string password
    3: optional string nickname
}

struct RegisterResp {
    1: base.BaseResp base_resp
    2: optional i64 user_id
}

// ==================== 用户登录 ====================
struct LoginReq {
    1: required string email
    2: required string password
}

struct LoginResp {
    1: base.BaseResp base_resp
    2: optional string token
    3: optional base.UserInfo user
}

// ==================== 获取用户信息 ====================
struct GetUserInfoReq {
    1: required i64 user_id
}

struct GetUserInfoResp {
    1: base.BaseResp base_resp
    2: optional base.UserInfo user
}

// ==================== 更新用户信息 ====================
struct UpdateUserInfoReq {
    1: required i64 user_id
    2: optional string nickname
    3: optional string avatar_url
}

struct UpdateUserInfoResp {
    1: base.BaseResp base_resp
}

// ==================== 修改密码 ====================
struct ChangePasswordReq {
    1: required i64 user_id
    2: required string old_password
    3: required string new_password
}

struct ChangePasswordResp {
    1: base.BaseResp base_resp
}

// ==================== 上传简历 ====================
struct UploadResumeReq {
    1: required i64 user_id
    2: required string file_url
    3: optional string parsed_content
}

struct UploadResumeResp {
    1: base.BaseResp base_resp
    2: optional i64 resume_id
}

// ==================== 获取简历信息 ====================
struct GetResumeReq {
    1: required i64 user_id
}

struct GetResumeResp {
    1: base.BaseResp base_resp
    2: optional base.ResumeInfo resume
}

// ==================== 删除简历 ====================
struct DeleteResumeReq {
    1: required i64 user_id
    2: required i64 resume_id
}

struct DeleteResumeResp {
    1: base.BaseResp base_resp
}

// ==================== 用户服务接口 ====================
service UserService {
    // 用户注册
    RegisterResp Register(1: RegisterReq req)
    
    // 用户登录
    LoginResp Login(1: LoginReq req)
    
    // 获取用户信息
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq req)
    
    // 更新用户信息
    UpdateUserInfoResp UpdateUserInfo(1: UpdateUserInfoReq req)
    
    // 修改密码
    ChangePasswordResp ChangePassword(1: ChangePasswordReq req)
    
    // 上传简历
    UploadResumeResp UploadResume(1: UploadResumeReq req)
    
    // 获取简历信息
    GetResumeResp GetResume(1: GetResumeReq req)
    
    // 删除简历
    DeleteResumeResp DeleteResume(1: DeleteResumeReq req)
}