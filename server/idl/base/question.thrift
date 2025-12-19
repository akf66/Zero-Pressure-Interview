namespace go base

// 题目实体
struct QuestionEntity {
    1: i64 id
    2: Question question
}

// 题目领域对象
struct Question {
    1: string category        // 分类：golang/java/mysql等
    2: i32 difficulty         // 难度：1-简单/2-中等/3-困难
    3: string title
    4: string content
    5: string answer
    6: list<string> tags      // 标签列表
    7: i64 created_at
}

// 题目难度枚举
enum QuestionDifficulty {
    QD_NOT_SPECIFIED = 0  // 未指定EASY = 1              // 简单
    MEDIUM = 2            // 中等
    HARD = 3              // 困难
}

// 分类信息
struct CategoryInfo {
    1: string name// 分类名称
    2: i32 count              // 题目数量
}