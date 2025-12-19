namespace go agent

include "../base/base.thrift"

// ==================== 开始面试 ====================
struct StartInterviewReq {
    1: required i64 user_id
    2: required i32 type              // 类型：1-专项/2-综合
    3: required string category       // 专项类型（如：golang/mysql）或岗位（如：Golang后端工程师）
    4: optional i32 difficulty        // 难度：1-简单/2-中等/3-困难（专项面试）
    5: optional i32 round             // 面试轮次（综合面试）：1-一面/2-二面/3-三面/4-HR面
    6: optional i64 resume_id         // 简历ID（可选，用于个性化问题）
}

struct StartInterviewResp {
    1: base.BaseResp base_resp
    2: optional i64 interview_id
    3: optional string first_question  // AI的第一个问题
}

// ==================== 提交回答 ====================
struct SubmitAnswerReq {
    1: required i64 interview_id
    2: required i64 user_id
    3: required string answer         // 用户的回答
}

struct SubmitAnswerResp {
    1: base.BaseResp base_resp
    2: optional string next_question  // AI的下一个问题
    3: optional bool is_finished      // 是否结束面试
}

// ==================== 结束面试 ====================
struct FinishInterviewReq {
    1: required i64 interview_id
    2: required i64 user_id
}

struct FinishInterviewResp {
    1: base.BaseResp base_resp
    2: optional i32 score             // 评分
    3: optional string evaluation     // AI评估报告
}

// ==================== 获取面试详情 ====================
struct GetInterviewDetailReq {
    1: required i64 interview_id
    2: required i64 user_id
}

struct GetInterviewDetailResp {
    1: base.BaseResp base_resp
    2: optional base.InterviewInfo interview
    3: optional list<base.InterviewMessage> messages  // 对话记录
}

// ==================== 获取面试历史 ====================
struct GetInterviewHistoryReq {
    1: required i64 user_id
    2: optional i32 type              // 类型筛选：1-专项/2-综合
    3: optional base.PageReq page
}

struct GetInterviewHistoryResp {
    1: base.BaseResp base_resp
    2: optional list<base.InterviewInfo> interviews
    3: optional base.PageResp page
}

// ==================== 简历分析 ====================
struct AnalyzeResumeReq {
    1: required i64 user_id
    2: required i64 resume_id
}

struct AnalyzeResumeResp {
    1: base.BaseResp base_resp
    2: optional string analysis       // AI分析结果
    3: optional list<string> suggestions  // 优化建议
    4: optional list<string> matched_positions  // 匹配的岗位
}

// ==================== 获取能力分析 ====================
struct GetAbilityAnalysisReq {
    1: required i64 user_id
}

struct AbilityScore {
    1: string dimension               // 能力维度（如：编程基础、算法、系统设计等）
    2: i32 score                      // 分数（0-100）
}

struct GetAbilityAnalysisResp {
    1: base.BaseResp base_resp
    2: optional list<AbilityScore> abilities  // 各维度能力评分
    3: optional string summary        // 综合评价
}

// ==================== Agent服务接口 ====================
service AgentService {
    // 开始面试
    StartInterviewResp StartInterview(1: StartInterviewReq req)
    
    // 提交回答
    SubmitAnswerResp SubmitAnswer(1: SubmitAnswerReq req)
    
    // 结束面试
    FinishInterviewResp FinishInterview(1: FinishInterviewReq req)
    
    // 获取面试详情
    GetInterviewDetailResp GetInterviewDetail(1: GetInterviewDetailReq req)
    
    // 获取面试历史
    GetInterviewHistoryResp GetInterviewHistory(1: GetInterviewHistoryReq req)
    
    // 简历分析
    AnalyzeResumeResp AnalyzeResume(1: AnalyzeResumeReq req)
    
    // 获取能力分析
    GetAbilityAnalysisResp GetAbilityAnalysis(1: GetAbilityAnalysisReq req)
}