package gochia

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

/*
Test Certification loading and connecting to local Chia Node
*/
func TestCreateClient(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.crt",
		ChiaKeyFile:  "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.key",
	}
	client, err := config.CreateClient()
	if err != nil {
		t.Fatal(err)
	}
	postBody, _ := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := client.Post("https://localhost:8555/get_blockchain_state", "application/json", responseBody)
	if err != nil {
		t.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	sb := string(body)
	t.Logf(sb)
}

/*
login to wallet
*/
func TestCreateClient2(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.crt",
		ChiaKeyFile:  "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.key",
	}
	client, err := config.CreateClient()
	if err != nil {
		t.Fatal(err)
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"fingerprint": 123,
		"type":        "start",
		"host":        "https://localhost:9256",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := client.Post("https://localhost:9256/log_in", "application/json", responseBody)
	if err != nil {
		t.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	sb := string(body)
	t.Logf(sb)
}

/*
get wallet balance
*/
func TestCreateClient3(t *testing.T) {
	config := &ChiaConfig{
		ChiaCertFile: "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.crt",
		ChiaKeyFile:  "/home/eddie/.chia/mainnet/config/ssl/full_node/private_full_node.key",
	}
	client, err := config.CreateClient()
	if err != nil {
		t.Fatal(err)
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"wallet_id": 1,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := client.Post("https://localhost:9256/get_wallet_balance", "application/json", responseBody)
	if err != nil {
		t.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	sb := string(body)
	t.Logf(sb)
}
