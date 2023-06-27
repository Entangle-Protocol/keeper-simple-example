package handlers

import (
	"encoding/hex"
	"fmt"
	"math/big"

	aggregationSpotter "github.com/Entangle-Protocol/off-chain-apps/shared/contracts/aggregationSpotter"
	"github.com/Entangle-Protocol/off-chain-apps/shared/contracts/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cfg "github.com/Entangle-Protocol/off-chain-apps/shared/cfg"
	eventbase "github.com/Entangle-Protocol/off-chain-apps/shared/eventBase"
	logger "github.com/Entangle-Protocol/off-chain-apps/shared/logger"
	"github.com/ethereum/go-ethereum/crypto"
)

type Handlers struct{}

var (
	ApplyStateFunctionSelector = calculateSelector("applyState(bytes)")
	ApplyEventFunctionSelector = calculateSelector("applyEvent(bytes)")
)

type ApplyEventParams struct {
	ChainId        uint64
	BlockNumber    *big.Int
	TxHash         [32]byte
	EventParams    []byte
	EventSignature [4]byte
}

type ApplyStateParams struct {
	ChainId     uint64
	BlockNumber uint64
	StateTypes  *big.Int
	Version     *big.Int
	State       []byte
}

type State struct {
	Sid    *big.Int
	Params []byte
}

type ChefState struct {
	OpTokenBalance *big.Int
	LpTokenBalance *big.Int
}

type StateTypes *big.Int

var (
	Synth StateTypes = big.NewInt(0)
	Dex   StateTypes = big.NewInt(1)
	Chef  StateTypes = big.NewInt(2)
)

type ProposeApplyStateParams struct {
	Params    []byte
	StateType *big.Int

	ProposeParams ProposeOperationParams
}

type ProposeApplyEventParams struct {
	Params         []byte
	TxHash         [32]byte
	EventSignature [4]byte

	ProposeParams ProposeOperationParams
}

type ProposeOperationParams struct {
	EventLog *types.Log
	Network  eventbase.ListeningNetwork
	TxChan   chan<- eventbase.ProcessData

	ContractAddress common.Address
	Selector        [4]byte

	EncodedParams []byte

	SuccessSentText string
	ProcessText     string
}

func calculateSelector(funcSignature string) [4]byte {
	// logger.Println("CalculateSelector for ", funcSignature)
	hash := crypto.Keccak256Hash([]byte(funcSignature))
	firstFourBytes := hash.Bytes()[:4]
	var selectorArray [4]byte
	copy(selectorArray[:], firstFourBytes)

	// logger.Printf("Selector for %s: %x\n", funcSignature, firstFourBytes)
	return selectorArray
}

func CalculateEventSelectorForEvent(eventLog *types.Log) [4]byte {
	// firstFourBytes := eventLog.Topics[0][:4]
	// var selectorArray [4]byte
	// copy(selectorArray[:], firstFourBytes)
	// logger.Printf("Selector For Event: %x\n", firstFourBytes)
	// return selectorArray
	var selector [4]byte
	if len(eventLog.Topics) > 0 {
		eventSig := eventLog.Topics[0]
		var eventSigBytes [32]byte
		copy(eventSigBytes[:], eventSig[:])

		copy(selector[:], eventSigBytes[0:4])

		logger.Printf("Event Selector: %x\n", selector)
	} else {
		logger.Fatalf("No topics present in the log.")
	}

	return selector
}

