package gochia

import (
	"testing"
)

func TestFarmerClient_GetSignagePoints(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	farmerClient := NewFarmerClient(config)
	signage, err := farmerClient.GetSignagePoints()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", signage)
}

func TestFarmerClient_GetSignagePoint(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	farmerClient := NewFarmerClient(config)
	signage, err := farmerClient.GetSignagePoint("0x48028f3473c5b6f37d859101549e89d3dd93222b5b4bfad6660b192b3a67a8d4")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", signage)
}
