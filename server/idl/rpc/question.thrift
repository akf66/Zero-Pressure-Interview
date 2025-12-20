namespace go question

include "../base/common.thrift"
include "../base/question.thrift"

// ==================== 创建题目 ====================
struct CreateQuestionRequest {
    1: string category                    // 分类：golang/java/mysql等
    2: question.QuestionDifficulty difficulty  // 难度：简单/中等/困难
    3: string title
    4: string content
    5: string answer
    6: list<string> tags
}

struct CreateQuestionResponse {
    1: common.BaseResponse base_resp
    2: i64 question_id
}

// ==================== 获取题目详情 ====================
struct GetQuestionRequest {
    1: i64 question_id
}

struct GetQuestionResponse {
    1: common.BaseResponse base_resp
    2: question.QuestionEntity question_entity
}

// ==================== 获取题目列表 ====================
struct GetQuestionListRequest {
    1: optional string category                    // 分类筛选
    2: optional question.QuestionDifficulty difficulty  // 难度筛选
    3: optional list<string> tags                  // 标签筛选
    4: optional string keyword                     // 关键词搜索
    5: common.PageRequest page
}

struct GetQuestionListResponse {
    1: common.BaseResponse base_resp
    2: list<question.QuestionEntity> questions
    3: common.PageResponse page
}

// ==================== 更新题目 ====================
struct UpdateQuestionRequest {
    1: i64 question_id
    2: string category
    3: question.QuestionDifficulty difficulty
    4: string title
    5: string content
    6: string answer
    7: list<string> tags
}

struct UpdateQuestionResponse {
    1: common.BaseResponse base_resp
}

// ==================== 删除题目 ====================
struct DeleteQuestionRequest {
    1: i64 question_id
}

struct DeleteQuestionResponse {
    1: common.BaseResponse base_resp
}

// ==================== 获取分类列表 ====================
struct GetCategoriesRequest {}

struct GetCategoriesResponse {
    1: common.BaseResponse base_resp
    2: list<question.CategoryInfo> categories
}

// ==================== 随机获取题目 ====================
struct GetRandomQuestionsRequest {
    1: string category                         // 分类
    2: question.QuestionDifficulty difficulty  // 难度
    3: i32 count                               // 数量
}

struct GetRandomQuestionsResponse {
    1: common.BaseResponse base_resp
    2: list<question.QuestionEntity> questions
}

// ==================== 收藏题目 ====================
struct FavoriteQuestionRequest {
    1: i64 user_id
    2: i64 question_id
    3: optional string idempotency_key  // 幂等性key，防止重复收藏
}

struct FavoriteQuestionResponse {
    1: common.BaseResponse base_resp
}

// ==================== 取消收藏 ====================
struct UnfavoriteQuestionRequest {
    1: i64 user_id
    2: i64 question_id
}

struct UnfavoriteQuestionResponse {
    1: common.BaseResponse base_resp
}

// ==================== 获取收藏列表 ====================
struct GetFavoriteQuestionsRequest {
    1: i64 user_id
    2: common.PageRequest page
}

struct GetFavoriteQuestionsResponse {
    1: common.BaseResponse base_resp
    2: list<question.QuestionEntity> questions
    3: common.PageResponse page
}

// ==================== 添加用户笔记 ====================
struct AddQuestionNoteRequest {
    1: i64 user_id
    2: i64 question_id
    3: string note
}

struct AddQuestionNoteResponse {
    1: common.BaseResponse base_resp
}

// ==================== 获取用户笔记 ====================
struct GetQuestionNoteRequest {
    1: i64 user_id
    2: i64 question_id
}

struct GetQuestionNoteResponse {
    1: common.BaseResponse base_resp
    2: string note
}

// ==================== Question服务接口 ====================
service QuestionService {
    // 健康检查
    common.HealthCheckResponse HealthCheck()
    
    // 题目管理（管理员）
    CreateQuestionResponse CreateQuestion(1: CreateQuestionRequest req)
    UpdateQuestionResponse UpdateQuestion(1: UpdateQuestionRequest req)
    DeleteQuestionResponse DeleteQuestion(1: DeleteQuestionRequest req)
    // 题目查询
    GetQuestionResponse GetQuestion(1: GetQuestionRequest req)
    GetQuestionListResponse GetQuestionList(1: GetQuestionListRequest req)
    GetCategoriesResponse GetCategories(1: GetCategoriesRequest req)
    GetRandomQuestionsResponse GetRandomQuestions(1: GetRandomQuestionsRequest req)
    
    // 用户收藏和笔记
    FavoriteQuestionResponse FavoriteQuestion(1: FavoriteQuestionRequest req)
    UnfavoriteQuestionResponse UnfavoriteQuestion(1: UnfavoriteQuestionRequest req)
    GetFavoriteQuestionsResponse GetFavoriteQuestions(1: GetFavoriteQuestionsRequest req)
    AddQuestionNoteResponse AddQuestionNote(1: AddQuestionNoteRequest req)
    GetQuestionNoteResponse GetQuestionNote(1: GetQuestionNoteRequest req)
}