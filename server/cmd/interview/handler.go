package main

import (
	"context"
	"time"
	"zpi/server/cmd/interview/config"
	"zpi/server/cmd/interview/pkg/ai"
	"zpi/server/shared/dal/sqlentity"
	"zpi/server/shared/dal/sqlfunc"
	"zpi/server/shared/errno"
	"zpi/server/shared/kitex_gen/base"
	"zpi/server/shared/kitex_gen/interview"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// InterviewServiceImpl implements the last service interface defined in the IDL.
type InterviewServiceImpl struct {
	*InterviewManager
}

type InterviewManager struct {
	Query *sqlfunc.Query
}

// HealthCheck 健康检查
func (s *InterviewServiceImpl) HealthCheck(ctx context.Context) (resp *base.HealthCheckResponse, err error) {
	return &base.HealthCheckResponse{
		Status:  "ok",
		Version: "1.0.0",
	}, nil
}

// StartInterview 开始面试
func (s *InterviewServiceImpl) StartInterview(ctx context.Context, req *interview.StartInterviewRequest) (resp *interview.StartInterviewResponse, err error) {
	resp = &interview.StartInterviewResponse{}

	// 参数验证
	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid user_id"))
		return resp, nil
	}
	if req.Category == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("category is required"))
		return resp, nil
	}

	// 创建面试记录
	now := time.Now()
	interviewRecord := &sqlentity.Interview{
		UserID:     uint64(req.UserId),
		Title:      generateInterviewTitle(req.Type, req.Category),
		Type:       getInterviewTypeString(req.Type),
		Difficulty: getDifficultyString(req.Difficulty),
		Status:     "in_progress",
		StartedAt:  &now,
	}

	if err := config.DB.Create(interviewRecord).Error; err != nil {
		klog.CtxErrorf(ctx, "create interview failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr.WithMessage("create interview failed"))
		return resp, nil
	}

	// 生成第一个问题
	aiClient := ai.GetClient()
	firstQuestion, err := aiClient.GenerateFirstQuestion(ctx, req.Category, req.Difficulty, "")
	if err != nil {
		klog.CtxErrorf(ctx, "generate first question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr.WithMessage("generate question failed"))
		return resp, nil
	}

	// 保存面试官的第一个问题
	message := &sqlentity.InterviewMessage{
		InterviewID: interviewRecord.ID,
		Role:        "interviewer",
		Content:     firstQuestion,
	}
	if err := config.DB.Create(message).Error; err != nil {
		klog.CtxErrorf(ctx, "save message failed: %v", err)
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.InterviewId = int64(interviewRecord.ID)
	resp.FirstQuestion = firstQuestion
	return resp, nil
}

// SubmitAnswer 提交回答
func (s *InterviewServiceImpl) SubmitAnswer(ctx context.Context, req *interview.SubmitAnswerRequest) (resp *interview.SubmitAnswerResponse, err error) {
	resp = &interview.SubmitAnswerResponse{}

	// 参数验证
	if req.InterviewId <= 0 || req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}
	if req.Answer == "" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("answer is required"))
		return resp, nil
	}

	// 查询面试记录
	var interviewRecord sqlentity.Interview
	if err := config.DB.Where("id = ? AND user_id = ?", req.InterviewId, req.UserId).First(&interviewRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.RecordNotFound.WithMessage("interview not found"))
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query interview failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 检查面试状态
	if interviewRecord.Status != "in_progress" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("interview is not in progress"))
		return resp, nil
	}

	// 保存用户回答
	candidateMsg := &sqlentity.InterviewMessage{
		InterviewID: uint64(req.InterviewId),
		Role:        "candidate",
		Content:     req.Answer,
	}
	if err := config.DB.Create(candidateMsg).Error; err != nil {
		klog.CtxErrorf(ctx, "save answer failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 获取对话历史
	var messages []sqlentity.InterviewMessage
	config.DB.Where("interview_id = ?", req.InterviewId).Order("created_at ASC").Find(&messages)

	// 转换为AI消息格式
	history := make([]ai.Message, len(messages))
	for i, msg := range messages {
		history[i] = ai.Message{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	// 生成下一个问题
	aiClient := ai.GetClient()
	nextQuestion, isFinished, err := aiClient.GenerateNextQuestion(ctx, interviewRecord.Type, history)
	if err != nil {
		klog.CtxErrorf(ctx, "generate next question failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 如果面试未结束，保存面试官的问题
	if !isFinished && nextQuestion != "" {
		interviewerMsg := &sqlentity.InterviewMessage{
			InterviewID: uint64(req.InterviewId),
			Role:        "interviewer",
			Content:     nextQuestion,
		}
		if err := config.DB.Create(interviewerMsg).Error; err != nil {
			klog.CtxErrorf(ctx, "save interviewer message failed: %v", err)
		}
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.NextQuestion = nextQuestion
	resp.IsFinished = isFinished
	return resp, nil
}

// FinishInterview 结束面试
func (s *InterviewServiceImpl) FinishInterview(ctx context.Context, req *interview.FinishInterviewRequest) (resp *interview.FinishInterviewResponse, err error) {
	resp = &interview.FinishInterviewResponse{}

	// 参数验证
	if req.InterviewId <= 0 || req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}

	// 查询面试记录
	var interviewRecord sqlentity.Interview
	if err := config.DB.Where("id = ? AND user_id = ?", req.InterviewId, req.UserId).First(&interviewRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.RecordNotFound.WithMessage("interview not found"))
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query interview failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 检查面试状态
	if interviewRecord.Status == "completed" {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("interview already completed"))
		return resp, nil
	}

	// 获取对话历史
	var messages []sqlentity.InterviewMessage
	config.DB.Where("interview_id = ?", req.InterviewId).Order("created_at ASC").Find(&messages)

	// 转换为AI消息格式
	history := make([]ai.Message, len(messages))
	for i, msg := range messages {
		history[i] = ai.Message{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	// AI评估
	aiClient := ai.GetClient()
	score, evaluation, err := aiClient.EvaluateInterview(ctx, interviewRecord.Type, history)
	if err != nil {
		klog.CtxErrorf(ctx, "evaluate interview failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 更新面试记录
	now := time.Now()
	scoreFloat := float64(score)
	updates := map[string]interface{}{
		"status":       "completed",
		"completed_at": now,
		"score":        scoreFloat,
		"feedback":     evaluation,
	}
	if interviewRecord.StartedAt != nil {
		duration := int32(now.Sub(*interviewRecord.StartedAt).Seconds())
		updates["duration"] = duration
	}

	if err := config.DB.Model(&interviewRecord).Updates(updates).Error; err != nil {
		klog.CtxErrorf(ctx, "update interview failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Score = score
	resp.Evaluation = evaluation
	return resp, nil
}

// GetInterviewDetail 获取面试详情
func (s *InterviewServiceImpl) GetInterviewDetail(ctx context.Context, req *interview.GetInterviewDetailRequest) (resp *interview.GetInterviewDetailResponse, err error) {
	resp = &interview.GetInterviewDetailResponse{}

	// 参数验证
	if req.InterviewId <= 0 || req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}

	// 查询面试记录
	var interviewRecord sqlentity.Interview
	if err := config.DB.Where("id = ? AND user_id = ?", req.InterviewId, req.UserId).First(&interviewRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = errno.BuildBaseResp(errno.RecordNotFound.WithMessage("interview not found"))
			return resp, nil
		}
		klog.CtxErrorf(ctx, "query interview failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	// 查询对话记录
	var messages []sqlentity.InterviewMessage
	config.DB.Where("interview_id = ?", req.InterviewId).Order("created_at ASC").Find(&messages)

	// 转换响应
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.InterviewEntity = convertToInterviewEntity(&interviewRecord)
	resp.Messages = convertToInterviewMessages(messages)
	return resp, nil
}

// GetInterviewHistory 获取面试历史
func (s *InterviewServiceImpl) GetInterviewHistory(ctx context.Context, req *interview.GetInterviewHistoryRequest) (resp *interview.GetInterviewHistoryResponse, err error) {
	resp = &interview.GetInterviewHistoryResponse{}

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

	// 构建查询
	query := config.DB.Model(&sqlentity.Interview{}).Where("user_id = ?", req.UserId)

	// 类型筛选
	if req.Type != nil && *req.Type != base.InterviewType_IT_NOT_SPECIFIED {
		query = query.Where("type = ?", getInterviewTypeString(*req.Type))
	}

	// 查询总数
	var total int64
	query.Count(&total)

	// 分页查询
	var interviews []sqlentity.Interview
	offset := (page - 1) * pageSize
	query.Order("created_at DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&interviews)

	// 转换响应
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Interviews = make([]*base.InterviewEntity, len(interviews))
	for i, iv := range interviews {
		resp.Interviews[i] = convertToInterviewEntity(&iv)
	}
	resp.Page = &base.PageResponse{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	return resp, nil
}

// AnalyzeResume 分析简历
func (s *InterviewServiceImpl) AnalyzeResume(ctx context.Context, req *interview.AnalyzeResumeRequest) (resp *interview.AnalyzeResumeResponse, err error) {
	resp = &interview.AnalyzeResumeResponse{}

	// 参数验证
	if req.UserId <= 0 || req.ResumeId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid parameters"))
		return resp, nil
	}

	// TODO: 调用 User 服务获取简历内容
	resumeContent := "模拟简历内容"

	// AI分析简历
	aiClient := ai.GetClient()
	analysis, err := aiClient.AnalyzeResume(ctx, resumeContent)
	if err != nil {
		klog.CtxErrorf(ctx, "analyze resume failed: %v", err)
		resp.BaseResp = errno.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Analysis = analysis.Analysis
	resp.Suggestions = analysis.Suggestions
	resp.MatchedPositions = analysis.MatchedPositions
	return resp, nil
}

// GetAbilityAnalysis 获取能力分析
func (s *InterviewServiceImpl) GetAbilityAnalysis(ctx context.Context, req *interview.GetAbilityAnalysisRequest) (resp *interview.GetAbilityAnalysisResponse, err error) {
	resp = &interview.GetAbilityAnalysisResponse{}

	// 参数验证
	if req.UserId <= 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamsErr.WithMessage("invalid user_id"))
		return resp, nil
	}

	// 查询用户的所有已完成面试
	var interviews []sqlentity.Interview
	config.DB.Where("user_id = ? AND status = ?", req.UserId, "completed").Find(&interviews)

	if len(interviews) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.Success)
		resp.Summary = "暂无面试记录，无法生成能力分析"
		return resp, nil
	}

	// 计算各维度能力评分（模拟）
	abilities := []*base.AbilityScore{
		{Dimension: "编程基础", Score: 75},
		{Dimension: "算法能力", Score: 70},
		{Dimension: "系统设计", Score: 65},
		{Dimension: "沟通表达", Score: 80},
		{Dimension: "问题解决", Score: 72},
	}

	// 计算平均分
	var totalScore float64
	for _, iv := range interviews {
		if iv.Score != nil {
			totalScore += *iv.Score
		}
	}
	avgScore := totalScore / float64(len(interviews))

	summary := "根据您的面试表现，您在沟通表达方面表现突出，编程基础扎实。建议加强系统设计和算法方面的学习。"
	if avgScore >= 80 {
		summary = "您的整体表现优秀！各项能力均衡发展，建议继续保持并挑战更高难度的面试。"
	} else if avgScore < 60 {
		summary = "建议加强基础知识的学习，多进行模拟面试练习，提升面试技巧。"
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Abilities = abilities
	resp.Summary = summary
	return resp, nil
}

// 辅助函数

func generateInterviewTitle(interviewType base.InterviewType, category string) string {
	typeStr := "面试"
	if interviewType == base.InterviewType_SPECIALIZED {
		typeStr = "专项面试"
	} else if interviewType == base.InterviewType_COMPREHENSIVE {
		typeStr = "综合面试"
	}
	return category + " " + typeStr
}

func getInterviewTypeString(t base.InterviewType) string {
	switch t {
	case base.InterviewType_SPECIALIZED:
		return "specialized"
	case base.InterviewType_COMPREHENSIVE:
		return "comprehensive"
	default:
		return "technical"
	}
}

func getDifficultyString(d int32) string {
	switch d {
	case 1:
		return "easy"
	case 2:
		return "medium"
	case 3:
		return "hard"
	default:
		return "medium"
	}
}

func convertToInterviewEntity(iv *sqlentity.Interview) *base.InterviewEntity {
	entity := &base.InterviewEntity{
		Id: int64(iv.ID),
		Interview: &base.Interview{
			UserId:   int64(iv.UserID),
			Category: iv.Title,
			Status:   getInterviewStatus(iv.Status),
		},
	}

	if iv.Score != nil {
		entity.Interview.Score = int32(*iv.Score)
	}
	if iv.Feedback != nil {
		entity.Interview.Evaluation = *iv.Feedback
	}
	entity.Interview.CreatedAt = iv.CreatedAt.Unix()
	if iv.CompletedAt != nil {
		entity.Interview.FinishedAt = iv.CompletedAt.Unix()
	}

	return entity
}

func getInterviewStatus(status string) base.InterviewStatus {
	switch status {
	case "in_progress":
		return base.InterviewStatus_IN_PROGRESS
	case "completed":
		return base.InterviewStatus_COMPLETED
	case "cancelled":
		return base.InterviewStatus_CANCELLED
	default:
		return base.InterviewStatus_IS_NOT_SPECIFIED
	}
}

func convertToInterviewMessages(messages []sqlentity.InterviewMessage) []*base.InterviewMessage {
	result := make([]*base.InterviewMessage, len(messages))
	for i, msg := range messages {
		result[i] = &base.InterviewMessage{
			Id:          int64(msg.ID),
			InterviewId: int64(msg.InterviewID),
			Role:        getMessageRole(msg.Role),
			Content:     msg.Content,
			CreatedAt:   msg.CreatedAt.Unix(),
		}
	}
	return result
}

func getMessageRole(role string) base.MessageRole {
	switch role {
	case "interviewer":
		return base.MessageRole_INTERVIEWER
	case "candidate":
		return base.MessageRole_CANDIDATE
	default:
		return base.MessageRole_MR_NOT_SPECIFIED
	}
}
