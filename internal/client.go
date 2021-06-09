package internal

import (
	"fmt"
	"github.com/rivo/tview"
	"gochia/pkg/gochia"
	"time"
)

type Client struct {
	CertFile    string
	KeyFile     string
	Fingerprint string

	app         *tview.Application
	nodeTable   *tview.Table
	walletTable *tview.Table
	farmTable   *tview.Table
}

const eib = 1024 * 1024 * 1024 * 1024 * 1024 * 1024

func (c *Client) Run() {
	c.setupUI()

	go c.processLoop()
	c.app.Run()
}

func (c *Client) processLoop() {
	config := &gochia.ChiaConfig{
		ChiaCertFile: c.CertFile,
		ChiaKeyFile:  c.KeyFile,
	}
	nodeClient := gochia.NewFullNodeClient(config)
	walletClient := gochia.NewWalletClient(config)
	harvesterClient := gochia.NewHarvesterClient(config)

	ticker := time.NewTicker(time.Second * 10)
	for range ticker.C {
		if state, err := nodeClient.GetBlockchainState(); err == nil && state.Success {
			c.nodeTable.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%s", getSyncState(&state.BlockchainState))))
			c.nodeTable.SetCell(1, 1, tview.NewTableCell(fmt.Sprintf("%d", state.BlockchainState.Peak.Height)))

			// TODO need investigation, seem sometimes Timestamp returned from the RPC call is not set.
			t := time.Unix(int64(state.BlockchainState.Peak.Timestamp), 0)
			c.nodeTable.SetCell(2, 1, tview.NewTableCell(fmt.Sprintf("%s (%d)", t.Format("2006-01-02 15:04:05"), state.BlockchainState.Peak.Timestamp)))

			c.nodeTable.SetCell(3, 1, tview.NewTableCell(fmt.Sprintf("%d", state.BlockchainState.Difficulty)))
			c.nodeTable.SetCell(4, 1, tview.NewTableCell(fmt.Sprintf("%.3f EiB", state.BlockchainState.Space/eib)))
		}
		if balance, err := walletClient.GetWalletBalance(1); err == nil && balance.Success {
			c.walletTable.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%f XCH", balance.WalletBalance.ConfirmedWalletBalance)))
		}
		if farmed, err := walletClient.GetFarmedAmount(); err == nil && farmed.Success {
			c.farmTable.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%.10f XCH", farmed.FarmedAmount/10e11)))
			c.farmTable.SetCell(1, 1, tview.NewTableCell(fmt.Sprintf("%.10f XCH", farmed.FeeAmount/10e11)))
		}
		if plots, err := harvesterClient.GetPlots(); err == nil && plots.Success {
			c.farmTable.SetCell(2, 1, tview.NewTableCell(fmt.Sprintf("%d Plots, Failed Plots: %d, NotFound Plots: %d\n", len(plots.Plots), len(plots.FailedToOpenFilenames), len(plots.NotFoundFilenames))))
		}

		c.app.Draw()
	}
}

func (c *Client) setupUI() {
	c.nodeTable = tview.NewTable()
	c.nodeTable.SetTitle("Node").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	c.nodeTable.SetCell(0, 0, tview.NewTableCell("Node State"))
	c.nodeTable.SetCell(1, 0, tview.NewTableCell("Peak Height"))
	c.nodeTable.SetCell(2, 0, tview.NewTableCell("Peak Height Timestamp"))
	c.nodeTable.SetCell(3, 0, tview.NewTableCell("Difficulty"))
	c.nodeTable.SetCell(4, 0, tview.NewTableCell("Estimated Network Space:"))

	c.walletTable = tview.NewTable()
	c.walletTable.SetTitle("Wallet").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	c.walletTable.SetCell(0, 0, tview.NewTableCell("Wallet Balance"))

	c.farmTable = tview.NewTable()
	c.farmTable.SetTitle("Farm").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	c.farmTable.SetCell(0, 0, tview.NewTableCell("Farmed"))
	c.farmTable.SetCell(1, 0, tview.NewTableCell("Fees"))
	c.farmTable.SetCell(2, 0, tview.NewTableCell("Farming"))

	mainPanel := tview.NewFlex()
	mainPanel.SetDirection(tview.FlexRow)
	mainPanel.AddItem(c.nodeTable, 0, 1, false)
	mainPanel.AddItem(c.walletTable, 0, 1, false)
	mainPanel.AddItem(c.farmTable, 0, 1, false)
	c.app = tview.NewApplication()
	c.app.SetRoot(mainPanel, true)
}

func getSyncState(state *gochia.BlockchainState) string {
	if state.Sync.Synced {
		return "Synced"
	} else {
		return fmt.Sprintf("Syncing %d / %d", state.Sync.SyncProgressHeight, state.Sync.SyncTipHeight)
	}
}
