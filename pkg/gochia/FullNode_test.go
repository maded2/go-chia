package gochia

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestFullNodeResponse(t *testing.T) {
	buffer, err := ioutil.ReadFile("full_node_response.json")
	if err != nil {
		t.Fatal(err)
	}
	var resp FullNodeResponse
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)
}
