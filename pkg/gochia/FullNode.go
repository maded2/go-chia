package gochia

type Block struct {
	ChallengeChainIpProof        Proof            `json:"challenge_chain_ip_proof"`
	ChallengeChainSpProof        Proof            `json:"challenge_chain_sp_proof"`
	FinishedSubSlots             []SubSlot        `json:"finished_sub_slots"`
	Foliage                      FoliageBlock     `json:"foliage"`
	RewardChainBlock             RewardChainBlock `json:"reward_chain_block"`
	InfusedChallengeChainIpProof Proof            `json:"infused_challenge_chain_ip_proof"`
	RewardChainIpProof           Proof            `json:"reward_chain_ip_proof"`
	RewardChainSpProof           Proof            `json:"reward_chain_sp_proof"`
	//TransactionsGenerator        null             `json:"transactions_generator"`

	// Only present on transaction block
	FoliageTransactionBlock FoliageTransactionBlock `json:"foliage_transaction_block"`
	TransactionsInfo        TransactionsInfo        `json:"transactions_info"`

	// Only present if requested by client
	HeaderHash string `json:"header_hash"`
}

type BlockHeader struct {
	data struct {
		AdditionsRoot           string `json:"additions_root"`
		AggregatedSignature     string `json:"aggregated_signature"`
		Cost                    string `json:"cost"`
		ExtensionData           string `json:"extension_data"`
		FarmerRewardsPuzzleHash string `json:"farmer_rewards_puzzle_hash"`
		FilterHash              string `json:"filter_hash"`
		GeneratorHash           string `json:"generator_hash"`
		Height                  uint64 `json:"height"`
		PoolTarget              struct {
			MaxHeight  uint64 `json:"max_height"`
			PuzzleHash string `json:"puzzle_hash"`
		} `json:"pool_target"`
		PrevHeaderHash       string  `json:"prev_header_hash"`
		ProofOfSpaceHash     string  `json:"proof_of_space_hash"`
		RemovalsRoot         string  `json:"removals_root"`
		Timestamp            uint64  `json:"timestamp"`
		TotalIters           uint64  `json:"total_iters"`
		TotalTransactionFees float64 `json:"total_transaction_fees"`
		Weight               string  `json:"weight"`
	}
	PlotSignature string
}

type BlockRecord struct {
	ChallengeBlockInfoHash             string    `json:"challenge_block_info_hash"`
	ChallengeVdfOutput                 VdfOutput `json:"challenge_vdf_output"`
	Deficit                            uint64    `json:"deficit"`
	FarmerPuzzleHash                   string    `json:"farmer_puzzle_hash"`
	Fees                               float64   `json:"fees"`
	FinishedChallengeSlotHashes        []string  `json:"finished_challenge_slot_hashes"`
	FinishedInfusedChallengeSlotHashes []string  `json:"finished_infused_challenge_slot_hashes"`
	FinishedRewardSlotHashes           []string  `json:"finished_reward_slot_hashes"`
	HeaderHash                         string    `json:"header_hash"`
	Height                             uint64    `json:"height"`
	InfusedChallengeVdfOutput          VdfOutput `json:"infused_challenge_vdf_output"`
	Overflow                           bool      `json:"overflow"`
	PoolPuzzleHash                     string    `json:"pool_puzzle_hash"`
	PrevTransactionBlockHash           string    `json:"prev_transaction_block_hash"`
	PrevTransactionBlockHeight         uint64    `json:"prev_transaction_block_height"`
	PrevHash                           string    `json:"prev_hash"`
	RequiredIters                      uint64    `json:"required_iters"`
	// Only transaction blocks have reward claim
	RewardClaimsIncorporated   []RewardClaim `json:"reward_claims_incorporated"`
	RewardInfusionNewChallenge string        `json:"reward_infusion_new_challenge"`
	SignagePointIndex          uint64        `json:"signage_point_index"`
	SubEpochSummaryIncluded    uint64        `json:"sub_epoch_summary_included"`
	SubSlotIters               uint64        `json:"sub_slot_iters"`
	// Only transaction blocks have timestamp
	Timestamp  uint64 `json:"timestamp"`
	TotalIters uint64 `json:"total_iters"`
	Weight     uint64 `json:"weight"`
}

