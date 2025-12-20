namespace go base

// 面试实体
struct InterviewEntity {
    1: i64 id
    2: Interview interview
}

// 面试领域对象
struct Interview {
    1: i64 user_id
    2: InterviewType type     // 类型：专项/综合
    3: string category        // 专项类型或岗位
    4: InterviewRound round   // 面试轮次（综合面试）
    5: InterviewStatus status // 状态：进行中/已完成/已取消
    6: i32 score              // 评分
    7: string evaluation      // AI评估
    8: i64 created_at
    9: i64 finished_at
}

// 面试类型枚举
enum InterviewType {
    IT_NOT_SPECIFIED = 0  // 未指定
    SPECIALIZED = 1       // 专项面试
    COMPREHENSIVE = 2     // 综合面试
}

// 面试状态枚举
enum InterviewStatus {
    IS_NOT_SPECIFIED = 0  // 未指定
    IN_PROGRESS = 1       // 进行中
    COMPLETED = 2         // 已完成
    CANCELLED = 3         // 已取消
}

// 面试轮次枚举（综合面试）
enum InterviewRound {
    IR_NOT_SPECIFIED = 0  // 未指定
    FIRST_ROUND = 1       // 一面（技术基础）
    SECOND_ROUND = 2      // 二面（技术深度）
    THIRD_ROUND = 3       // 三面（综合能力）
    HR_ROUND = 4          // HR面
}

// 面试对话消息
struct InterviewMessage {
    1: i64 id
    2: i64 interview_id
    3: MessageRole role       // 消息角色：面试官/候选人
    4: string content
    5: i64 created_at
}

// 消息角色枚举
enum MessageRole {
    MR_NOT_SPECIFIED = 0  // 未指定
    INTERVIEWER = 1       // 面试官
    CANDIDATE = 2         // 候选人
}

// 能力评分
struct AbilityScore {
    1: string dimension       // 能力维度（如：编程基础、算法、系统设计等）
    2: i32 score              // 分数（0-100）
}