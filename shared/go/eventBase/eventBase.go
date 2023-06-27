package eventBase

import (
	"context"
	"math/big"

	"github.com/Entangle-Protocol/off-chain-apps/shared/cfg"
	"github.com/Entangle-Protocol/off-chain-apps/shared/logger"
	"github.com/Entangle-Protocol/off-chain-apps/shared/transactor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TargetNetwork struct {
	Client     *ethclient.Client
	Contracts  map[string]cfg.Contract
	Transactor *bind.TransactOpts
	ChainId    uint64
}

type ListeningNetwork struct {
	Name      string
	Client    *ethclient.Client
	Contracts []cfg.Contract
	Events    []string
	Handler   EventHandlerFunc
	ChainId   uint64
	Target    *TargetNetwork
	// Logger    *log.Logger
}

type ProcessData struct {
	Func       func(*bind.TransactOpts) (string, *types.Transaction, error)
	Transactor *bind.TransactOpts
	Client     *ethclient.Client
}

var targetNetworkDict = make(map[uint64]*TargetNetwork)
var chainIDCache = make(map[string]*big.Int)

type EventHandlerFunc func(*types.Log, ListeningNetwork, chan<- ProcessData)

var Handler handlerInterface
var EventHandlers = make(map[string]EventHandlerFunc)

type handlerInterface interface {
	HandleUSDTTransferEvent(*types.Log, ListeningNetwork, chan<- ProcessData)
	HandleProposalApproved(*types.Log, ListeningNetwork, chan<- ProcessData)
}

func InitializeEventHandlers(handler handlerInterface) {
	EventHandlers = map[string]EventHandlerFunc{
		"handleUSDTTransferEvent": handler.HandleUSDTTransferEvent,
	}
}

func LoadTargetNetwork(config *cfg.TargetNetworkConfig, keystoreFileName string, keystorePassword string) *TargetNetwork {
	// Load client
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		logger.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	var chainID *big.Int
	if cachedChainID, ok := chainIDCache[config.RPCURL]; ok {
		chainID = cachedChainID
	} else {
		chainID, err = client.ChainID(context.Background())
		if err != nil {
			logger.Fatalf("Failed to get ChainID: %v", err)
		}
		chainIDCache[config.RPCURL] = chainID
	}

	// Load transactor
	transactor, err := transactor.LoadTransactor(client, chainID, keystoreFileName, keystorePassword)
	if err != nil {
		logger.Fatalf("Failed to create transactor: %v", err)
	}

	var contractsData map[string]cfg.Contract = make(map[string]cfg.Contract)

	// Load contracts
	for _, contractConfig := range config.Contracts {
		var contractInfo cfg.Contract
		if contractConfig.ABI != "" && contractConfig.Address != "" && contractConfig.Name != "" {
			contractABI, err := cfg.LoadContractABI(contractConfig.ABI)
			if err != nil {
				logger.Fatalf("Failed to load contract ABI: %v", err)
			}

			instance := bind.NewBoundContract(common.HexToAddress(contractConfig.Address), *contractABI, client, client, client)
			contractInfo = cfg.Contract{
				Address:  common.HexToAddress(contractConfig.Address),
				ABI:      contractABI,
				Instance: instance,
			}
		}

		contractsData[contractConfig.Name] = contractInfo
	}

	// Create network
	network := TargetNetwork{
		Client:     client,
		Contracts:  contractsData,
		Transactor: transactor,
		ChainId:    chainID.Uint64(),
	}

	return &network
}

func LoadListeningNetworks(configs []cfg.NetworkConfig, keystoreFilename string, keystorePassword string) []ListeningNetwork {
	var networks []ListeningNetwork
	for _, config := range configs {
		// Load client
		client, err := ethclient.Dial(config.RPCURL)
		if err != nil {
			logger.Fatalf("Failed to connect to the RPC  client: %v", err)
		}

		var chainID *big.Int
		if cachedChainID, ok := chainIDCache[config.RPCURL]; ok {
			chainID = cachedChainID
		} else {
			chainID, err = client.ChainID(context.Background())
			if err != nil {
				logger.Fatalf("Failed to get ChainID: %v", err)
			}
			chainIDCache[config.RPCURL] = chainID
		}

		// Load contracts
		var contractInfo cfg.Contract
		contractConfig := config.Contracts[0]
		if contractConfig.ABI != "" && contractConfig.Address != "" && contractConfig.Name != "" {
			contractABI, err := cfg.LoadContractABI(contractConfig.ABI)
			if err != nil {
				logger.Fatalf("Failed to load contract ABI: %v", err)
			}

			instance := bind.NewBoundContract(common.HexToAddress(contractConfig.Address), *contractABI, client, client, client)
			contractInfo = cfg.Contract{
				Address:  common.HexToAddress(contractConfig.Address),
				ABI:      contractABI,
				Instance: instance,
			}
		}

		var contracts []cfg.Contract
		contracts = append(contracts, contractInfo)

		var target *TargetNetwork
		if config.TargetConfig != nil {
			target = LoadTargetNetwork(config.TargetConfig, keystoreFilename, keystorePassword)
		}

		// Create network
		network := ListeningNetwork{
			Name:      config.Name,
			Client:    client,
			Contracts: contracts,
			Events:    config.ContractEvents,
			Handler:   EventHandlers[config.Handler],
			ChainId:   chainID.Uint64(),
			// Logger:    createCustomLogger(config.LoggerFile),
			Target: target,
		}
		networks = append(networks, network)
	}
	return networks
}

func PrepareTargetNetworkList(configList []*cfg.TargetNetworkConfig, keystoreFilename string, keystorePassword string) []*cfg.TargetNetworkConfig {
	for _, config := range configList {
		network := LoadTargetNetwork(config, keystoreFilename, keystorePassword)
		if network != nil {
			targetNetworkDict[network.ChainId] = network
		}
	}

	return configList
}

func GetTargetNetworkByChainID(destinationChainID uint64) *TargetNetwork {
	return targetNetworkDict[destinationChainID]
}
