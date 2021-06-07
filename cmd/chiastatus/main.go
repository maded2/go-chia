package main

import (
	"flag"
	"fmt"
	"gochia/internal"
	"os"
)

func main() {
	certFileStr := flag.String("CertFile", "", "")
	keyFileStr := flag.String("keyfile", "", "")
	fingerprint := flag.String("fingerprint", "", "")
	flag.Parse()
	if !flag.Parsed() {
		flag.Usage()
		return
	}

	certFile := *certFileStr
	keyFile := *keyFileStr
	if len(*certFileStr) == 0 || len(*keyFileStr) == 0 {
		if homeDir, err := os.UserHomeDir(); err == nil {
			certFile = fmt.Sprintf("%s%c%s", homeDir, os.PathSeparator, ".chia/mainnet/config/ssl/full_node/private_full_node.crt")
			keyFile = fmt.Sprintf("%s%c%s", homeDir, os.PathSeparator, ".chia/mainnet/config/ssl/full_node/private_full_node.key")
		}
	}

	client := &internal.Client{
		CertFile:    certFile,
		KeyFile:     keyFile,
		Fingerprint: *fingerprint,
	}
	client.Run()
}
