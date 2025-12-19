package main

import (
	"context"
	question "zpi/server/shared/kitex_gen/question"
)

// QuestionServiceImpl implements the last service interface defined in the IDL.
type QuestionServiceImpl struct{}

// CreateQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) CreateQuestion(ctx context.Context, req *question.CreateQuestionReq) (resp *question.CreateQuestionResp, err error) {
	// TODO: Your code here...
	return
}

// GetQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetQuestion(ctx context.Context, req *question.GetQuestionReq) (resp *question.GetQuestionResp, err error) {
	// TODO: Your code here...
	return
}

// GetQuestionList implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetQuestionList(ctx context.Context, req *question.GetQuestionListReq) (resp *question.GetQuestionListResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) UpdateQuestion(ctx context.Context, req *question.UpdateQuestionReq) (resp *question.UpdateQuestionResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) DeleteQuestion(ctx context.Context, req *question.DeleteQuestionReq) (resp *question.DeleteQuestionResp, err error) {
	// TODO: Your code here...
	return
}

// GetCategories implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetCategories(ctx context.Context, req *question.GetCategoriesReq) (resp *question.GetCategoriesResp, err error) {
	// TODO: Your code here...
	return
}

// GetRandomQuestions implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetRandomQuestions(ctx context.Context, req *question.GetRandomQuestionsReq) (resp *question.GetRandomQuestionsResp, err error) {
	// TODO: Your code here...
	return
}

// FavoriteQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) FavoriteQuestion(ctx context.Context, req *question.FavoriteQuestionReq) (resp *question.FavoriteQuestionResp, err error) {
	// TODO: Your code here...
	return
}

// UnfavoriteQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) UnfavoriteQuestion(ctx context.Context, req *question.UnfavoriteQuestionReq) (resp *question.UnfavoriteQuestionResp, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteQuestions implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetFavoriteQuestions(ctx context.Context, req *question.GetFavoriteQuestionsReq) (resp *question.GetFavoriteQuestionsResp, err error) {
	// TODO: Your code here...
	return
}

// AddQuestionNote implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) AddQuestionNote(ctx context.Context, req *question.AddQuestionNoteReq) (resp *question.AddQuestionNoteResp, err error) {
	// TODO: Your code here...
	return
}

// GetQuestionNote implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetQuestionNote(ctx context.Context, req *question.GetQuestionNoteReq) (resp *question.GetQuestionNoteResp, err error) {
	// TODO: Your code here...
	return
}
