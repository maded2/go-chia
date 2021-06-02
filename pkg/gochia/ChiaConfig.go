package gochia

import (
	"crypto/tls"
	"net/http"
)

type ChiaConfig struct {
	ChiaCertFile  string // ~/.chia/mainnet/config/ssl/full_node/private_full_node.crt
	ChiaKeyFile   string // ~/.chia/mainnet/config/ssl/full_node/private_full_node.key
	FullNodePort  int
	FarmerPort    int
	HarvesterPort int
	WalletPort    int
}

func (cc *ChiaConfig) CreateClient() (client *http.Client, err error) {
	cert, err := tls.LoadX509KeyPair(cc.ChiaCertFile, cc.ChiaKeyFile)
	if err != nil {
		return
	}
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			},
		},
	}
	return
}
