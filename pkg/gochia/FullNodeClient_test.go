package gochia

import (
	"testing"
)

func TestFullNodeClient_GetBlockchainState(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.crt",
		ChiaKeyFile:  "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.key",
	}
	fullNodeClient := NewFullNodeClient(config)
	state, err := fullNodeClient.GetBlockchainState()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", state)
}
