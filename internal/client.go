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

	app                   *tview.Application
	nodeTable             *tview.Table
	walletTable           *tview.Table
	farmTable             *tview.Table
	unfinishedBlocksTable *tview.Table
	latestChallengeTable  *tview.Table
}

const gib = 1024 * 1024 * 1024
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
	farmerClient := gochia.NewFarmerClient(config)

	ticker := time.NewTicker(time.Second * 2)
	c.requestAndDraw(nodeClient, walletClient, harvesterClient, farmerClient)
	for range ticker.C {
		c.requestAndDraw(nodeClient, walletClient, harvesterClient, farmerClient)
	}
}

func (c *Client) requestAndDraw(nodeClient *gochia.FullNodeClient, walletClient *gochia.WalletClient, harvesterClient *gochia.HarvesterClient, farmerClient *gochia.FarmerClient) {
	if state, err := nodeClient.GetBlockchainState(); err == nil && state.Success {
		c.nodeTable.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%s", getSyncState(&state.BlockchainState))))
		c.nodeTable.SetCell(1, 1, tview.NewTableCell(fmt.Sprintf("%d", state.BlockchainState.Peak.Height)))

		// TODO need investigation, seem sometimes Timestamp returned from the RPC call is not set.
		if state.BlockchainState.Peak.Timestamp > 0 {
			t := time.Unix(int64(state.BlockchainState.Peak.Timestamp), 0)
			c.nodeTable.SetCell(2, 1, tview.NewTableCell(fmt.Sprintf("%s (%d)", t.Format("2006-01-02 15:04:05"), state.BlockchainState.Peak.Timestamp)))
		}
		c.nodeTable.SetCell(3, 1, tview.NewTableCell(fmt.Sprintf("%d", state.BlockchainState.Difficulty)))
		c.nodeTable.SetCell(4, 1, tview.NewTableCell(fmt.Sprintf("%.3f EiB", state.BlockchainState.Space/eib)))

		if headers, err := nodeClient.GetUnfinishedBlockHeaders(state.BlockchainState.Peak.Height); err == nil && headers.Success {
			c.unfinishedBlocksTable.Clear()
			c.unfinishedBlocksTable.SetCell(0, 0, tview.NewTableCell("Header Hash"))
			c.unfinishedBlocksTable.SetCell(0, 1, tview.NewTableCell("Timestamp"))
			for n, header := range headers.Headers {
				c.unfinishedBlocksTable.SetCell(n+1, 0, tview.NewTableCell(fmt.Sprintf("%s", header.Foliage.FoliageTransactionBlockHash)))
				t := time.Unix(int64(header.FoliageTransactionBlock.Timestamp), 0)
				c.unfinishedBlocksTable.SetCell(n+1, 1, tview.NewTableCell(fmt.Sprintf("%s", t.Format("2006-01-02 15:04:05"))))
			}
		}
	}
	if balance, err := walletClient.GetWalletBalance(1); err == nil && balance.Success {
		c.walletTable.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%f XCH", balance.WalletBalance.ConfirmedWalletBalance)))
		c.walletTable.SetCell(1, 1, tview.NewTableCell(fmt.Sprintf("%f XCH", balance.WalletBalance.SpendableBalance)))
		c.walletTable.SetCell(2, 1, tview.NewTableCell(fmt.Sprintf("%f XCH", balance.WalletBalance.UnconfirmedWalletBalance)))
		c.walletTable.SetCell(3, 1, tview.NewTableCell(fmt.Sprintf("%f XCH", balance.WalletBalance.PendingChange)))
	}
	//if farmed, err := walletClient.GetFarmedAmount(); err == nil && farmed.Success {
	//	c.farmTable.SetCell(0, 1, tview.NewTableCell(fmt.Sprintf("%.10f XCH", farmed.FarmedAmount/10e11)))
	//	c.farmTable.SetCell(1, 1, tview.NewTableCell(fmt.Sprintf("%.10f XCH", farmed.FeeAmount/10e11)))
	//}
	if challenges, err := farmerClient.GetSignagePoints(); err == nil && challenges.Success {
		c.latestChallengeTable.Clear()
		c.latestChallengeTable.SetCell(0, 0, tview.NewTableCell("Challenge Hash"))
		c.latestChallengeTable.SetCell(0, 1, tview.NewTableCell("Index"))

		for n, challenge := range challenges.SignagePoints {
			c.latestChallengeTable.SetCell(n+1, 0, tview.NewTableCell(fmt.Sprintf("%s", challenge.SignagePoint.ChallengeHash)))
			c.latestChallengeTable.SetCell(n+1, 1, tview.NewTableCell(fmt.Sprintf("%d", challenge.SignagePoint.SignagePointIndex)))
		}
		c.latestChallengeTable.ScrollToEnd()
	}

	if plots, err := harvesterClient.GetPlots(); err == nil && plots.Success {
		c.farmTable.Clear()
		c.farmTable.SetTitle(fmt.Sprintf("Farming: %d Plots, %d Error Plots, %d Missing Plots", len(plots.Plots), len(plots.FailedToOpenFilenames), len(plots.FailedToOpenFilenames)))
		c.farmTable.SetCell(0, 0, tview.NewTableCell("Plot Size"))
		c.farmTable.SetCell(0, 1, tview.NewTableCell("Plot Key"))
		c.farmTable.SetCell(0, 2, tview.NewTableCell("Filename"))
		for n, plot := range plots.Plots {
			c.farmTable.SetCell(n+1, 0, tview.NewTableCell(fmt.Sprintf("K-%d, %.1f GiB", plot.Size, float64(plot.FileSize)/gib)))
			c.farmTable.SetCell(n+1, 1, tview.NewTableCell(plot.PlotPublicKey))
			c.farmTable.SetCell(n+1, 2, tview.NewTableCell(plot.Filename))
		}
		c.farmTable.ScrollToBeginning()
	}

	c.app.Draw()
}

