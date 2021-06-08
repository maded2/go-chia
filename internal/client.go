package internal

import (
	"fmt"
	"gochia/pkg/gochia"
	"time"
)

type Client struct {
	CertFile    string
	KeyFile     string
	Fingerprint string
}

const eib = 1024 * 1024 * 1024 * 1024 * 1024 * 1024

func (c *Client) Run() {
	nodeClient := gochia.NewFullNodeClient(&gochia.ChiaConfig{
		ChiaCertFile: c.CertFile,
		ChiaKeyFile:  c.KeyFile,
	})
	if state, err := nodeClient.GetBlockchainState(); err == nil && state.Success {
		fmt.Printf("Node State: %s\n", getSyncState(&state.BlockchainState))
		fmt.Printf("Peak Height: %d\n", state.BlockchainState.Peak.Height)

		t := time.Unix(int64(state.BlockchainState.Peak.Timestamp), 0)
		fmt.Printf("Peak Height Timestamp: %s\n", t.Format("2006-01-02 15:04:05"))

		fmt.Printf("Diffculty: %d\n", state.BlockchainState.Difficulty)
		fmt.Printf("Estimated Network Space: %.3f EiB\n", state.BlockchainState.Space/eib)
	} else {
		fmt.Printf("Can not get Chia Node state: %s\n", err)
	}
	walletClient := gochia.NewWalletClient(&gochia.ChiaConfig{
		ChiaCertFile: c.CertFile,
		ChiaKeyFile:  c.KeyFile,
	})
	if balance, err := walletClient.GetWalletBalance(1); err == nil && balance.Success {
		fmt.Printf("Wallet Balance: %f XCH\n", balance.WalletBalance.ConfirmedWalletBalance)
	} else {
		fmt.Printf("Can not get Chia Wallet balance: %s\n", err)
	}
	if farmed, err := walletClient.GetFarmedAmount(); err == nil && farmed.Success {
		fmt.Printf("Farmed: %.10f XCH\n", farmed.FarmedAmount/10e11)
		fmt.Printf("Fees: %.10f XCH\n", farmed.FeeAmount/10e11)
	} else {
		fmt.Printf("Can not get Chia Farmed: %s\n", err)
	}
}

func getSyncState(state *gochia.BlockchainState) string {
	if state.Sync.Synced {
		return "Synced"
	} else {
		return fmt.Sprintf("Syncing %d / %d", state.Sync.SyncProgressHeight, state.Sync.SyncTipHeight)
	}
}
