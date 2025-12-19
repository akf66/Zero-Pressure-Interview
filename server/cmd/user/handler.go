package main

import (
	"context"
	user "zpi/server/shared/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	// TODO: Your code here...
	return
}

// ChangePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (resp *user.ChangePasswordResp, err error) {
	// TODO: Your code here...
	return
}

// UploadResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadResume(ctx context.Context, req *user.UploadResumeReq) (resp *user.UploadResumeResp, err error) {
	// TODO: Your code here...
	return
}

// GetResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetResume(ctx context.Context, req *user.GetResumeReq) (resp *user.GetResumeResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteResume implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteResume(ctx context.Context, req *user.DeleteResumeReq) (resp *user.DeleteResumeResp, err error) {
	// TODO: Your code here...
	return
}
