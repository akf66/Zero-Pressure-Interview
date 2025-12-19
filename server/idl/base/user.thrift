namespace go base

// 用户实体（带ID）
struct UserEntity {
    1: i64 id
    2: User user
}

// 用户领域对象
struct User {
    1: string email
    2: string nickname
    3: string avatar_url
    4: i64 created_at    // 创建时间戳（秒）
}

// 用户状态枚举
enum UserStatus {
    US_NOT_SPECIFIED = 0  // 未指定（必须有0值）
    ACTIVE = 1            // 激活
    INACTIVE = 2          // 未激活
    BANNED = 3            // 封禁
}

// 简历实体
struct ResumeEntity {
    1: i64 id
    2: Resume resume
}

// 简历领域对象
struct Resume {
    1: i64 user_id
    2: string file_url
    3: string parsed_content  // JSON格式的解析内容
    4: i64 created_at
}