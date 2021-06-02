package gochia

import (
	"fmt"
)

type FullNodeClient struct {
	config *ChiaConfig
}

func NewFullNodeClient(config *ChiaConfig) *FullNodeClient {
	if config.FullNodePort == 0 {
		config.FullNodePort = 8555
	}
	return &FullNodeClient{
		config: config,
	}
}

func (client *FullNodeClient) GetBlockchainState() (*ChiaBlockchainState, error) {
	var response FullNodeResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_blockchain_state", client.config.FullNodePort), map[string]interface{}{}, &response)
	if err != nil {
		return nil, err
	}
	return &response.State, nil
}
