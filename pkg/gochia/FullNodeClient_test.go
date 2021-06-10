package gochia

import (
	"testing"
)

func TestFullNodeClient_GetBlockchainState(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	fullNodeClient := NewFullNodeClient(config)
	state, err := fullNodeClient.GetBlockchainState()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", state)
}

func TestFullNodeClient_GetUnfinishedBlockHeaders(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	fullNodeClient := NewFullNodeClient(config)
	response, err := fullNodeClient.GetUnfinishedBlockHeaders(409311)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", response)
}
