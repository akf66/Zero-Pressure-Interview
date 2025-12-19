namespace go base

// 统一响应结构
struct BaseResp {
    1: i32 code          // 错误码：0-成功，其他-失败
    2: string message    // 响应消息
}

// 分页请求
struct PageReq {
    1: i32 page = 1      // 页码，从1开始
    2: i32 page_size = 10 // 每页数量
}

// 分页响应
struct PageResp {
    1: i64 total         // 总数
    2: i32 page// 当前页
    3: i32 page_size     // 每页数量
}

// 用户基本信息
struct UserInfo {
    1: i64 id
    2: string email
    3: string nickname
    4: string avatar_url
    5: i64 created_at    // 创建时间戳（秒）
}

// 简历信息
struct ResumeInfo {
    1: i64 id
    2: i64 user_id
    3: string file_url
    4: string parsed_content  // JSON格式的解析内容
    5: i64 created_at
}

// 题目信息
struct QuestionInfo {
    1: i64 id
    2: string category        // 分类：golang/java/mysql等
    3: i32 difficulty         // 难度：1-简单/2-中等/3-困难
    4: string title
    5: string content
    6: string answer
    7: list<string> tags      // 标签列表
    8: i64 created_at
}

// 面试记录信息
struct InterviewInfo {
    1: i64 id
    2: i64 user_id
    3: i32 type               // 类型：1-专项/2-综合
    4: string category        // 专项类型或岗位
    5: i32 round              // 面试轮次（综合面试）
    6: i32 status             // 状态：0-进行中/1-已完成
    7: i32 score              // 评分
    8: string evaluation      // AI评估
    9: i64 created_at
    10: i64 finished_at
}

// 面试对话消息
struct InterviewMessage {
    1: i64 id
    2: i64 interview_id
    3: string role// interviewer/candidate
    4: string content
    5: i64 created_at
}

// 错误码定义
const i32 SUCCESS = 0
const i32 PARAM_ERROR = 10001
const i32 NOT_LOGIN = 10002
const i32 NO_PERMISSION = 10003
const i32 USER_NOT_EXIST = 20001
const i32 PASSWORD_ERROR = 20002
const i32 INTERVIEW_NOT_EXIST = 30001
const i32 QUESTION_NOT_EXIST = 40001
const i32 FILE_UPLOAD_ERROR = 50001