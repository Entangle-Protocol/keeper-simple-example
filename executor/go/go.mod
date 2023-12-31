module github.com/Entangle-Protocol/off-chain-apps/executor

go 1.19

require github.com/ethereum/go-ethereum v1.12.0

require gopkg.in/yaml.v2 v2.4.0 // indirect

require (
	github.com/Entangle-Protocol/off-chain-apps/shared v0.0.0
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/deckarep/golang-set/v2 v2.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/holiman/uint256 v1.2.2-0.20230321075855-87b91420868c // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)

replace (
	// github.com/Entangle-Protocol/off-chain-apps/shared/contracts/aggregationSpotter => ../../shared/go/contracts/aggregationSpotter

	// github.com/Entangle-Protocol/off-chain-apps/shared/contracts/entangleDex => ../../shared/go/contracts/entangleDex

	// github.com/Entangle-Protocol/off-chain-apps/shared/contracts/stateSpotter => ../../shared/go/contracts/stateSpotter

	// github.com/Entangle-Protocol/off-chain-apps/shared/contracts/masterSynthChef => ../../shared/go/contracts/masterSynthChef

	// github.com/Entangle-Protocol/off-chain-apps/shared/contracts/balanceSpotter => ../../shared/go/contracts/balanceSpotter

	// github.com/Entangle-Protocol/off-chain-apps/shared/logger => ../../shared/go/logger

	// github.com/Entangle-Protocol/off-chain-apps/shared/proccessing => ../../shared/go/proccessing

	// github.com/Entangle-Protocol/off-chain-apps/shared/cfg => ../../shared/go/cfg

	// github.com/Entangle-Protocol/off-chain-apps/shared/transactor => ../../shared/go/transactor

	// github.com/Entangle-Protocol/off-chain-apps/shared/events => ../../shared/go/events

	// github.com/Entangle-Protocol/off-chain-apps/shared/eventBase => ../../shared/go/eventBase

	github.com/Entangle-Protocol/off-chain-apps/executor/handlers => ./handlers
	github.com/Entangle-Protocol/off-chain-apps/shared => ../../shared/go
)
