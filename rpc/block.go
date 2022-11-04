package rpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetLatestBlockNum(client rpc.Client) (*big.Int, error) {
	var latestBlockNumHex string
	err := client.Call(&latestBlockNumHex, "eth_blockNumber")
	if err != nil {
		return big.NewInt(0), err
	}
	latestBlockNum, err := hexutil.DecodeBig(latestBlockNumHex)
	if err != nil {
		return big.NewInt(0), err
	}
	return latestBlockNum, err
}

func GetLatestBlock(client rpc.Client) (Block, error) {
	var lastBlock Block
	err := client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		return Block{}, err
	}
	return lastBlock, nil
}

func GetBlockByHash(client rpc.Client, blockHash string, detail bool) (Block, error) {
	var block Block
	err := client.Call(&block, "eth_getBlockByHash", blockHash, detail)
	if err != nil {
		return Block{}, err
	}
	return block, nil
}

func GetBlockByNumber(client rpc.Client, blockNumHex string, detail bool) (Block, error) {
	var block Block
	err := client.Call(&block, "eth_getBlockByNumber", blockNumHex, detail)
	if err != nil {
		return Block{}, err
	}
	return block, nil
}
