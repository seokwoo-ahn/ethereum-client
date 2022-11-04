package rpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
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

func GetTxCountByBlockHash(client rpc.Client, blockHash string) (*big.Int, error) {
	var txCountHex string
	err := client.Call(&txCountHex, "eth_getBlockTransactionCountByHash", blockHash)
	if err != nil {
		return big.NewInt(0), err
	}
	txCount, err := hexutil.DecodeBig(txCountHex)
	if err != nil {
		return big.NewInt(0), err
	}
	return txCount, err
}

func GetTxCountByBlockNum(client rpc.Client, blockNumHex string) (*big.Int, error) {
	var txCountHex string
	err := client.Call(&txCountHex, "eth_getBlockTransactionCountByNumber", blockNumHex)
	if err != nil {
		return big.NewInt(0), err
	}
	txCount, err := hexutil.DecodeBig(txCountHex)
	if err != nil {
		return big.NewInt(0), err
	}
	return txCount, err
}
