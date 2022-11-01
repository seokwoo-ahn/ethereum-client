package rpc

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func ExecuteTx(client rpc.Client, tx types.TxData) (CallResp, error) {
	var response CallResp
	err := client.Call(&response, "eth_call", tx, "latest")
	if err != nil {
		return CallResp{}, err
	}
	return response, nil
}
