package rpc

// InitRPC 初始化所有RPC客户端
func InitRPC() {
	InitUser()
	InitInterview()
	InitQuestion()
	InitStorage()
}
