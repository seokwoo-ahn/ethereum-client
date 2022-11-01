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

	// address, err := rpc.GetAccounts(*client)
	// latestBlockNum, err := rpc.GetLatestBlockNum(*client)
	chainId, err := rpc.GetChainId(*client)
	if err != nil {
		panic(err)
	}
	// fmt.Println(address)
	// fmt.Println(latestBlockNum)
	fmt.Println(chainId)
}
