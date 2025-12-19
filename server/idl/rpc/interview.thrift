namespace go interview

include "../base/common.thrift"
include "../base/interview.thrift"

// ==================== 开始面试 ====================
struct StartInterviewRequest {
    1: i64 user_id
    2: i32 type              // 类型：1-专项/2-综合
    3: string category       // 专项类型（如：golang/mysql）或岗位（如：Golang后端工程师）
    4: i32 difficulty        // 难度：1-简单/2-中等/3-困难（专项面试）
    5: i32 round             // 面试轮次（综合面试）：1-一面/2-二面/3-三面/4-HR面
    6: i64 resume_id         // 简历ID（可选，用于个性化问题）
}

struct StartInterviewResponse {
    1: common.BaseResponse base_resp
    2: i64 interview_id
    3: string first_question  // AI的第一个问题
}

// ==================== 提交回答 ====================
struct SubmitAnswerRequest {
    1: i64 interview_id
    2: i64 user_id
    3: string answer         // 用户的回答
}

struct SubmitAnswerResponse {
    1: common.BaseResponse base_resp
    2: string next_question  // AI的下一个问题
    3: bool is_finished      // 是否结束面试
}

// ==================== 结束面试 ====================
struct FinishInterviewRequest {
    1: i64 interview_id
    2: i64 user_id
}

struct FinishInterviewResponse {
    1: common.BaseResponse base_resp
    2: i32 score             // 评分
    3: string evaluation     // AI评估报告
}

// ==================== 获取面试详情 ====================
struct GetInterviewDetailRequest {
    1: i64 interview_id
    2: i64 user_id
}

struct GetInterviewDetailResponse {
    1: common.BaseResponse base_resp
    2: interview.InterviewEntity interview_entity
    3: list<interview.InterviewMessage> messages  // 对话记录
}

// ==================== 获取面试历史 ====================
struct GetInterviewHistoryRequest {
    1: i64 user_id
    2: i32 type              // 类型筛选：1-专项/2-综合
    3: common.PageRequest page
}

struct GetInterviewHistoryResponse {
    1: common.BaseResponse base_resp
    2: list<interview.InterviewEntity> interviews
    3: common.PageResponse page
}

// ==================== 简历分析 ====================
struct AnalyzeResumeRequest {
    1: i64 user_id
    2: i64 resume_id
}

struct AnalyzeResumeResponse {
    1: common.BaseResponse base_resp
    2: string analysis       // AI分析结果
    3: list<string> suggestions  // 优化建议
    4: list<string> matched_positions  // 匹配的岗位
}

// ==================== 获取能力分析 ====================
struct GetAbilityAnalysisRequest {
    1: i64 user_id
}

struct GetAbilityAnalysisResponse {
    1: common.BaseResponse base_resp
    2: list<interview.AbilityScore> abilities  // 各维度能力评分
    3: string summary        // 综合评价
}

// ==================== Interview服务接口 ====================
service InterviewService {
    StartInterviewResponse StartInterview(1: StartInterviewRequest req)
    SubmitAnswerResponse SubmitAnswer(1: SubmitAnswerRequest req)
    FinishInterviewResponse FinishInterview(1: FinishInterviewRequest req)
    GetInterviewDetailResponse GetInterviewDetail(1: GetInterviewDetailRequest req)
    GetInterviewHistoryResponse GetInterviewHistory(1: GetInterviewHistoryRequest req)
    AnalyzeResumeResponse AnalyzeResume(1: AnalyzeResumeRequest req)
    GetAbilityAnalysisResponse GetAbilityAnalysis(1: GetAbilityAnalysisRequest req)
}