# go-chia (WIP)
Golang RPC API to the Chia Node

All the Chia processes is only available to local connection with the default ports below.  Also in order access the RPC API it also need access to the chia full node certificate and private keys.

Chia Process Default Ports

    Full Node: 8555
    Farmer:    8559
    Harvester: 8560
    Wallet:    9256

Default location of the Chia Full Node Certificate and Private Keys

	~/.chia/mainnet/config/ssl/full_node/private_full_node.crt
	~/.chia/mainnet/config/ssl/full_node/private_full_node.key