func ProposeOperationApplyEvent(applyParams ProposeApplyEventParams) {

	applyEventParams := ApplyEventParams{
		ChainId:        applyParams.ProposeParams.Network.ChainId,
		BlockNumber:    big.NewInt(0).SetUint64(applyParams.ProposeParams.EventLog.BlockNumber),
		TxHash:         applyParams.TxHash,
		EventParams:    applyParams.Params,
		EventSignature: applyParams.EventSignature,
	}

	logger.Println("ChainId ", applyEventParams.ChainId)
	logger.Println("BlockNumber ", applyEventParams.BlockNumber)
	logger.Println("TxHash ", hex.EncodeToString(applyEventParams.TxHash[:]))
	logger.Println("EventParams ", hex.EncodeToString(applyEventParams.EventParams))
	logger.Println("EventSignature ", hex.EncodeToString(applyEventParams.EventSignature[:]))
	// bytes, err := hex.DecodeString(hexString)
	// if err != nil {
	// 	// Обработка ошибок
	// }
	// ABI args
	// arguments := abi.Arguments{
	// 	{
	// 		Type: cfg.Uint64Type, //  ChainId
	// 	},
	// 	{
	// 		Type: cfg.Uint256Type, //  BlockNumber
	// 	},
	// 	{
	// 		Type: cfg.Bytes32Type, // StateTypes
	// 	},
	// 	{
	// 		Type: cfg.BytesType, //  State
	// 	},
	// 	{
	// 		Type: cfg.Bytes4Type, //  Sig
	// 	},
	// }
	structApplyEventParams, _ := abi.NewType("tuple", "structApplyEventParams", []abi.ArgumentMarshaling{
		{Name: "ChainId", Type: "uint64"},
		{Name: "BlockNumber", Type: "uint256"},
		{Name: "TxHash", Type: "bytes32"},
		{Name: "EventParams", Type: "bytes"},
		{Name: "EventSignature", Type: "bytes4"},
	})

	args := abi.Arguments{
		{Type: structApplyEventParams, Name: "instance"},
	}

	applyEventParamsEncoded, err := args.Pack(&applyEventParams)
	// bytes := []byte{211, 96, 232, 88}
	// applyEventParamsEncoded, err := arguments.Pack(applyEventParams.ChainId, applyEventParams.BlockNumber, applyEventParams.TxHash, applyEventParams.EventParams, applyEventParams.EventSignature)
	if err != nil {
		logger.Fatalf("Failed to pack applyEventParams: %v", err)
	}

	str := hex.EncodeToString(applyEventParamsEncoded)
	logger.Printf("applyEventParamsEncoded: %v\n", str)

	applyParams.ProposeParams.Selector = ApplyEventFunctionSelector
	applyParams.ProposeParams.EncodedParams = applyEventParamsEncoded

	proposeOperation(applyParams.ProposeParams)
}

func proposeOperationApplyState(applyParams ProposeApplyStateParams) {

	StateParams := ApplyStateParams{
		ChainId:     applyParams.ProposeParams.Network.ChainId,
		BlockNumber: applyParams.ProposeParams.EventLog.BlockNumber,
		StateTypes:  applyParams.StateType,
		Version:     big.NewInt(1),
		State:       applyParams.Params,
	}

	// Описание аргументов функции в формате ABI
	arguments := abi.Arguments{
		{
			Type: cfg.Uint64Type, //  ChainId
		},
		{
			Type: cfg.Uint256Type, //  BlockNumber
		},
		{
			Type: cfg.Uint256Type, // StateTypes
		},
		{
			Type: cfg.Uint256Type, // Version
		},
		{
			Type: cfg.BytesType, //  State
		},
	}

	applyStateParamsEncoded, err := arguments.Pack(StateParams.ChainId, big.NewInt(0).SetUint64(StateParams.BlockNumber), &StateParams.StateTypes, &StateParams.Version, StateParams.State)
	if err != nil {
		logger.Fatalf("Failed to pack data: %v", err)
	}

	str := hex.EncodeToString(applyStateParamsEncoded)
	logger.Printf("applyStateParamsEncoded: %v\n", str)

	applyParams.ProposeParams.Selector = ApplyStateFunctionSelector
	applyParams.ProposeParams.EncodedParams = applyStateParamsEncoded

	proposeOperation(applyParams.ProposeParams)

}

func proposeOperation(params ProposeOperationParams) {

	logger.Printf("ProposeOperation %d params.ContractAddress: %s\n", params.Network.Target.ChainId, params.ContractAddress.Hex())

	operationData := aggregationSpotter.AggregationSpotterOperationData{
		Contr:            params.ContractAddress,
		FunctionSelector: params.Selector,
		Params:           params.EncodedParams,
	}

	logger.LogProposeOperation(operationData)

	var OpTxIdBytes [32]byte
	byteArr := params.EventLog.TxHash.Bytes()
	copy(OpTxIdBytes[32-len(byteArr):], byteArr)

	aggregationSpotterContract, err := aggregationSpotter.NewAggregationSpotter(params.Network.Target.Contracts["AggregationSpotter"].Address, params.Network.Target.Client)
	if err != nil {
		logger.Fatalf("Failed to bind AggregationSpotter contract: %v", err)
	}

	proposeOperationFunc := func(transactor *bind.TransactOpts) (string, *types.Transaction, error) {
		// Run proposeOperation for AggregationSpotter

		logger.Println(params.ProcessText)
		tx, err := aggregationSpotterContract.ProposeOperation(transactor, OpTxIdBytes, operationData)
		return params.SuccessSentText, tx, err
	}

	// Send proposeOperation to txChan
	processData := eventbase.ProcessData{
		Func:       proposeOperationFunc,
		Transactor: params.Network.Target.Transactor,
		Client:     params.Network.Target.Client,
	}

	params.TxChan <- processData
}

func (h *Handlers) HandleProposalApproved(eventLog *types.Log, network eventbase.ListeningNetwork, txChan chan<- eventbase.ProcessData) {
}

