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

type Block struct {
	Number           string
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	ReceiptsRoot     string
	Miner            string
	Difficulty       string
	TotalDifficulty  string
	ExtraData        string
	Size             string
	GasLimit         string
	GasUsed          string
	TimeStamp        string
	Transactions     []interface{}
	Uncles           []string
}