type BlockchainState struct {
	Difficulty   uint64      `json:"difficulty"`
	Peak         BlockRecord `json:"peak"`
	Space        float64     `json:"space"` // TODO need to use uint128
	MempoolSize  uint64      `json:"mempool_size"`
	SubSlotIters uint64      `json:"sub_slot_iters"`
	sync         struct {
		SyncMode           bool   `json:"sync_mode"`
		SyncProgressHeight uint64 `json:"sync_progress_height"`
		SyncTipHeight      uint64 `json:"sync_tip_height"`
		Synced             bool   `json:"synced"`
	}
}

type CoinRecord struct {
	coin struct {
		Amount         float64 `json:"amount"`
		ParentCoinInfo string  `json:"parent_coin_info"`
		PuzzleHash     string  `json:"puzzle_hash"`
	}
	Coinbase            bool   `json:"coinbase"`
	ConfirmedBlockIndex uint64 `json:"confirmed_block_index"`
	Spent               bool   `json:"spent"`
	SpentBlockIndex     uint64 `json:"spent_block_index"`
	Timestamp           uint64 `json:"timestamp"`
}

type FoliageBlock struct {
	FoliageBlockData struct {
		ExtensionData          string `json:"extension_data"`
		FarmerRewardPuzzleHash string `json:"farmer_reward_puzzle_hash"`
		PoolSignature          string `json:"pool_signature"`
		PoolTarget             struct {
			MaxHeight  uint64 `json:"max_height"`
			PuzzleHash string `json:"puzzle_hash"`
		} `json:"pool_target"`
		UnfinishedRewardBlockHash string
	}
	FoliageBlockDataSignature        string `json:"foliage_block_data_signature"`
	FoliageTransactionBlockHash      string `json:"foliage_transaction_block_hash"`
	FoliageTransactionBlockSignature string `json:"foliage_transaction_block_signature"`
	PrevBlockHash                    string `json:"prev_block_hash"`
	RewardBlockHash                  string `json:"reward_block_hash"`
}

type FoliageTransactionBlock struct {
	AdditionsRoot            string `json:"additions_root"`
	FilterHash               string `json:"filter_hash"`
	PrevTransactionBlockHash string `json:"prev_transaction_block_hash"`
	RemovalsRoot             string `json:"removals_root"`
	Timestamp                uint64 `json:"timestamp"`
	TransactionsInfoHash     string `json:"transactions_info_hash"`
}

type Proof struct {
	Witness     string `json:"witness"`
	WitnessType uint   `json:"witness_type"`
}

type ProofOfSpace struct {
	Challenge              string `json:"challenge"`
	PlotPublicKey          string `json:"plot_public_key"`
	PoolContractPuzzleHash string `json:"pool_contract_puzzle_hash"`
	PoolPublicKey          string `json:"pool_public_key"`
	Proof                  string `json:"proof"`
	Size                   uint64 `json:"size"`
}

type RewardChainBlock struct {
	ChallengeChainIpVdf        Vdf          `json:"challenge_chain_ip_vdf"`
	ChallengeChainSpSignature  string       `json:"challenge_chain_sp_signature"`
	ChallengeChainSpVdf        Vdf          `json:"challenge_chain_sp_vdf"`
	Height                     uint64       `json:"height"`
	InfusedChallengeChainIpVdf Vdf          `json:"infused_challenge_chain_ip_vdf"`
	IsTransactionBlock         bool         `json:"is_transaction_block"`
	PosSsCcChallengeHash       string       `json:"pos_ss_cc_challenge_hash"`
	ProofOfSpace               ProofOfSpace `json:"proof_of_space"`
	RewardChainIpVdf           Vdf          `json:"reward_chain_ip_vdf"`
	RewardChainSpSignature     string       `json:"reward_chain_sp_signature"`
	RewardChainSpVdf           Vdf          `json:"reward_chain_sp_vdf"`
	SignagePointIndex          uint64       `json:"signage_point_index"`
	TotalIters                 uint64       `json:"total_iters"`
	Weight                     string       `json:"weight"`
}

