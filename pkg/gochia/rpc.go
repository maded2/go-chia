package gochia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func rpc(config *ChiaConfig, url string, parameters map[string]interface{}, result interface{}) error {
	client, err := config.CreateClient()
	if err != nil {
		return err
	}
	postBody, _ := json.Marshal(parameters)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := client.Post(url, "application/json", responseBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if debugResponse {
		fmt.Printf("%s\n", body)
	}
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}
