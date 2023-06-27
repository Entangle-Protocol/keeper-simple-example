package proccessing

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	eventbase "github.com/Entangle-Protocol/off-chain-apps/shared/eventBase"
	logger "github.com/Entangle-Protocol/off-chain-apps/shared/logger"
	transactor "github.com/Entangle-Protocol/off-chain-apps/shared/transactor"
)

func sendTransaction(signedTx *types.Transaction, client *ethclient.Client, successText string) {
	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Errorf("Failed to send transaction: %v", err)
	}

	logger.Printf("Proccessing: Sent transaction %s to chain %d\n", signedTx.Hash().Hex(), signedTx.ChainId())

	for {
		receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			if err.Error() == "not found" {
				time.Sleep(5 * time.Second)
				continue
			} else {
				logger.Errorf("Proccessing: Failed to get transaction receipt: %v", err)
				return
			}
		}

		if receipt.Status == types.ReceiptStatusFailed {
			logger.Errorf("Proccessing: Transaction failed")
		} else {
			logger.Printf("Proccessing: Success mined tx %s\n - %s", signedTx.Hash().Hex(), successText)
			// logger.Println(successText)
		}
		break
	}
}

func ProccessTransactions(txChan <-chan eventbase.ProcessData) {
	// Loop forever, signing and sending transactions as they come in
	for data := range txChan {
		// uniq TransactOpts instance for every transaction
		newTransactor := transactor.CloneTransactOpts(data.Transactor)

		succesText, signedTx, err := data.Func(newTransactor)
		if err != nil {
			logger.Fatalf("Failed to create transaction: %v", err)
		}

		// Increase nonce
		data.Transactor.Nonce.Add(data.Transactor.Nonce, big.NewInt(1))

		// send transaction
		go sendTransaction(signedTx, data.Client, succesText)
	}
}
