package tools

import (
	"errors"
	"zpi/server/shared/errno"
	"zpi/server/shared/kitex_gen/base"
)

func BuildBaseResp(err error) *base.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *base.BaseResponse {
	return &base.BaseResponse{
		StatusCode: int64(err.ErrCode),
		StatusMsg:  err.ErrMsg,
	}
}

func ParseBaseResp(resp *base.BaseResponse) error {
	if base.ErrCode(resp.StatusCode) == errno.Success.ErrCode {
		return nil
	}
	return errno.NewErrNo(base.ErrCode(resp.StatusCode), resp.StatusMsg)
}
