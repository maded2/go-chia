package gochia

import (
	"fmt"
)

type FarmerClient struct {
	config *ChiaConfig
}

func NewFarmerClient(config *ChiaConfig) *FarmerClient {
	if config.FullNodePort == 0 {
		config.FarmerPort = 8559
	}
	return &FarmerClient{
		config: config,
	}
}

func (client *FarmerClient) GetSignagePoints() (*SignagePointsResponse, error) {
	var response SignagePointsResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_signage_points", client.config.FarmerPort),
		map[string]interface{}{}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FarmerClient) GetSignagePoint(sp_hash bytes32) (*SignagePointResponse, error) {
	var response SignagePointResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_signage_point", client.config.FarmerPort),
		map[string]interface{}{
			"sp_hash": sp_hash,
		}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
