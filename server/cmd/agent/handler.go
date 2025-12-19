package main

import (
	"context"
	agent "zpi/server/shared/kitex_gen/agent"
)

// AgentServiceImpl implements the last service interface defined in the IDL.
type AgentServiceImpl struct{}

// StartInterview implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) StartInterview(ctx context.Context, req *agent.StartInterviewReq) (resp *agent.StartInterviewResp, err error) {
	// TODO: Your code here...
	return
}

// SubmitAnswer implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) SubmitAnswer(ctx context.Context, req *agent.SubmitAnswerReq) (resp *agent.SubmitAnswerResp, err error) {
	// TODO: Your code here...
	return
}

// FinishInterview implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) FinishInterview(ctx context.Context, req *agent.FinishInterviewReq) (resp *agent.FinishInterviewResp, err error) {
	// TODO: Your code here...
	return
}

// GetInterviewDetail implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) GetInterviewDetail(ctx context.Context, req *agent.GetInterviewDetailReq) (resp *agent.GetInterviewDetailResp, err error) {
	// TODO: Your code here...
	return
}

// GetInterviewHistory implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) GetInterviewHistory(ctx context.Context, req *agent.GetInterviewHistoryReq) (resp *agent.GetInterviewHistoryResp, err error) {
	// TODO: Your code here...
	return
}

// AnalyzeResume implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) AnalyzeResume(ctx context.Context, req *agent.AnalyzeResumeReq) (resp *agent.AnalyzeResumeResp, err error) {
	// TODO: Your code here...
	return
}

// GetAbilityAnalysis implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) GetAbilityAnalysis(ctx context.Context, req *agent.GetAbilityAnalysisReq) (resp *agent.GetAbilityAnalysisResp, err error) {
	// TODO: Your code here...
	return
}
