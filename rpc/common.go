package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetAccount(client rpc.Client) ([]common.Address, error) {
	var addresses []common.Address
	err := client.Call(&addresses, "eth_accounts")
	if err != nil {
		return addresses, err
	}
	return addresses, nil
}
