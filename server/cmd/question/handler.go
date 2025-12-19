package main

import (
	"context"
	question "zpi/server/shared/kitex_gen/question"
)

// QuestionServiceImpl implements the last service interface defined in the IDL.
type QuestionServiceImpl struct{}

// CreateQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) CreateQuestion(ctx context.Context, req *question.CreateQuestionRequest) (resp *question.CreateQuestionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetQuestion(ctx context.Context, req *question.GetQuestionRequest) (resp *question.GetQuestionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetQuestionList implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetQuestionList(ctx context.Context, req *question.GetQuestionListRequest) (resp *question.GetQuestionListResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) UpdateQuestion(ctx context.Context, req *question.UpdateQuestionRequest) (resp *question.UpdateQuestionResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) DeleteQuestion(ctx context.Context, req *question.DeleteQuestionRequest) (resp *question.DeleteQuestionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCategories implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetCategories(ctx context.Context, req *question.GetCategoriesRequest) (resp *question.GetCategoriesResponse, err error) {
	// TODO: Your code here...
	return
}

// GetRandomQuestions implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetRandomQuestions(ctx context.Context, req *question.GetRandomQuestionsRequest) (resp *question.GetRandomQuestionsResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) FavoriteQuestion(ctx context.Context, req *question.FavoriteQuestionRequest) (resp *question.FavoriteQuestionResponse, err error) {
	// TODO: Your code here...
	return
}

// UnfavoriteQuestion implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) UnfavoriteQuestion(ctx context.Context, req *question.UnfavoriteQuestionRequest) (resp *question.UnfavoriteQuestionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteQuestions implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetFavoriteQuestions(ctx context.Context, req *question.GetFavoriteQuestionsRequest) (resp *question.GetFavoriteQuestionsResponse, err error) {
	// TODO: Your code here...
	return
}

// AddQuestionNote implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) AddQuestionNote(ctx context.Context, req *question.AddQuestionNoteRequest) (resp *question.AddQuestionNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// GetQuestionNote implements the QuestionServiceImpl interface.
func (s *QuestionServiceImpl) GetQuestionNote(ctx context.Context, req *question.GetQuestionNoteRequest) (resp *question.GetQuestionNoteResponse, err error) {
	// TODO: Your code here...
	return
}
