package main

import (
	"context"
	"zpi/server/shared/dal/sqlfunc"
	base "zpi/server/shared/kitex_gen/base"
	user "zpi/server/shared/kitex_gen/user"

	"github.com/hertz-contrib/paseto"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	EncryptManager
	TokenGenerator
	*UserManager
}

type UserManager struct {
	Query *sqlfunc.Query
}

type EncryptManager interface {
	EncryptPassword(code string) string
}

type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangePassword(ctx context.Context, req *user.ChangePasswordRequest) (resp *user.ChangePasswordResponse, err error) {
	// TODO: Your code here...
	return
}

// UploadResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadResume(ctx context.Context, req *user.UploadResumeRequest) (resp *user.UploadResumeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetResume(ctx context.Context, req *user.GetResumeRequest) (resp *user.GetResumeResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteResume(ctx context.Context, req *user.DeleteResumeRequest) (resp *user.DeleteResumeResponse, err error) {
	// TODO: Your code here...
	return
}

// HealthCheck implements the UserServiceImpl interface.
func (s *UserServiceImpl) HealthCheck(ctx context.Context) (resp *base.HealthCheckResponse, err error) {
	// TODO: Your code here...
	return
}
