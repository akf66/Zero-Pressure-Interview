namespace go interview

include "../base/common.thrift"
include "../base/interview.thrift"

// ==================== 开始面试 ====================
struct StartInterviewRequest {
    1: i32 type (api.body="type", api.vd="$ > 0 && $ <= 2"),              // 类型：1-专项/2-综合
    2: string category (api.body="category", api.vd="len($) > 0"),        // 专项类型或岗位
    3: i32 difficulty (api.body="difficulty", api.vd="$ >= 0 && $ <= 3"), // 难度：1-简单/2-中等/3-困难
    4: i32 round (api.body="round", api.vd="$ >= 0 && $ <= 4"),           // 面试轮次
    5: i64 resume_id (api.body="resume_id"),                              // 简历ID（可选）
}

struct StartInterviewResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: i64 interview_id (api.body="interview_id"),4: string first_question (api.body="first_question"),
}

// ==================== 提交回答 ====================
struct SubmitAnswerRequest {
    1: string answer (api.body="answer", api.vd="len($) > 0"),
}

struct SubmitAnswerResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: string next_question (api.body="next_question"),
    4: bool is_finished (api.body="is_finished"),
}

// ==================== 结束面试 ====================
struct FinishInterviewResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: i32 score (api.body="score"),
    4: string evaluation (api.body="evaluation"),
}

// ==================== 获取面试详情 ====================
struct GetInterviewDetailResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: interview.InterviewEntity interview_entity (api.body="interview"),
    4: list<interview.InterviewMessage> messages (api.body="messages"),
}

// ==================== 获取面试历史 ====================
struct GetInterviewHistoryRequest {
    1: i32 type (api.query="type"),           // 类型筛选
    2: i32 page (api.query="page"),
    3: i32 page_size (api.query="page_size"),
}

struct GetInterviewHistoryResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: list<interview.InterviewEntity> interviews (api.body="interviews"),
    4: common.PageResponse page (api.body="page"),
}

// ==================== 简历分析 ====================
struct AnalyzeResumeRequest {
    1: i64 resume_id (api.body="resume_id", api.vd="$ > 0"),
}

struct AnalyzeResumeResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: string analysis (api.body="analysis"),
    4: list<string> suggestions (api.body="suggestions"),5: list<string> matched_positions (api.body="matched_positions"),
}

// ==================== 获取能力分析 ====================
struct GetAbilityAnalysisRequest {}

struct GetAbilityAnalysisResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: list<interview.AbilityScore> abilities (api.body="abilities"),4: string summary (api.body="summary"),
}

// ==================== 面试服务接口 ====================
service InterviewService {
    StartInterviewResponse StartInterview(1: StartInterviewRequest req) (api.post="/api/v1/interview/start"),
    SubmitAnswerResponse SubmitAnswer(1: SubmitAnswerRequest req, 2: i64 interview_id) (api.post="/api/v1/interview/:interview_id/answer", api.path="interview_id"),
    FinishInterviewResponse FinishInterview(1: i64 interview_id) (api.post="/api/v1/interview/:interview_id/finish", api.path="interview_id"),
    GetInterviewDetailResponse GetInterviewDetail(1: i64 interview_id) (api.get="/api/v1/interview/:interview_id", api.path="interview_id"),
    GetInterviewHistoryResponse GetInterviewHistory(1: GetInterviewHistoryRequest req) (api.get="/api/v1/interview/history"),
    AnalyzeResumeResponse AnalyzeResume(1: AnalyzeResumeRequest req) (api.post="/api/v1/interview/analyze-resume"),
    GetAbilityAnalysisResponse GetAbilityAnalysis(1: GetAbilityAnalysisRequest req) (api.get="/api/v1/interview/ability-analysis"),
}