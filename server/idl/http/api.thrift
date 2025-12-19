namespace go api

include "../base/base.thrift"

// ==================== 用户相关接口 ====================

// 用户注册
struct RegisterRequest {
    1: required string email (api.body="email")
    2: required string password (api.body="password")
    3: optional string nickname (api.body="nickname")
}

struct RegisterResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional i64 user_id (api.body="user_id")
}

// 用户登录
struct LoginRequest {
    1: required string email (api.body="email")
    2: required string password (api.body="password")
}

struct LoginResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional string token (api.body="token")
    4: optional base.UserInfo user (api.body="user")
}

// 获取用户信息
struct GetProfileResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional base.UserInfo user (api.body="user")
}

// 更新用户信息
struct UpdateProfileRequest {
    1: optional string nickname (api.body="nickname")
    2: optional string avatar_url (api.body="avatar_url")
}

struct UpdateProfileResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
}

// 修改密码
struct ChangePasswordRequest {
    1: required string old_password (api.body="old_password")
    2: required string new_password (api.body="new_password")
}

struct ChangePasswordResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
}

// 上传简历
struct UploadResumeRequest {
    1: required string file_url (api.body="file_url")
    2: optional string parsed_content (api.body="parsed_content")
}

struct UploadResumeResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional i64 resume_id (api.body="resume_id")
}

// 获取简历
struct GetResumeResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional base.ResumeInfo resume (api.body="resume")
}

// ==================== 面试相关接口 ====================

// 开始面试
struct StartInterviewRequest {
    1: required i32 type (api.body="type")
    2: required string category (api.body="category")
    3: optional i32 difficulty (api.body="difficulty")
    4: optional i32 round (api.body="round")
    5: optional i64 resume_id (api.body="resume_id")
}

struct StartInterviewResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional i64 interview_id (api.body="interview_id")
    4: optional string first_question (api.body="first_question")
}

// 提交回答
struct SubmitAnswerRequest {
    1: required string answer (api.body="answer")
}

struct SubmitAnswerResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional string next_question (api.body="next_question")
    4: optional bool is_finished (api.body="is_finished")
}

// 结束面试
struct FinishInterviewResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional i32 score (api.body="score")4: optional string evaluation (api.body="evaluation")
}

// 获取面试详情
struct GetInterviewDetailResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional base.InterviewInfo interview (api.body="interview")4: optional list<base.InterviewMessage> messages (api.body="messages")
}

// 获取面试历史
struct GetInterviewHistoryRequest {
    1: optional i32 type (api.query="type")
    2: optional i32 page (api.query="page")
    3: optional i32 page_size (api.query="page_size")
}

struct GetInterviewHistoryResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional list<base.InterviewInfo> interviews (api.body="interviews")4: optional base.PageResp page (api.body="page")
}

// 简历分析
struct AnalyzeResumeRequest {
    1: required i64 resume_id (api.body="resume_id")
}

struct AnalyzeResumeResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional string analysis (api.body="analysis")
    4: optional list<string> suggestions (api.body="suggestions")
    5: optional list<string> matched_positions (api.body="matched_positions")
}

// 获取能力分析
struct GetAbilityAnalysisResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional list<map<string, string>> abilities (api.body="abilities")
    4: optional string summary (api.body="summary")
}

// ==================== 题库相关接口 ====================

// 获取分类列表
struct GetCategoriesResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional list<map<string, string>> categories (api.body="categories")
}

// 获取题目列表
struct GetQuestionListRequest {
    1: optional string category (api.query="category")
    2: optional i32 difficulty (api.query="difficulty")
    3: optional string keyword (api.query="keyword")
    4: optional i32 page (api.query="page")
    5: optional i32 page_size (api.query="page_size")
}

struct GetQuestionListResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional list<base.QuestionInfo> questions (api.body="questions")
    4: optional base.PageResp page (api.body="page")
}

// 获取题目详情
struct GetQuestionDetailResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional base.QuestionInfo question (api.body="question")
}

// 收藏题目
struct FavoriteQuestionRequest {
    1: required i64 question_id (api.body="question_id")
}

