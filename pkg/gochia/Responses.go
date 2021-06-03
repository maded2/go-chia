package gochia

type RpcResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

type LoginResponse struct {
	RpcResponse
	BackupInfo BackupInfo `json:"backup_info"`
	BackupPath string     `json:"backup_path"`
}

type PublicKeysResponse struct {
	PublicKeyFingerprints []string `json:"public_key_fingerprints"`
}

type PrivateKeyResponse struct {
	RpcResponse
	PrivateKey []string `json:"private_key"`
}

type GenerateMnemonicResponse struct {
	RpcResponse
	Mnemonic []string `json:"mnemonic"`
}

type AddKeyResponse struct {
	RpcResponse
	Word string `json:"word"`
}

type SyncStatusResponse struct {
	RpcResponse
	Syncing bool `json:"syncing"`
}

type HeightResponse struct {
	RpcResponse
	Height int `json:"height"`
}

type WalletsResponse struct {
	RpcResponse
	Wallets []WalletInfo `json:"wallets"`
}

type WalletBalanceResponse struct {
	RpcResponse
	WalletBalance WalletBalance `json:"wallet_balance"`
}

type TransactionResponse struct {
	RpcResponse
	Transaction   Transaction `json:"transaction"`
	TransactionId string      `json:"transaction_id"`
}

type TransactionsResponse struct {
	RpcResponse
	Transactions []Transaction `json:"transactions"`
	WalletId     int           `json:"wallet_id"`
}

type NextAddressResponse struct {
	RpcResponse
	WalletId int    `json:"wallet_id"`
	Address  string `json:"address"`
}

type PlotsResponse struct {
	RpcResponse
	FailedToOpenFilenames []string `json:"failed_to_open_filenames"`
	NotFoundFilenames     []string `json:"not_found_filenames"`
	Plots                 []Plot   `json:"plots"`
}

type PlotDirectoriesResponse struct {
	RpcResponse
	Directories []string `json:"directories"`
}
