package gochia

type WalletBalance struct {
	WalletId                 int     `json:"wallet_id"`
	ConfirmedWalletBalance   float64 `json:"confirmed_wallet_balance"`
	MaxSendAmount            float64 `json:"max_send_amount"`
	PendingChange            float64 `json:"pending_change"`
	PendingCoinRemovalCount  int     `json:"pending_coin_removal_count"`
	SpendableBalance         float64 `json:"spendable_balance"`
	UnconfirmedWalletBalance float64 `json:"unconfirmed_wallet_balance"`
	UnspentCoinCount         int     `json:"unspent_coin_count"`
}

type ChiaWalletResponse struct {
	Balance WalletBalance `json:"wallet_balance"`
	Success bool          `json:"success"`
}
