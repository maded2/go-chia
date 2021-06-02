package gochia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type WalletClient struct {
	config *ChiaConfig
}

func NewWalletClient(config *ChiaConfig) *WalletClient {
	if config.FullNodePort == 0 {
		config.WalletPort = 9256
	}
	return &WalletClient{
		config: config,
	}
}

func (client *WalletClient) GetWalletBalance(walletId int) (*WalletBalance, error) {
	c, err := client.config.CreateClient()
	if err != nil {
		return nil, err
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"wallet_id": walletId,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := c.Post(fmt.Sprintf("https://localhost:%d/get_wallet_balance", client.config.WalletPort), "application/json", responseBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var state ChiaWalletResponse
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, err
	}
	return &state.Balance, nil
}

func (client *WalletClient) LogIn(fingerprint int) (bool, error) {
	c, err := client.config.CreateClient()
	if err != nil {
		return false, err
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"fingerprint": fingerprint,
		"type":        "start",
		"host":        fmt.Sprintf("https://localhost:%d", client.config.WalletPort),
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := c.Post(fmt.Sprintf("https://localhost:%d/log_in", client.config.WalletPort), "application/json", responseBody)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var state ChiaWalletResponse
	err = json.Unmarshal(body, &state)
	if err != nil {
		return false, err
	}
	return state.Success, err
}
