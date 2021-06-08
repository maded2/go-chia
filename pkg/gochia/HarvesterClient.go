package gochia

import (
	"fmt"
)

type HarvesterClient struct {
	config *ChiaConfig
}

func NewHarvesterClient(config *ChiaConfig) *HarvesterClient {
	if config.FullNodePort == 0 {
		config.HarvesterPort = 8560
	}
	return &HarvesterClient{
		config: config,
	}
}

func (client *HarvesterClient) GetPlots() (*PlotsResponse, error) {
	var response PlotsResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_plots", client.config.HarvesterPort),
		map[string]interface{}{}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
