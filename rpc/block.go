package rpc

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetLatestBlock(client rpc.Client) (types.Block, error) {
	var lastBlock types.Block
	err := client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		return lastBlock, err
	}
	return lastBlock, nil
}
