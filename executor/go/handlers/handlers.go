package handlers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Entangle-Protocol/off-chain-apps/shared/contracts/aggregationSpotter"
	eventbase "github.com/Entangle-Protocol/off-chain-apps/shared/eventBase"
	logger "github.com/Entangle-Protocol/off-chain-apps/shared/logger"
)

type Handlers struct{}

func (h *Handlers) HandleUSDTTransferEvent(eventLog *types.Log, network eventbase.ListeningNetwork, txChan chan<- eventbase.ProcessData) {
}

func InitializeHandlers() {
	var HandlersInstance = &Handlers{}

	eventbase.InitializeEventHandlers(HandlersInstance)
}

func (h *Handlers) HandleProposalApproved(eventLog *types.Log, network eventbase.ListeningNetwork, txChan chan<- eventbase.ProcessData) {

	contract, err := aggregationSpotter.NewAggregationSpotter(network.Contracts[0].Address, network.Target.Client)
	if err != nil {
		logger.Fatalf("Failed to bind AggregationSpotter contract: %v", err)
	}

	eventName := ""
	var opHash *big.Int = big.NewInt(0)
	for _, event := range network.Contracts[0].ABI.Events {
		if eventLog.Topics[0].Hex() == event.ID.Hex() {
			eventName = event.Name
			if event.Name == "ProposalApproved" {
				eventData, err := contract.ParseProposalApproved(*eventLog)
				if err != nil {
					logger.Fatalf("Failed parese Approve event: %v", err)
				}

				opHash = eventData.OpHash
			}
			break
		}
	}

	logger.Printf("Handle %s event with name %s, opHash %s\n", network.Name, eventName, opHash)

	eventInfoText := fmt.Sprintf("'%v': block %d nonce %d txId %v [opHash %d]", eventName, eventLog.BlockNumber, network.Target.Transactor.Nonce, eventLog.TxHash.String(), opHash)
	processText := fmt.Sprintf("%v tx processed %v", network.Name, eventInfoText)
	successSentText := fmt.Sprintf("%v success sent %v", network.Name, eventInfoText)

	executeOperationFunc := func(transactor *bind.TransactOpts) (string, *types.Transaction, error) {
		// Запуск транзакции proposeOperation в контракте AggregationSpotter

		logger.Println(processText)
		tx, err := contract.ExecuteOperation(transactor, opHash)
		return successSentText, tx, err
	}

	// Отправьте транзакцию в канал
	processData := eventbase.ProcessData{
		Func:       executeOperationFunc,
		Transactor: network.Target.Transactor,
		Client:     network.Target.Client,
	}

	txChan <- processData
}
