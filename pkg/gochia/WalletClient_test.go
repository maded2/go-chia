package gochia

import (
	"testing"
)

func TestWalletClient_GetWalletBalance(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.crt",
		ChiaKeyFile:  "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.key",
	}
	fullNodeClient := NewWalletClient(config)
	balance, err := fullNodeClient.GetWalletBalance(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", balance)
}
