package gochia

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
