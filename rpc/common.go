package rpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetAccounts(client rpc.Client) ([]common.Address, error) {
	var addresses []common.Address
	err := client.Call(&addresses, "eth_accounts")
	if err != nil {
		return []common.Address{}, err
	}
	return addresses, nil
}

func GetChainId(client rpc.Client) (*big.Int, error) {
	var chainIdHex string
	err := client.Call(&chainIdHex, "eth_chainId")
	if err != nil {
		return big.NewInt(0), err
	}
	chainId, err := hexutil.DecodeBig(chainIdHex)
	if err != nil {
		return big.NewInt(0), err
	}
	return chainId, err
}

func GetGasPrice(client rpc.Client) (*big.Int, error) {
	var gasPriceHex string
	err := client.Call(&gasPriceHex, "eth_chainId")
	if err != nil {
		return big.NewInt(0), err
	}
	gasPrice, err := hexutil.DecodeBig(gasPriceHex)
	if err != nil {
		return big.NewInt(0), err
	}
	return gasPrice, err
}

func GetClientVersion(client rpc.Client) (string, error) {
	var clientVersion string
	err := client.Call(&clientVersion, "web3_clientVersion")
	if err != nil {
		return "", err
	}
	return clientVersion, err
}
