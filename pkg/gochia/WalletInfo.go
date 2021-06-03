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
