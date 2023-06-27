package transactor

import (
	"context"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"

	logger "github.com/Entangle-Protocol/off-chain-apps/shared/logger"
)

var transactorCache = make(map[string]*bind.TransactOpts)
var keyCache = make(map[string]*keystore.Key)

func getKey(keystoreFilePath string, keystorePassword string) (*keystore.Key, error) {
	cacheKey := keystoreFilePath + keystorePassword
	if key, ok := keyCache[cacheKey]; ok {
		return key, nil
	}

	// Read keystore
	keyjson, err := ioutil.ReadFile(keystoreFilePath)
	if err != nil {
		return nil, err
	}

	// Unencrypt keystore
	key, err := keystore.DecryptKey(keyjson, keystorePassword)
	if err != nil {
		return nil, err
	}

	logger.Println("Loaded key for ", key.Address.Hex())

	keyCache[cacheKey] = key
	return key, nil
}

// LoadTransactor read keystore and create TransactOpts
func LoadTransactor(client *ethclient.Client, chainID *big.Int, keystoreFileName string, keystorePassword string) (*bind.TransactOpts, error) {

	keystoreFilePath := "../../shared/keystore/" + keystoreFileName
	if client == nil {
		return nil, logger.Errorf("Client is nil")
	}

	cacheKey := keystoreFilePath + chainID.String()
	if cachedTransactor, ok := transactorCache[cacheKey]; ok {
		// logger.Println("Use cached transactor for ", chainID.String())
		return cachedTransactor, nil
	}

	key, err := getKey(keystoreFilePath, keystorePassword)
	if err != nil {
		return nil, err
	}

	// new TransactOpts instance
	transactor, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	// logger.Println("Create transactor for chain", chainID.String())

	transactor.NoSend = true
	/// To be implemented - calculate GasLimit
	transactor.GasLimit = 10000000

	go updateGasPrice(transactor, client, chainID)

	nonce, err := client.PendingNonceAt(context.Background(), transactor.From)
	if err != nil {
		logger.Fatalf("Failed to get nonce: %v", err)
	}

	nonceBig := big.NewInt(0)
	nonceBig.SetUint64(nonce)
	transactor.Nonce = nonceBig

	transactorCache[cacheKey] = transactor

	return transactor, nil
}

func updateGasPrice(transactor *bind.TransactOpts, client *ethclient.Client, chainID *big.Int) {
	for {
		suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			logger.Fatalf("Failed to get gas price: %v", err)
		}
		transactor.GasPrice = suggestedGasPrice
		// logger.Println("Updated gas price for chain", chainID.String(), ": ", suggestedGasPrice)
		// Update Gasprice every minute
		time.Sleep(time.Minute)
	}
}

func CloneTransactOpts(opts *bind.TransactOpts) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     opts.From,
		Nonce:    new(big.Int).Set(opts.Nonce),
		Signer:   opts.Signer,
		GasPrice: new(big.Int).Set(opts.GasPrice),
		GasLimit: opts.GasLimit,
		Context:  opts.Context,
		NoSend:   opts.NoSend,
	}
}
