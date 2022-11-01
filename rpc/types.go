package rpc

type CallResp struct {
	Result string
	Error
}

type Error struct {
	Code    int
	Message string
	Data    string
}
