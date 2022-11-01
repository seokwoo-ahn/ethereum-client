package client

import (
	"ethereum-client/config"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

func ConnectClient(config *config.Config) (*rpc.Client, error) {
	apiKey := config.ApiKey
	client, err := rpc.Dial("https://Goerli.infura.io/v3/" + apiKey)
	if err != nil {
		log.Fatalf("Could not connect to Infura: %v", err)
	}
	return client, err
}
