namespace go question

include "../base/common.thrift"
include "../base/question.thrift"

// ==================== 获取分类列表 ====================
struct GetCategoriesRequest {}

struct GetCategoriesResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: list<question.CategoryInfo> categories (api.body="categories"),
}

// ==================== 获取题目列表 ====================
struct GetQuestionListRequest {
    1: string category (api.query="category"),
    2: i32 difficulty (api.query="difficulty"),
    3: string keyword (api.query="keyword"),
    4: i32 page (api.query="page"),
    5: i32 page_size (api.query="page_size"),
}

struct GetQuestionListResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: list<question.QuestionEntity> questions (api.body="questions"),
    4: common.PageResponse page (api.body="page"),
}

// ==================== 获取题目详情 ====================
struct GetQuestionDetailResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: question.QuestionEntity question_entity (api.body="question"),
}

// ==================== 收藏题目 ====================
struct FavoriteQuestionRequest {
    1: i64 question_id (api.body="question_id", api.vd="$ > 0"),
}

// ==================== 获取收藏列表 ====================
struct GetFavoriteQuestionsRequest {
    1: i32 page (api.query="page"),
    2: i32 page_size (api.query="page_size"),
}

struct GetFavoriteQuestionsResponse {
    1: i64 status_code (api.body="code"),
    2: string status_msg (api.body="message"),
    3: list<question.QuestionEntity> questions (api.body="questions"),
    4: common.PageResponse page (api.body="page"),
}

// ==================== 题库服务接口 ====================
service QuestionService {
    GetCategoriesResponse GetCategories(1: GetCategoriesRequest req) (api.get="/api/v1/question/categories"),
    GetQuestionListResponse GetQuestionList(1: GetQuestionListRequest req) (api.get="/api/v1/question/list"),
    GetQuestionDetailResponse GetQuestionDetail(1: i64 question_id) (api.get="/api/v1/question/:question_id", api.path="question_id"),
    common.NilResponse FavoriteQuestion(1: FavoriteQuestionRequest req) (api.post="/api/v1/question/favorite"),
    common.NilResponse UnfavoriteQuestion(1: i64 question_id) (api.delete="/api/v1/question/favorite/:question_id", api.path="question_id"),GetFavoriteQuestionsResponse GetFavoriteQuestions(1: GetFavoriteQuestionsRequest req) (api.get="/api/v1/question/favorites"),
}