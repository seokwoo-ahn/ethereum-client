package main

import (
	"flag"
	"fmt"
	"log"

	"ethereum-client/config"

	"github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
	Number string
}

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	config := config.NewConfig(*configFlag)
	apiKey := config.ApiKey
	client, err := rpc.Dial("https://Goerli.infura.io/v3/" + apiKey)
	if err != nil {
		log.Fatalf("Could not connect to Infura: %v", err)
	}

	var lastBlock Block
	err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		fmt.Println("Cannot get the latest block:", err)
		return
	}

	fmt.Printf("Latest block: %v\n", lastBlock.Number)
}
