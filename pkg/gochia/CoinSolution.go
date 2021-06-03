package gochia

type CoinSolution struct {
	Coin     Coin        `json:"coin"`
	Solution interface{} `json:"solution"`
}