func (c *Client) setupUI() {
	c.nodeTable = tview.NewTable()
	c.nodeTable.SetTitle("Node").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	c.nodeTable.SetCell(0, 0, tview.NewTableCell("Node State"))
	c.nodeTable.SetCell(1, 0, tview.NewTableCell("Peak Height"))
	c.nodeTable.SetCell(2, 0, tview.NewTableCell("Peak Height Timestamp"))
	c.nodeTable.SetCell(3, 0, tview.NewTableCell("Difficulty"))
	c.nodeTable.SetCell(4, 0, tview.NewTableCell("Estimated Network Space:"))

	c.unfinishedBlocksTable = tview.NewTable()
	c.unfinishedBlocksTable.SetTitle("Unfinished Block").SetBorder(true).SetTitleAlign(tview.AlignLeft)

	panel1 := tview.NewFlex()
	panel1.SetDirection(tview.FlexColumn)
	panel1.AddItem(c.nodeTable, 0, 1, false)
	panel1.AddItem(c.unfinishedBlocksTable, 0, 1, false)

	c.walletTable = tview.NewTable()
	c.walletTable.SetTitle("Wallet").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	c.walletTable.SetCell(0, 0, tview.NewTableCell("Total Balance"))
	c.walletTable.SetCell(1, 0, tview.NewTableCell("Spendable Balance"))
	c.walletTable.SetCell(2, 0, tview.NewTableCell("Pending Balance"))
	c.walletTable.SetCell(3, 0, tview.NewTableCell("Pending Change"))

	c.latestChallengeTable = tview.NewTable()
	c.latestChallengeTable.SetTitle("Latest Block Challenges").SetBorder(true).SetTitleAlign(tview.AlignLeft)
	panel2 := tview.NewFlex()
	panel2.SetDirection(tview.FlexColumn)
	panel2.AddItem(c.walletTable, 0, 1, false)
	panel2.AddItem(c.latestChallengeTable, 0, 1, false)

	c.farmTable = tview.NewTable()
	c.farmTable.SetTitle("Farm").SetBorder(true).SetTitleAlign(tview.AlignLeft)

	mainPanel := tview.NewFlex()
	mainPanel.SetDirection(tview.FlexRow)
	mainPanel.AddItem(panel1, 0, 1, false)
	mainPanel.AddItem(panel2, 0, 1, false)
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
