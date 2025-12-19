namespace go question

include "../base/base.thrift"

// ==================== 创建题目 ====================
struct CreateQuestionReq {
    1: required string category       // 分类：golang/java/mysql等
    2: required i32 difficulty        // 难度：1-简单/2-中等/3-困难
    3: required string title
    4: required string content
    5: optional string answer
    6: optional list<string> tags
}

struct CreateQuestionResp {
    1: base.BaseResp base_resp
    2: optional i64 question_id
}

// ==================== 获取题目详情 ====================
struct GetQuestionReq {
    1: required i64 question_id
}

struct GetQuestionResp {
    1: base.BaseResp base_resp
    2: optional base.QuestionInfo question
}

// ==================== 获取题目列表 ====================
struct GetQuestionListReq {
    1: optional string category       // 分类筛选
    2: optional i32 difficulty        // 难度筛选
    3: optional list<string> tags     // 标签筛选
    4: optional string keyword        // 关键词搜索
    5: optional base.PageReq page
}

struct GetQuestionListResp {
    1: base.BaseResp base_resp
    2: optional list<base.QuestionInfo> questions
    3: optional base.PageResp page
}

// ==================== 更新题目 ====================
struct UpdateQuestionReq {
    1: required i64 question_id
    2: optional string category
    3: optional i32 difficulty
    4: optional string title
    5: optional string content
    6: optional string answer
    7: optional list<string> tags
}

struct UpdateQuestionResp {
    1: base.BaseResp base_resp
}

// ==================== 删除题目 ====================
struct DeleteQuestionReq {
    1: required i64 question_id
}

struct DeleteQuestionResp {
    1: base.BaseResp base_resp
}

// ==================== 获取分类列表 ====================
struct GetCategoriesReq {
}

struct CategoryInfo {
    1: string name    // 分类名称
    2: i32 count                      // 题目数量
}

struct GetCategoriesResp {
    1: base.BaseResp base_resp
    2: optional list<CategoryInfo> categories
}

// ==================== 随机获取题目 ====================
struct GetRandomQuestionsReq {
    1: required string category       // 分类
    2: optional i32 difficulty        // 难度
    3: required i32 count             // 数量
}

struct GetRandomQuestionsResp {
    1: base.BaseResp base_resp
    2: optional list<base.QuestionInfo> questions
}

// ==================== 收藏题目 ====================
struct FavoriteQuestionReq {
    1: required i64 user_id
    2: required i64 question_id
}

struct FavoriteQuestionResp {
    1: base.BaseResp base_resp
}

// ==================== 取消收藏 ====================
struct UnfavoriteQuestionReq {
    1: required i64 user_id
    2: required i64 question_id
}

struct UnfavoriteQuestionResp {
    1: base.BaseResp base_resp
}

// ==================== 获取收藏列表 ====================
struct GetFavoriteQuestionsReq {
    1: required i64 user_id
    2: optional base.PageReq page
}

struct GetFavoriteQuestionsResp {
    1: base.BaseResp base_resp
    2: optional list<base.QuestionInfo> questions
    3: optional base.PageResp page
}

// ==================== 添加用户笔记 ====================
struct AddQuestionNoteReq {
    1: required i64 user_id
    2: required i64 question_id
    3: required string note
}

struct AddQuestionNoteResp {
    1: base.BaseResp base_resp
}

// ==================== 获取用户笔记 ====================
struct GetQuestionNoteReq {
    1: required i64 user_id
    2: required i64 question_id
}

struct GetQuestionNoteResp {
    1: base.BaseResp base_resp
    2: optional string note
}

// ==================== Question服务接口 ====================
service QuestionService {
    // 创建题目
    CreateQuestionResp CreateQuestion(1: CreateQuestionReq req)
    
    // 获取题目详情
    GetQuestionResp GetQuestion(1: GetQuestionReq req)
    
    // 获取题目列表
    GetQuestionListResp GetQuestionList(1: GetQuestionListReq req)
    
    // 更新题目
    UpdateQuestionResp UpdateQuestion(1: UpdateQuestionReq req)
    
    // 删除题目
    DeleteQuestionResp DeleteQuestion(1: DeleteQuestionReq req)
    
    // 获取分类列表
    GetCategoriesResp GetCategories(1: GetCategoriesReq req)
    
    // 随机获取题目
    GetRandomQuestionsResp GetRandomQuestions(1: GetRandomQuestionsReq req)
    
    // 收藏题目
    FavoriteQuestionResp FavoriteQuestion(1: FavoriteQuestionReq req)
    
    // 取消收藏
    UnfavoriteQuestionResp UnfavoriteQuestion(1: UnfavoriteQuestionReq req)
    
    // 获取收藏列表
    GetFavoriteQuestionsResp GetFavoriteQuestions(1: GetFavoriteQuestionsReq req)
    
    // 添加用户笔记
    AddQuestionNoteResp AddQuestionNote(1: AddQuestionNoteReq req)
    
    // 获取用户笔记
    GetQuestionNoteResp GetQuestionNote(1: GetQuestionNoteReq req)
}