package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"
	"zpi/server/shared/dal/sqlentity"
	"zpi/server/shared/dal/sqlfunc"
	"zpi/server/shared/errno"
	"zpi/server/shared/kitex_gen/base"
	"zpi/server/shared/kitex_gen/question"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// QuestionServiceImpl implements the last service interface defined in the IDL.
type QuestionServiceImpl struct {
	*QuestionManager
}

type QuestionManager struct {
	Query *sqlfunc.Query
}

// HealthCheck 健康检查
func (s *QuestionServiceImpl) HealthCheck(ctx context.Context) (resp *base.HealthCheckResponse, err error) {
	return &base.HealthCheckResponse{
		Status:  "ok",
		Version: "1.0.0",
	}, nil
}

// CreateQuestion 创建题目
func (s *QuestionServiceImpl) CreateQuestion(ctx context.Context, req *question.CreateQuestionRequest) (resp *question.CreateQuestionResponse, err error) {
	resp = &question.CreateQuestionResponse{}

	// 参数验证
	if req.Title == "" || req.Content == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("title and content are required"))
		return resp, nil
	}
	if req.Category == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("category is required"))
		return resp, nil
	}

	// 序列化标签
	var tagsJSON *string
	if len(req.Tags) > 0 {
		tagsBytes, _ := json.Marshal(req.Tags)
		tagsStr := string(tagsBytes)
		tagsJSON = &tagsStr
	}

	q := s.Query.Question

	// 创建题目记录
	questionRecord := &sqlentity.Question{
		Title:      req.Title,
		Content:    req.Content,
		Type:       req.Category,
		Category:   &req.Category,
		Difficulty: getDifficultyString(req.Difficulty),
		Tags:       tagsJSON,
		Answer:     &req.Answer,
		Status:     1,
	}

	if err := q.WithContext(ctx).Create(questionRecord); err != nil {
		klog.CtxErrorf(ctx, "create question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr.WithMessage("create question failed"))
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.QuestionId = int64(questionRecord.ID)
	return resp, nil
}

// GetQuestion 获取题目详情
func (s *QuestionServiceImpl) GetQuestion(ctx context.Context, req *question.GetQuestionRequest) (resp *question.GetQuestionResponse, err error) {
	resp = &question.GetQuestionResponse{}

	// 参数验证
	if req.QuestionId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid question_id"))
		return resp, nil
	}

	// 查询题目
	q := s.Query.Question
	questionRecord, err := q.WithContext(ctx).Where(q.ID.Eq(uint64(req.QuestionId)), q.Status.Eq(1)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.QuestionNotFound)
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.QuestionEntity = convertToQuestionEntity(questionRecord)
	return resp, nil
}

