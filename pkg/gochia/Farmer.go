package gochia

//type ProofOfSpace struct {
//	PlotIdentifier string `json:"plot_identifier"`
//	Proof          string  `json:"proof"`
//}

type SignagePoint struct {
	ChallengeHash     string `json:"challenge_hash"`
	ChallengeChainSp  string `json:"challenge_chain_sp"`
	RewardChainSp     string `json:"reward_chain_sp"`
	Difficulty        uint   `json:"difficulty"`
	SubSlotIters      uint64 `json:"sub_slot_iters"`
	SignagePointIndex uint64 `json:"signage_point_index"`
}

type SignagePointResponse struct {
	RpcResponse
	SignagePoint SignagePoint   `json:"signage_point"`
	Proofs       []ProofOfSpace `json:"proofs"`
}

type SignagePointsResponse struct {
	RpcResponse
	SignagePoints []struct {
		SignagePoint SignagePoint   `json:"signage_point"`
		Proofs       []ProofOfSpace `json:"proofs"`
	} `json:"signage_points"`
}
type RewardTargetResponse struct {
	RpcResponse
	FarmerTarget string `json:"farmer_target"`
	PoolTarget   string `json:"pool_target"`
	HaveFarmerSk bool   `json:"have_farmer_sk"`
	HavePoolSk   bool   `json:"have_pool_sk"`
}
