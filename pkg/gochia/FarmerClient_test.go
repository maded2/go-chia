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
	signage, err := farmerClient.GetSignagePoint("")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", signage)
}
