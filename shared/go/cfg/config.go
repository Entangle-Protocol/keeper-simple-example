package cfg

import (
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"

	logger "github.com/Entangle-Protocol/off-chain-apps/shared/logger"
)

type Config struct {
	ListeningData []NetworkConfig        `yaml:"ListeningData"`
	TargetList    []*TargetNetworkConfig `yaml:"TargetList"`
}

type TargetNetworkConfig struct {
	RPCURL    string           `yaml:"rpcURL"`
	ChainID   *big.Int         `yaml:"chainID"`
	Contracts []ContractConfig `yaml:"contracts"`
}

type ContractConfig struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	ABI     string `yaml:"abi"`
}

type NetworkConfig struct {
	Name           string               `yaml:"name"`
	RPCURL         string               `yaml:"rpcURL"`
	Handler        string               `yaml:"handler"`
	ContractEvents []string             `yaml:"contractEvents"`
	Contracts      []ContractConfig     `yaml:"contracts"`
	TargetConfig   *TargetNetworkConfig `yaml:"targetNetwork"`
	// LoggerFile     string               `yaml:"logger"`
}

// --------------------------------------------
type Contract struct {
	Address  common.Address
	ABI      *abi.ABI
	Instance *bind.BoundContract
}

type Bytes32 [32]byte

var Uint256ArrayType abi.Type
var Uint256Type abi.Type
var Uint128Type abi.Type
var Uint64Type abi.Type
var Bytes32Type abi.Type
var Bytes4Type abi.Type
var BytesType abi.Type

func init() {
	var err error

	Uint256ArrayType, err = abi.NewType("uint256[]", "", nil)
	if err != nil {
		logger.Fatalf("Failed to create uint256 type: %v", err)
	}

	Uint256Type, err = abi.NewType("uint256", "", nil)
	if err != nil {
		logger.Fatalf("Failed to create uint256 type: %v", err)
	}

	Uint128Type, err = abi.NewType("uint128", "", nil)
	if err != nil {
		logger.Fatalf("Failed to create uint128 type: %v", err)
	}

	Uint64Type, err = abi.NewType("uint64", "", nil)
	if err != nil {
		logger.Fatalf("Failed to create new ABI type: %v", err)
	}

	Bytes4Type, err = abi.NewType("bytes4", "", nil)
	if err != nil {
		logger.Fatalf("Failed to create bytes type: %v", err)
	}

	Bytes32Type, err = abi.NewType("bytes32", "", nil)
	if err != nil {
		logger.Fatalf("Failed to create bytes type: %v", err)
	}

	BytesType, err = abi.NewType("bytes", "", nil)
	if err != nil {
		logger.Fatalf("Failed to create bytes type: %v", err)
	}
}

func LoadContractABI(path string) (*abi.ABI, error) {
	contractABI, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, logger.Errorf("failed to read contract ABI file: %v", err)
	}
	contract, err := abi.JSON(strings.NewReader(string(contractABI)))
	if err != nil {
		return nil, logger.Errorf("failed to parse contract ABI: %v", err)
	}
	return &contract, nil
}

func ParseConfigFile(filename string) Config {
	// Read config file
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		logger.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	// Parse config file
	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		logger.Fatalf("Failed to parse config file: %v", err)
	}

	return config
}