// GetQuestionList 获取题目列表
func (s *QuestionServiceImpl) GetQuestionList(ctx context.Context, req *question.GetQuestionListRequest) (resp *question.GetQuestionListResponse, err error) {
	resp = &question.GetQuestionListResponse{}

	// 分页参数
	page := int32(1)
	pageSize := int32(10)
	if req.Page != nil {
		if req.Page.Page > 0 {
			page = req.Page.Page
		}
		if req.Page.PageSize > 0 {
			pageSize = req.Page.PageSize
		}
	}

	q := s.Query.Question

	// 构建查询
	query := q.WithContext(ctx).Where(q.Status.Eq(1))

	// 分类筛选
	if req.Category != nil && *req.Category != "" {
		query = query.Where(q.Category.Eq(*req.Category))
	}

	// 难度筛选
	if req.Difficulty != nil && *req.Difficulty != base.QuestionDifficulty_QD_NOT_SPECIFIED {
		query = query.Where(q.Difficulty.Eq(getDifficultyString(*req.Difficulty)))
	}

	// 关键词搜索
	if req.Keyword != nil && *req.Keyword != "" {
		keyword := "%" + *req.Keyword + "%"
		query = query.Where(q.Where(q.Title.Like(keyword)).Or(q.Content.Like(keyword)))
	}

	// 标签筛选 - 暂时跳过，因为 gorm gen 不支持 JSON_CONTAINS
	// 如果需要标签筛选，可以在后续使用原生 SQL
	_ = req.Tags

	// 查询总数
	total, err := query.Count()
	if err != nil {
		klog.CtxErrorf(ctx, "count questions failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 分页查询
	offset := int((page - 1) * pageSize)
	questions, err := query.Order(q.CreatedAt.Desc()).Offset(offset).Limit(int(pageSize)).Find()
	if err != nil {
		klog.CtxErrorf(ctx, "query questions failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 转换响应
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Questions = make([]*base.QuestionEntity, len(questions))
	for i, question := range questions {
		resp.Questions[i] = convertToQuestionEntity(question)
	}
	resp.Page = &base.PageResponse{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	return resp, nil
}

// UpdateQuestion 更新题目
func (s *QuestionServiceImpl) UpdateQuestion(ctx context.Context, req *question.UpdateQuestionRequest) (resp *question.UpdateQuestionResponse, err error) {
	resp = &question.UpdateQuestionResponse{}

	// 参数验证
	if req.QuestionId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid question_id"))
		return resp, nil
	}

	q := s.Query.Question

	// 查询题目是否存在
	questionRecord, err := q.WithContext(ctx).Where(q.ID.Eq(uint64(req.QuestionId))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.QuestionNotFound)
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.Category != "" {
		updates["category"] = req.Category
		updates["type"] = req.Category
	}
	if req.Difficulty != base.QuestionDifficulty_QD_NOT_SPECIFIED {
		updates["difficulty"] = getDifficultyString(req.Difficulty)
	}
	if req.Answer != "" {
		updates["answer"] = req.Answer
	}
	if len(req.Tags) > 0 {
		tagsBytes, _ := json.Marshal(req.Tags)
		updates["tags"] = string(tagsBytes)
	}

	// 更新
	if len(updates) > 0 {
		if _, err := q.WithContext(ctx).Where(q.ID.Eq(questionRecord.ID)).Updates(updates); err != nil {
			klog.CtxErrorf(ctx, "update question failed: %v", err)
			resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
			return resp, nil
		}
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteQuestion 删除题目
func (s *QuestionServiceImpl) DeleteQuestion(ctx context.Context, req *question.DeleteQuestionRequest) (resp *question.DeleteQuestionResponse, err error) {
	resp = &question.DeleteQuestionResponse{}

	// 参数验证
	if req.QuestionId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid question_id"))
		return resp, nil
	}

	q := s.Query.Question

	// 软删除
	result, err := q.WithContext(ctx).Where(q.ID.Eq(uint64(req.QuestionId))).Delete()
	if err != nil {
		klog.CtxErrorf(ctx, "delete question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}
	if result.RowsAffected == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.QuestionNotFound)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetCategories 获取分类列表
func (s *QuestionServiceImpl) GetCategories(ctx context.Context, req *question.GetCategoriesRequest) (resp *question.GetCategoriesResponse, err error) {
	resp = &question.GetCategoriesResponse{}

	q := s.Query.Question

	// 查询分类及数量 - 使用原生 SQL
	type CategoryCount struct {
		Category string
		Count    int32
	}
	var results []CategoryCount

	err = q.WithContext(ctx).
		Select(q.Category, q.ID.Count().As("count")).
		Where(q.Status.Eq(1), q.Category.IsNotNull()).
		Group(q.Category).
		Scan(&results)

	if err != nil {
		klog.CtxErrorf(ctx, "query categories failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 转换响应
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Categories = make([]*base.CategoryInfo, len(results))
	for i, r := range results {
		resp.Categories[i] = &base.CategoryInfo{
			Name:  r.Category,
			Count: r.Count,
		}
	}
	return resp, nil
}

// GetRandomQuestions 随机获取题目
func (s *QuestionServiceImpl) GetRandomQuestions(ctx context.Context, req *question.GetRandomQuestionsRequest) (resp *question.GetRandomQuestionsResponse, err error) {
	resp = &question.GetRandomQuestionsResponse{}

	// 参数验证
	count := req.Count
	if count <= 0 {
		count = 5
	}
	if count > 20 {
		count = 20
	}

	q := s.Query.Question

	// 构建查询
	query := q.WithContext(ctx).Where(q.Status.Eq(1))

	// 分类筛选
	if req.Category != "" {
		query = query.Where(q.Category.Eq(req.Category))
	}

	// 难度筛选
	if req.Difficulty != base.QuestionDifficulty_QD_NOT_SPECIFIED {
		query = query.Where(q.Difficulty.Eq(getDifficultyString(req.Difficulty)))
	}

	// 获取符合条件的题目ID
	var questionIDs []uint64
	err = query.Pluck(q.ID, &questionIDs)
	if err != nil {
		klog.CtxErrorf(ctx, "query question ids failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	if len(questionIDs) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.Success)
		resp.Questions = []*base.QuestionEntity{}
		return resp, nil
	}

	// 随机选择
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(questionIDs), func(i, j int) {
		questionIDs[i], questionIDs[j] = questionIDs[j], questionIDs[i]
	})

	if int(count) > len(questionIDs) {
		count = int32(len(questionIDs))
	}
	selectedIDs := questionIDs[:count]

	// 查询选中的题目
	questions, err := q.WithContext(ctx).Where(q.ID.In(selectedIDs...)).Find()
	if err != nil {
		klog.CtxErrorf(ctx, "query questions failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 转换响应
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Questions = make([]*base.QuestionEntity, len(questions))
	for i, question := range questions {
		resp.Questions[i] = convertToQuestionEntity(question)
	}
	return resp, nil
}

// FavoriteQuestion 收藏题目
func (s *QuestionServiceImpl) FavoriteQuestion(ctx context.Context, req *question.FavoriteQuestionRequest) (resp *question.FavoriteQuestionResponse, err error) {
	resp = &question.FavoriteQuestionResponse{}

	// 参数验证
	if req.UserId <= 0 || req.QuestionId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}

	q := s.Query.Question

	// 检查题目是否存在
	_, err = q.WithContext(ctx).Where(q.ID.Eq(uint64(req.QuestionId)), q.Status.Eq(1)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.QuestionNotFound)
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	uf := s.Query.UserFavorite

	// 检查是否已收藏
	existing, _ := uf.WithContext(ctx).Where(uf.UserID.Eq(uint64(req.UserId)), uf.QuestionID.Eq(uint64(req.QuestionId))).First()
	if existing != nil {
		// 已经收藏，直接返回成功
		resp.BaseResp = errno.BuildBaseResp(errno.Success)
		return resp, nil
	}

	// 创建收藏记录
	favorite := &sqlentity.UserFavorite{
		UserID:     uint64(req.UserId),
		QuestionID: uint64(req.QuestionId),
	}
	if err := uf.WithContext(ctx).Create(favorite); err != nil {
		klog.CtxErrorf(ctx, "create favorite failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// UnfavoriteQuestion 取消收藏
func (s *QuestionServiceImpl) UnfavoriteQuestion(ctx context.Context, req *question.UnfavoriteQuestionRequest) (resp *question.UnfavoriteQuestionResponse, err error) {
	resp = &question.UnfavoriteQuestionResponse{}

	// 参数验证
	if req.UserId <= 0 || req.QuestionId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}

	uf := s.Query.UserFavorite

	// 删除收藏记录
	_, err = uf.WithContext(ctx).Where(uf.UserID.Eq(uint64(req.UserId)), uf.QuestionID.Eq(uint64(req.QuestionId))).Delete()
	if err != nil {
		klog.CtxErrorf(ctx, "delete favorite failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteQuestions 获取收藏列表
func (s *QuestionServiceImpl) GetFavoriteQuestions(ctx context.Context, req *question.GetFavoriteQuestionsRequest) (resp *question.GetFavoriteQuestionsResponse, err error) {
	resp = &question.GetFavoriteQuestionsResponse{}

	// 参数验证
	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid user_id"))
		return resp, nil
	}

	// 分页参数
	page := int32(1)
	pageSize := int32(10)
	if req.Page != nil {
		if req.Page.Page > 0 {
			page = req.Page.Page
		}
		if req.Page.PageSize > 0 {
			pageSize = req.Page.PageSize
		}
	}

	uf := s.Query.UserFavorite

	// 查询收藏的题目ID
	var favoriteIDs []uint64
	err = uf.WithContext(ctx).Where(uf.UserID.Eq(uint64(req.UserId))).Pluck(uf.QuestionID, &favoriteIDs)
	if err != nil {
		klog.CtxErrorf(ctx, "query favorite ids failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	if len(favoriteIDs) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.Success)
		resp.Questions = []*base.QuestionEntity{}
		resp.Page = &base.PageResponse{Total: 0, Page: page, PageSize: pageSize}
		return resp, nil
	}

	// 查询总数
	total := int64(len(favoriteIDs))

	q := s.Query.Question

	// 分页查询题目
	offset := int((page - 1) * pageSize)
	questions, err := q.WithContext(ctx).
		Where(q.ID.In(favoriteIDs...), q.Status.Eq(1)).
		Order(q.CreatedAt.Desc()).
		Offset(offset).
		Limit(int(pageSize)).
		Find()
	if err != nil {
		klog.CtxErrorf(ctx, "query questions failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 转换响应
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Questions = make([]*base.QuestionEntity, len(questions))
	for i, question := range questions {
		resp.Questions[i] = convertToQuestionEntity(question)
	}
	resp.Page = &base.PageResponse{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	return resp, nil
}

// AddQuestionNote 添加笔记
func (s *QuestionServiceImpl) AddQuestionNote(ctx context.Context, req *question.AddQuestionNoteRequest) (resp *question.AddQuestionNoteResponse, err error) {
	resp = &question.AddQuestionNoteResponse{}

	// 参数验证
	if req.UserId <= 0 || req.QuestionId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}
	if req.Note == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("note is required"))
		return resp, nil
	}

	q := s.Query.Question

	// 检查题目是否存在
	_, err = q.WithContext(ctx).Where(q.ID.Eq(uint64(req.QuestionId)), q.Status.Eq(1)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.QuestionNotFound)
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	un := s.Query.UserNote

	// 检查笔记是否已存在
	existingNote, _ := un.WithContext(ctx).Where(un.UserID.Eq(uint64(req.UserId)), un.QuestionID.Eq(uint64(req.QuestionId))).First()

	if existingNote != nil {
		// 更新现有笔记
		if _, err := un.WithContext(ctx).Where(un.ID.Eq(existingNote.ID)).Update(un.Note, req.Note); err != nil {
			klog.CtxErrorf(ctx, "update note failed: %v", err)
			resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
			return resp, nil
		}
	} else {
		// 创建新笔记
		note := &sqlentity.UserNote{
			UserID:     uint64(req.UserId),
			QuestionID: uint64(req.QuestionId),
			Note:       req.Note,
		}
		if err := un.WithContext(ctx).Create(note); err != nil {
			klog.CtxErrorf(ctx, "create note failed: %v", err)
			resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
			return resp, nil
		}
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetQuestionNote 获取笔记
func (s *QuestionServiceImpl) GetQuestionNote(ctx context.Context, req *question.GetQuestionNoteRequest) (resp *question.GetQuestionNoteResponse, err error) {
	resp = &question.GetQuestionNoteResponse{}

	// 参数验证
	if req.UserId <= 0 || req.QuestionId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}

	un := s.Query.UserNote

	// 查询笔记
	note, err := un.WithContext(ctx).Where(un.UserID.Eq(uint64(req.UserId)), un.QuestionID.Eq(uint64(req.QuestionId))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.Success)
			resp.Note = ""
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query note failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Note = note.Note
	return resp, nil
}

// 辅助函数

func getDifficultyString(d base.QuestionDifficulty) string {
	switch d {
	case base.QuestionDifficulty_EASY:
		return "easy"
	case base.QuestionDifficulty_MEDIUM:
		return "medium"
	case base.QuestionDifficulty_HARD:
		return "hard"
	default:
		return "medium"
	}
}

func getDifficultyEnum(d string) int32 {
	switch d {
	case "easy":
		return 1
	case "medium":
		return 2
	case "hard":
		return 3
	default:
		return 2
	}
}

func convertToQuestionEntity(q *sqlentity.Question) *base.QuestionEntity {
	entity := &base.QuestionEntity{
		Id: int64(q.ID),
		Question: &base.Question{
			Title:      q.Title,
			Content:    q.Content,
			Difficulty: getDifficultyEnum(q.Difficulty),
			CreatedAt:  q.CreatedAt.Unix(),
		},
	}

	if q.Category != nil {
		entity.Question.Category = *q.Category
	}
	if q.Answer != nil {
		entity.Question.Answer = *q.Answer
	}

	// 解析标签
	if q.Tags != nil && *q.Tags != "" {
		var tags []string
		if err := json.Unmarshal([]byte(*q.Tags), &tags); err == nil {
			entity.Question.Tags = tags
		}
	}

	return entity
}
