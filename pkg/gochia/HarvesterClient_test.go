package gochia

import (
	"testing"
)

func TestHarvesterClient_GetPlots(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: testCertFile,
		ChiaKeyFile:  testKeyFile,
	}
	harvesterClient := NewHarvesterClient(config)
	plots, err := harvesterClient.GetPlots()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", plots)
}
