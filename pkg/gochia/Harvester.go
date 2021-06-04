package gochia

type Plot struct {
	FileSize               uint64 `json:"file_size"`
	Filename               string `json:"filename"`
	PlotSeed               string `json:"plot-seed"`
	PlotPublicKey          string `json:"plot_public_key"`
	PoolContractPuzzleHash string `json:"pool_contract_puzzle_hash"`
	PoolPublicKey          string `json:"pool_public_key"`
	Size                   uint64 `json:"size"`
	TimeModified           uint64 `json:"time_modified"`
}

type PlotsResponse struct {
	RpcResponse
	FailedToOpenFilenames []string `json:"failed_to_open_filenames"`
	NotFoundFilenames     []string `json:"not_found_filenames"`
	Plots                 []Plot   `json:"plots"`
}

type PlotDirectoriesResponse struct {
	RpcResponse
	Directories []string `json:"directories"`
}