struct FavoriteQuestionResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
}

// 获取收藏列表
struct GetFavoriteQuestionsRequest {
    1: optional i32 page (api.query="page")
    2: optional i32 page_size (api.query="page_size")
}

struct GetFavoriteQuestionsResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional list<base.QuestionInfo> questions (api.body="questions")
    4: optional base.PageResp page (api.body="page")
}

// ==================== 文件相关接口 ====================

// 获取上传URL
struct GetUploadUrlRequest {
    1: required string file_name (api.query="file_name")
    2: required string file_type (api.query="file_type")3: optional string content_type (api.query="content_type")
}

struct GetUploadUrlResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional string upload_url (api.body="upload_url")
    4: optional string file_key (api.body="file_key")5: optional i64 expires_at (api.body="expires_at")
}

// 确认上传
struct ConfirmUploadRequest {
    1: required string file_key (api.body="file_key")
    2: required string file_type (api.body="file_type")
}

struct ConfirmUploadResponse {
    1: i32 code (api.body="code")
    2: string message (api.body="message")
    3: optional string file_url (api.body="file_url")
}

// ==================== API服务定义 ====================
service APIService {
    // 用户接口
    RegisterResponse Register(1: RegisterRequest req) (api.post="/api/v1/user/register")
    LoginResponse Login(1: LoginRequest req) (api.post="/api/v1/user/login")
    GetProfileResponse GetProfile() (api.get="/api/v1/user/profile")
    UpdateProfileResponse UpdateProfile(1: UpdateProfileRequest req) (api.put="/api/v1/user/profile")
    ChangePasswordResponse ChangePassword(1: ChangePasswordRequest req) (api.post="/api/v1/user/password")
    UploadResumeResponse UploadResume(1: UploadResumeRequest req) (api.post="/api/v1/user/resume")
    GetResumeResponse GetResume() (api.get="/api/v1/user/resume")
    
    // 面试接口
    StartInterviewResponse StartInterview(1: StartInterviewRequest req) (api.post="/api/v1/interview/start")
    SubmitAnswerResponse SubmitAnswer(1: SubmitAnswerRequest req, 2: i64 interview_id) (api.post="/api/v1/interview/:interview_id/answer", api.path="interview_id")
    FinishInterviewResponse FinishInterview(1: i64 interview_id) (api.post="/api/v1/interview/:interview_id/finish", api.path="interview_id")
    GetInterviewDetailResponse GetInterviewDetail(1: i64 interview_id) (api.get="/api/v1/interview/:interview_id", api.path="interview_id")
    GetInterviewHistoryResponse GetInterviewHistory(1: GetInterviewHistoryRequest req) (api.get="/api/v1/interview/history")
    AnalyzeResumeResponse AnalyzeResume(1: AnalyzeResumeRequest req) (api.post="/api/v1/interview/analyze-resume")
    GetAbilityAnalysisResponse GetAbilityAnalysis() (api.get="/api/v1/interview/ability-analysis")
    
    // 题库接口
    GetCategoriesResponse GetCategories() (api.get="/api/v1/question/categories")
    GetQuestionListResponse GetQuestionList(1: GetQuestionListRequest req) (api.get="/api/v1/question/list")
    GetQuestionDetailResponse GetQuestionDetail(1: i64 question_id) (api.get="/api/v1/question/:question_id", api.path="question_id")
    FavoriteQuestionResponse FavoriteQuestion(1: FavoriteQuestionRequest req) (api.post="/api/v1/question/favorite")
    FavoriteQuestionResponse UnfavoriteQuestion(1: i64 question_id) (api.delete="/api/v1/question/favorite/:question_id", api.path="question_id")
    GetFavoriteQuestionsResponse GetFavoriteQuestions(1: GetFavoriteQuestionsRequest req) (api.get="/api/v1/question/favorites")
    
    // 文件接口
    GetUploadUrlResponse GetUploadUrl(1: GetUploadUrlRequest req) (api.get="/api/v1/file/upload-url")
    ConfirmUploadResponse ConfirmUpload(1: ConfirmUploadRequest req) (api.post="/api/v1/file/confirm")
}