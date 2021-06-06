package gochia

type WalletType uint

const (
	STANDARD_WALLET  = 0
	RATE_LIMITED     = 1
	ATOMIC_SWAP      = 2
	AUTHORIZED_PAYEE = 3
	MULTI_SIG        = 4
	CUSTODY          = 5
	COLOURED_COIN    = 6
	RECOVERABLE      = 7
)

type WalletInfo struct {
	Id                    uint          `json:"id"`
	Name                  string        `json:"name"`
	Type                  WalletType    `json:"Type"`
	Data                  interface{}   `json:"data"`
	BalanceTotal          float64       `json:"balance_total"`
	BalancePending        float64       `json:"balance_pending"`
	BalanceSpendable      float64       `json:"balance_spendable"`
	BalanceFrozen         float64       `json:"balance_frozen"`
	BalanceChange         float64       `json:"balance_change"`
	Transactions          []Transaction `json:"transactions"`
	Address               string        `json:"address"`
	Colour                string        `json:"colour"`
	SendingTransaction    bool          `json:"sending_transaction"`
	SendTransactionResult string        `json:"send_transaction_result"`
}

type BackupInfo struct {
	BackupHost string `json:"backup_host"`
	Downloaded bool   `json:"downloaded"`
}

type Coin struct {
	ConfirmedBlockIndex uint       `json:"confirmed_block_index"`
	SpentBlockIndex     uint       `json:"spent_block_index"`
	Spent               bool       `json:"spent"`
	Coinbase            bool       `json:"coinbase"`
	WalletType          WalletType `json:"wallet_type"`
	WalletId            uint       `json:"wallet_id"`
}

type CoinSolution struct {
	Coin     Coin        `json:"coin"`
	Solution interface{} `json:"solution"`
}

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

type SpendBundle struct {
	CoinSolutions       []CoinSolution `json:"coin_solutions"`
	AggregatedSignature string         `json:"aggregated_signature"`
}

type Transaction struct {
	ConfirmedAtIndex uint        `json:"confirmed_at_index"`
	CreatedAtTime    uint64      `json:"created_at_time"`
	ToAddress        string      `json:"to_address"`
	Amount           float64     `json:"amount"`
	FeeAmount        float64     `json:"fee_amount"`
	Incoming         bool        `json:"incoming"`
	Confirmed        bool        `json:"confirmed"`
	Sent             float64     `json:"sent"`
	SpendBundle      SpendBundle `json:"spend_bundle"`
	Additions        []Coin      `json:"additions"`
	Removals         []Coin      `json:"removals"`
	WalletId         uint        `json:"wallet_id"`
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
