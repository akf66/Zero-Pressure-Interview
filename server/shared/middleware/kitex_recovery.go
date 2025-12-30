package middleware

import (
	"context"
	"fmt"
	"runtime"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
)

// KitexRecovery Kitex 服务的 panic 恢复中间件
func KitexRecovery() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			defer func() {
				if r := recover(); r != nil {
					// 获取堆栈信息
					buf := make([]byte, 4096)
					n := runtime.Stack(buf, false)
					stack := string(buf[:n])

					klog.CtxErrorf(ctx, "[KitexRecovery] panic recovered: %v\nstack: %s", r, stack)

					// 将 panic 转换为错误
					err = fmt.Errorf("internal server error: %v", r)
				}
			}()
			return next(ctx, req, resp)
		}
	}
}
