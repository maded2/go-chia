package gochia

import (
	"fmt"
)

type FullNodeClient struct {
	config *ChiaConfig
}

func NewFullNodeClient(config *ChiaConfig) *FullNodeClient {
	if config.FullNodePort == 0 {
		config.FullNodePort = 8555
	}
	return &FullNodeClient{
		config: config,
	}
}

func (client *FullNodeClient) GetBlockchainState() (*BlockchainStateResponse, error) {
	var response BlockchainStateResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_blockchain_state", client.config.FullNodePort), map[string]interface{}{}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetNetworkSpace(newerBlockHeaderHash, olderBlockHeaderHash string) (*NetspaceResponse, error) {
	var response NetspaceResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_network_space", client.config.FullNodePort), map[string]interface{}{
		"newer_block_header_hash": newerBlockHeaderHash,
		"older_block_header_hash": olderBlockHeaderHash,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetBlocks(start, end uint64, excludeHeaderHash bool) (*BlocksResponse, error) {
	var response BlocksResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_blocks", client.config.FullNodePort), map[string]interface{}{
		"start":               start,
		"end":                 end,
		"exclude_header_hash": excludeHeaderHash,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetBlock(headerHash string) (*BlockResponse, error) {
	var response BlockResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_block", client.config.FullNodePort), map[string]interface{}{
		"header_hash": headerHash,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetBlockRecordByHeight(height uint64) (*BlockRecordResponse, error) {
	var response BlockRecordResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_block_record_by_height", client.config.FullNodePort), map[string]interface{}{
		"height": height,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetBlockRecord(hash string) (*BlockRecordResponse, error) {
	var response BlockRecordResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_block_record", client.config.FullNodePort), map[string]interface{}{
		"header_hash": hash,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetUnfinishedBlockHeaders(height uint32) (*UnfinishedBlockHeadersResponse, error) {
	var response UnfinishedBlockHeadersResponse
	params := map[string]interface{}{
		"height": height,
	}
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_unfinished_block_headers", client.config.FullNodePort), params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetUnspentCoins(puzzleHash string, startHeight uint64, endHeight uint64) (*CoinResponse, error) {
	var response CoinResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_coin_records_by_puzzle_hash", client.config.FullNodePort), map[string]interface{}{
		"puzzle_hash":         puzzleHash,
		"start_height":        startHeight,
		"end_height":          endHeight,
		"include_spent_coins": false,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetCoinRecordByName(name string) (*CoinRecordResponse, error) {
	var response CoinRecordResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_coin_record_by_name", client.config.FullNodePort), map[string]interface{}{
		"name": name,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *FullNodeClient) GetAdditionsAndRemovals(hash string) (*AdditionsAndRemovalsResponse, error) {
	var response AdditionsAndRemovalsResponse
	err := rpc(client.config, fmt.Sprintf("https://localhost:%d/get_additions_and_removals", client.config.FullNodePort), map[string]interface{}{
		"hash": hash,
	}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
