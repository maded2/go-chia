package gochia

type RpcResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}
