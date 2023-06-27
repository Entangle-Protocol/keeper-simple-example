package events

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	eventbase "github.com/Entangle-Protocol/off-chain-apps/shared/eventBase"
	logger "github.com/Entangle-Protocol/off-chain-apps/shared/logger"
)

func ListenForEvent(network eventbase.ListeningNetwork, txChan chan<- eventbase.ProcessData) error {
	eventStringList := ""
	for _, event := range network.Events {
		eventABI := network.Contracts[0].ABI.Events[event]

		query := ethereum.FilterQuery{
			Addresses: []common.Address{network.Contracts[0].Address},
			Topics:    [][]common.Hash{{eventABI.ID}},
		}

		eventStringList += eventABI.Name + " "

		// Create a channel to receive the events
		events := make(chan types.Log)

		sub, err := network.Client.SubscribeFilterLogs(context.Background(), query, events)
		if err != nil {
			logger.Fatalf("Failed to subscribe to contract events: %v", err)
		}

		// Handle events in a goroutine
		go func(network eventbase.ListeningNetwork, txChan chan<- eventbase.ProcessData) {
			for {
				select {
				case err := <-sub.Err():
					logger.Fatalf("Got error from event subscription: %v", err)

				case eventLog := <-events:
					logger.Printf("Got event for %s\n", network.Name)
					network.Handler(&eventLog, network, txChan)
				}
			}
		}(network, txChan)
	}
	logger.Printf("Listening %s for events %s\n", network.Name, eventStringList)

	return nil
}
