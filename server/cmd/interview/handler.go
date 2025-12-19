package main

import (
	"context"
	interview "zpi/server/shared/kitex_gen/interview"
)

// InterviewServiceImpl implements the last service interface defined in the IDL.
type InterviewServiceImpl struct{}

// StartInterview implements the InterviewServiceImpl interface.
func (s *InterviewServiceImpl) StartInterview(ctx context.Context, req *interview.StartInterviewRequest) (resp *interview.StartInterviewResponse, err error) {
	// TODO: Your code here...
	return
}

// SubmitAnswer implements the InterviewServiceImpl interface.
func (s *InterviewServiceImpl) SubmitAnswer(ctx context.Context, req *interview.SubmitAnswerRequest) (resp *interview.SubmitAnswerResponse, err error) {
	// TODO: Your code here...
	return
}

// FinishInterview implements the InterviewServiceImpl interface.
func (s *InterviewServiceImpl) FinishInterview(ctx context.Context, req *interview.FinishInterviewRequest) (resp *interview.FinishInterviewResponse, err error) {
	// TODO: Your code here...
	return
}

// GetInterviewDetail implements the InterviewServiceImpl interface.
func (s *InterviewServiceImpl) GetInterviewDetail(ctx context.Context, req *interview.GetInterviewDetailRequest) (resp *interview.GetInterviewDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// GetInterviewHistory implements the InterviewServiceImpl interface.
func (s *InterviewServiceImpl) GetInterviewHistory(ctx context.Context, req *interview.GetInterviewHistoryRequest) (resp *interview.GetInterviewHistoryResponse, err error) {
	// TODO: Your code here...
	return
}

// AnalyzeResume implements the InterviewServiceImpl interface.
func (s *InterviewServiceImpl) AnalyzeResume(ctx context.Context, req *interview.AnalyzeResumeRequest) (resp *interview.AnalyzeResumeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAbilityAnalysis implements the InterviewServiceImpl interface.
func (s *InterviewServiceImpl) GetAbilityAnalysis(ctx context.Context, req *interview.GetAbilityAnalysisRequest) (resp *interview.GetAbilityAnalysisResponse, err error) {
	// TODO: Your code here...
	return
}
