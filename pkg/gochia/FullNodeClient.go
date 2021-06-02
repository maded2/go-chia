package gochia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	c, err := client.config.CreateClient()
	if err != nil {
		return nil, err
	}
	postBody, _ := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := c.Post(fmt.Sprintf("https://localhost:%d/get_blockchain_state", client.config.FullNodePort), "application/json", responseBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var state FullNodeResponse
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, err
	}
	return &state.State, nil
}
