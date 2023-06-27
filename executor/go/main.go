package main

import (
	"os"

	handlers "github.com/Entangle-Protocol/off-chain-apps/executor/handlers"
	cfg "github.com/Entangle-Protocol/off-chain-apps/shared/cfg"
	eventbase "github.com/Entangle-Protocol/off-chain-apps/shared/eventBase"
	events "github.com/Entangle-Protocol/off-chain-apps/shared/events"
	logger "github.com/Entangle-Protocol/off-chain-apps/shared/logger"
	proccessing "github.com/Entangle-Protocol/off-chain-apps/shared/proccessing"
)

func main() {
	logger.Println("Start Executor")

	args := os.Args[1:]
	if len(args) < 1 {
		logger.Fatal("Need arg keystore file name")
	}

	if len(args) < 2 {
		logger.Fatal("Need arg keystore password in params")
	}

	keystoreFileName := args[0]
	keystorePassword := args[1]

	// Parse config file
	config := cfg.ParseConfigFile("../config.yaml")

	handlers.InitializeHandlers()

	// Load target networks
	logger.Printf("Prepearing TargetNetworkList...\n")

	eventbase.PrepareTargetNetworkList(config.TargetList, keystoreFileName, keystorePassword)

	logger.Println("Prepearing TargetNetworkList - Success!")

	logger.Printf("Loading ListeningNetworks...")
	// Load listening networks
	listeningNetworks := eventbase.LoadListeningNetworks(config.ListeningData, keystoreFileName, keystorePassword)

	logger.Println(" - Success!")

	// Buffer size 100, adjust as needed
	txChan := make(chan eventbase.ProcessData, 100)

	// Start the transaction signer and sender
	go proccessing.ProccessTransactions(txChan)

	logger.Printf("Total %d contracts to listen\n", len(listeningNetworks))
	// Start listening to events
	logger.Println("Start to listening all events at every networks...")
	for _, network := range listeningNetworks {
		// go network.listenForEvents()
		go events.ListenForEvent(network, txChan)
	}

	// Wait indefinitely
	select {}
}