func InitializeHandlers() {
	var HandlersInstance = &Handlers{}

	eventbase.InitializeEventHandlers(HandlersInstance)
}

func (h *Handlers) HandleUSDTTransferEvent(eventLog *types.Log, network eventbase.ListeningNetwork, txChan chan<- eventbase.ProcessData) {

	contract, err := erc20.NewErc20(network.Contracts[0].Address, network.Client)
	if err != nil {
		logger.Fatalf("Failed to bind ERC20 contract: %v", err)
	}

	var Sid *big.Int = big.NewInt(0)
	var eventName string = ""
	var amount *big.Int = big.NewInt(0)
	var addressFromOrTo string
	var opId *big.Int = big.NewInt(0)
	for _, event := range network.Contracts[0].ABI.Events {
		if eventLog.Topics[0].Hex() == event.ID.Hex() {
			eventName = event.Name
			if event.Name == "Transfer" {
				eventData, err := contract.ParseTransfer(*eventLog)
				if err != nil {
					logger.Fatalf("Failed parse Transfer event: %v", err)
				}

				Sid = eventData.Sid
				amount = eventData.Amount
				addressFromOrTo = eventData.From.String()
				opId = eventData.OpId
			}
			break
		}
	}

	logger.Printf("Handle %s event with name %s: sid %d, amount %d, address %s, opID %d\n", network.Name, eventName, Sid, amount, addressFromOrTo, opId)

	totalSupply, err := contract.GetTotalSupply(nil, Sid)
	if err != nil {
		logger.Fatalf("Failed to get total supply: %v", err)
	}

	price, err := contract.GetPrice(nil, Sid)
	if err != nil {
		logger.Fatalf("Failed to get price: %v", err)
	}

	logger.Printf("Handle %s event with name %s: Factory TotalSupply %d, Price %d\n", network.Name, eventName, totalSupply, price)

	// OperationData for transaction
	params := stateSpotter.StateParamsLibraryESynthStateParams{Price: price, Supply: totalSupply}
	arguments := abi.Arguments{
		{
			Type: cfg.Uint256Type,
		},
		{
			Type: cfg.Uint256Type,
		},
	}

	encodedParams, err := arguments.Pack(params.Price, params.Supply)
	if err != nil {
		logger.Fatalf("Failed to pack params: %v", err)
	}

	state := State{Sid: Sid, Params: encodedParams}
	arguments = abi.Arguments{
		{
			Type: cfg.Uint128Type,
		},
		{
			Type: cfg.BytesType,
		},
	}

	encodedState, err := arguments.Pack(state.Sid, state.Params)
	if err != nil {
		logger.Fatalf("Failed to pack encodedState data: %v", err)
	}

	eventInfoText := fmt.Sprintf("'%v': block %d nonce %d txId %v [sid %d, amount %d, address %v, opId %d]", eventName, eventLog.BlockNumber, network.Target.Transactor.Nonce, eventLog.TxHash.String(), Sid, amount, addressFromOrTo, opId)
	processText := fmt.Sprintf("ApplyState processed for %v %v", network.Name, eventInfoText)
	sussecSentText := fmt.Sprintf("ApplyState success sent for %v %v", network.Name, eventInfoText)

	proposeOperationParams := ProposeOperationParams{
		EventLog:        eventLog,
		Network:         network,
		TxChan:          txChan,
		SuccessSentText: sussecSentText,
		ProcessText:     processText,
		ContractAddress: network.Target.Contracts["StateSpotter"].Address,
	}

	proposeOperationApplyState(
		ProposeApplyStateParams{
			Params:        encodedState,
			StateType:     Synth,
			ProposeParams: proposeOperationParams,
		})

	if opId.Cmp(big.NewInt(0)) == 0 {
		logger.Println("Event opId is 0, skip send ApplyEvent")
	} else {
		logger.Println("Event opID is not 0, need send ApplyEvent")

		proposeOperationParams.ProcessText = fmt.Sprintf("ApplyEvent processed for %v  %v", network.Name, eventInfoText)
		proposeOperationParams.SuccessSentText = fmt.Sprintf("ApplyEvent success sent for %v %v", network.Name, eventInfoText)
		proposeOperationParams.ContractAddress = network.Target.Contracts["BalanceSpotter"].Address

		ProposeOperationApplyEvent(
			ProposeApplyEventParams{
				Params:         eventLog.Data,
				TxHash:         eventLog.TxHash,
				EventSignature: CalculateEventSelectorForEvent(eventLog),
				ProposeParams:  proposeOperationParams,
			})

		if addressFromOrTo == network.Target.Contracts["SynthDex"].Address.String() {
			h.HandleSynthDexEvent(eventLog, network, txChan)
		}
	}
}
