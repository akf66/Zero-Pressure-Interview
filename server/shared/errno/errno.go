package errno

import (
	"fmt"

	"zpi/server/shared/kitex_gen/base"
)

// ErrNo 错误号结构
type ErrNo struct {
	ErrCode base.ErrCode
	ErrMsg  string
}

// Error 实现 error 接口
func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// NewErrNo 创建新的错误号
func NewErrNo(code base.ErrCode, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

// WithMessage 修改错误消息
func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// Response HTTP 响应结构
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 预定义的错误常量
var (
	// 成功
	Success = NewErrNo(base.ErrCode_Success, "success")

	// 通用错误 1xxxx
	NoRoute            = NewErrNo(base.ErrCode_NoRoute, "no route")
	NoMethod           = NewErrNo(base.ErrCode_NoMethod, "no method")
	BadRequest         = NewErrNo(base.ErrCode_BadRequest, "bad request")
	ParamsErr          = NewErrNo(base.ErrCode_ParamsErr, "params error")
	AuthorizeFail      = NewErrNo(base.ErrCode_AuthorizeFail, "authorize failed")
	TooManyRequest     = NewErrNo(base.ErrCode_TooManyRequest, "too many requests")
	ServiceErr         = NewErrNo(base.ErrCode_ServiceErr, "service error")
	RecordNotFound     = NewErrNo(base.ErrCode_RecordNotFound, "record not found")
	RecordAlreadyExist = NewErrNo(base.ErrCode_RecordAlreadyExist, "record already exist")

	// 用户服务错误 2xxxx
	RPCUserSrvErr     = NewErrNo(base.ErrCode_RPCUserSrvErr, "rpc user service error")
	UserSrvErr        = NewErrNo(base.ErrCode_UserSrvErr, "user service error")
	UserNotFound      = NewErrNo(base.ErrCode_UserNotFound, "user not found")
	PasswordError     = NewErrNo(base.ErrCode_PasswordError, "password error")
	EmailAlreadyExist = NewErrNo(base.ErrCode_EmailAlreadyExist, "email already exist")
	InvalidPassword   = NewErrNo(base.ErrCode_InvalidPassword, "invalid password")

	// 面试服务错误 3xxxx
	RPCInterviewSrvErr       = NewErrNo(base.ErrCode_RPCInterviewSrvErr, "rpc interview service error")
	InterviewSrvErr          = NewErrNo(base.ErrCode_InterviewSrvErr, "interview service error")
	InterviewNotFound        = NewErrNo(base.ErrCode_InterviewNotFound, "interview not found")
	InterviewAlreadyFinished = NewErrNo(base.ErrCode_InterviewAlreadyFinished, "interview already finished")
	InvalidInterviewType     = NewErrNo(base.ErrCode_InvalidInterviewType, "invalid interview type")

	// 题库服务错误 4xxxx
	RPCQuestionSrvErr = NewErrNo(base.ErrCode_RPCQuestionSrvErr, "rpc question service error")
	QuestionSrvErr    = NewErrNo(base.ErrCode_QuestionSrvErr, "question service error")
	QuestionNotFound  = NewErrNo(base.ErrCode_QuestionNotFound, "question not found")
	InvalidDifficulty = NewErrNo(base.ErrCode_InvalidDifficulty, "invalid difficulty")
	CategoryNotFound  = NewErrNo(base.ErrCode_CategoryNotFound, "category not found")

	// 存储服务错误 5xxxx
	RPCStorageSrvErr = NewErrNo(base.ErrCode_RPCStorageSrvErr, "rpc storage service error")
	StorageSrvErr    = NewErrNo(base.ErrCode_StorageSrvErr, "storage service error")
	FileUploadError  = NewErrNo(base.ErrCode_FileUploadError, "file upload error")
	FileNotFound     = NewErrNo(base.ErrCode_FileNotFound, "file not found")
	InvalidFileType  = NewErrNo(base.ErrCode_InvalidFileType, "invalid file type")
	FileSizeExceeded = NewErrNo(base.ErrCode_FileSizeExceeded, "file size exceeded")
)

// ConvertErr 将 error 转换为 ErrNo
func ConvertErr(err error) ErrNo {
	if err == nil {
		return Success
	}

	// 如果已经是 ErrNo 类型，直接返回
	if e, ok := err.(ErrNo); ok {
		return e
	}

	// 否则返回通用服务错误
	return ServiceErr.WithMessage(err.Error())
}

// BuildBaseResp 构建 BaseResponse
func BuildBaseResp(err error) *base.BaseResponse {
	if err == nil {
		return &base.BaseResponse{
			StatusCode: int64(base.ErrCode_Success),
			StatusMsg:  "success",
		}
	}

	e := ConvertErr(err)
	return &base.BaseResponse{
		StatusCode: int64(e.ErrCode),
		StatusMsg:  e.ErrMsg,
	}
}