type RewardClaim struct {
	Amount         float64 `json:"amount"`
	ParentCoinInfo string  `json:"parent_coin_info"`
	PuzzleHash     string  `json:"puzzle_hash"`
}

type SubSlot struct {
	ChallengeChain struct {
		ChallengeChainEndOfSlotVdf       Vdf    `json:"challenge_chain_end_of_slot_vdf"`
		InfusedChallengeChainSubSlotHash string `json:"infused_challenge_chain_sub_slot_hash"`
		NewDifficulty                    string `json:"new_difficulty"`
		NewSubSlotIters                  string `json:"new_sub_slot_iters"`
		SubepochSummaryHash              string `json:"subepoch_summary_hash"`
	} `json:"challenge_chain"`
	InfusedChallengeChain struct {
		InfusedChallengeChainEndOfSlotVdf Vdf `json:"infused_challenge_chain_end_of_slot_vdf"`
	} `json:"infused_challenge_chain"`
	Proofs struct {
		ChallengeChainSlotProof        Proof `json:"challenge_chain_slot_proof"`
		InfusedChallengeChainSlotProof Proof `json:"infused_challenge_chain_slot_proof"`
		RewardChainSlotProof           Proof `json:"reward_chain_slot_proof"`
	} `json:"proofs"`
	RewardChain struct {
		ChallengeChainSubSlotHash        string `json:"challenge_chain_sub_slot_hash"`
		Deficit                          uint64 `json:"deficit"`
		EndOfSlotVdf                     Vdf    `json:"end_of_slot_vdf"`
		InfusedChallengeChainSubSlotHash string `json:"infused_challenge_chain_sub_slot_hash"`
	} `json:"reward_chain"`
}

type TransactionsInfo struct {
	AggregatedSignature      string        `json:"aggregated_signature"`
	Cost                     string        `json:"cost"`
	Fees                     float64       `json:"fees"`
	GeneratorRoot            string        `json:"generator_root"`
	PreviousGeneratorsRoot   string        `json:"previous_generators_root"`
	RewardClaimsIncorporated []RewardClaim `json:"reward_claims_incorporated"`
}

type Vdf struct {
	Challenge          string    `json:"challenge"`
	NumberOfIterations string    `json:"number_of_iterations"`
	Output             VdfOutput `json:"output"`
}

type VdfOutput struct {
	Data string `json:"data"`
}

type BlockchainStateResponse struct {
	RpcResponse
	BlockchainState BlockchainState `json:"blockchain_state"`
}

type UnfinishedBlockHeadersResponse struct {
	RpcResponse
	Headers []BlockHeader `json:"headers"`
}

type HeaderResponse struct {
	RpcResponse
	Header BlockHeader `json:"header"`
}

type BlockResponse struct {
	RpcResponse
	Block Block `json:"block"`
}

type BlockRecordResponse struct {
	RpcResponse
	BlockRecord BlockRecord `json:"block_record"`
}

type CoinResponse struct {
	RpcResponse
	CoinRecords []CoinRecord `json:"coin_records"`
}

type CoinRecordResponse struct {
	RpcResponse
	CoinRecord CoinRecord `json:"coin_record"`
}

type AdditionsAndRemovalsResponse struct {
	RpcResponse
	Additions []CoinRecord `json:"additions"`
	Removals  []CoinRecord `json:"removals"`
}

type NetspaceResponse struct {
	RpcResponse
	Space uint64 `json:"space"`
}

type BlocksResponse struct {
	RpcResponse
	Blocks []Block `json:"blocks"`
}
