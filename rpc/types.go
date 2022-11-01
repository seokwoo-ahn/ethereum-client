package rpc

type CallResp struct {
	Jsonrpc string
	Id      int
	Result  string
	Error
}

type Error struct {
	Code    int
	Message string
	Data    string
}
