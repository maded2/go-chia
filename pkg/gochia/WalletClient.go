package gochia

import (
	"fmt"
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

func (client *WalletClient) GetWalletBalance(walletId int) (*WalletBalanceResponse, error) {
	var response WalletBalanceResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_wallet_balance", client.config.WalletPort),
		map[string]interface{}{
			"wallet_id": walletId,
		}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *WalletClient) LogIn(fingerprint int) (bool, error) {
	var response WalletBalanceResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/log_in", client.config.WalletPort),
		map[string]interface{}{

			"fingerprint": fingerprint,
			"type":        "start",
			"host":        fmt.Sprintf("https://localhost:%d", client.config.WalletPort),
		}, &response)
	if err != nil {
		return false, err
	}
	return response.Success, nil
}
