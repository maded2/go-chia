package gochia

type ChiaBlockchainState struct {
	Difficulty                  uint     `json:"difficulty"`
	GenesisChallengeInitialized bool     `json:"genesis_challenge_initialized"`
	MempoolSize                 uint     `json:"mempool_size"`
	Sync                        ChiaSync `json:"sync"`
	Space                       uint64   `json:"space"`
	Peak                        ChiaPeak `json:"peak"`
}

type FullNodeResponse struct {
	State   ChiaBlockchainState `json:"blockchain_state"`
	Success bool                `json:"success"`
}

type ChiaPeak struct {
	Height uint `json:"height"`
}

type ChiaSync struct {
	SyncMode bool `json:"sync_mode"`
	Synced   bool `json:"synced"`
}
