package gochia

type SpendBundle struct {
	CoinSolutions       []CoinSolution `json:"coin_solutions"`
	AggregatedSignature string         `json:"aggregated_signature"`
}
