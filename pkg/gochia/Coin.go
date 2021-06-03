package gochia

type Coin struct {
	ConfirmedBlockIndex uint       `json:"confirmed_block_index"`
	SpentBlockIndex     uint       `json:"spent_block_index"`
	Spent               bool       `json:"spent"`
	Coinbase            bool       `json:"coinbase"`
	WalletType          WalletType `json:"wallet_type"`
	WalletId            uint       `json:"wallet_id"`
}
