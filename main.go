package main

import (
	"flag"
	"fmt"

	"ethereum-client/client"
	"ethereum-client/config"
	"ethereum-client/rpc"
)

type Block struct {
	Number string
}

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	config := config.NewConfig(*configFlag)
	client, err := client.ConnectClient(config)
	if err != nil {
		panic(err)
	}

	// blockHash := "0xffac2cfa6ff05d28c83656e7dd67badd00f2ee015d5934a74a07d18b761ca526"
	// blockNum := "0x7860fe"

	// address, err := rpc.GetAccounts(*client)
	// latestBlockNum, err := rpc.GetLatestBlockNum(*client)
	// chainId, err := rpc.GetChainId(*client)
	// gasPrice, err := rpc.GetGasPrice(*client)
	// block, err := rpc.GetLatestBlock(*client)
	// block, err := rpc.GetBlockByHash(*client, blockHash, false)
	// block, err := rpc.GetBlockByNumber(*client, "0x9", false)
	// txCount, err := rpc.GetTxCountByBlockHash(*client, blockHash)
	// txCount, err := rpc.GetTxCountByBlockNum(*client, blockNum)
	// clientVersion, err := rpc.GetClientVersion(*client)
	// peerCount, err := rpc.GetNetPeerCount(*client)
	isListening, err := rpc.IsListening(*client)
	if err != nil {
		panic(err)
	}
	// fmt.Println(address)
	// fmt.Println(latestBlockNum)
	// fmt.Println(chainId)
	// fmt.Println(gasPrice)
	// fmt.Println(block)
	// fmt.Println(txCount)
	// fmt.Println(clientVersion)
	// fmt.Println(peerCount)
	fmt.Println(isListening)
}
