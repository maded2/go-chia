package gochia

import (
	"testing"
)

func TestWalletClient_GetWalletBalance(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	fullNodeClient := NewWalletClient(config)
	balance, err := fullNodeClient.GetWalletBalance(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", balance)
}

func TestWalletClient_GetFarmedAmount(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	fullNodeClient := NewWalletClient(config)
	balance, err := fullNodeClient.GetFarmedAmount()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", balance)
}

func TestWalletClient_LogIn(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	fullNodeClient := NewWalletClient(config)
	success, err := fullNodeClient.LogIn(testFingerprint) // replace with real fingerprint to run this test
	if err != nil {
		t.Fatal(err)
	}
	t.Log(success)
}
