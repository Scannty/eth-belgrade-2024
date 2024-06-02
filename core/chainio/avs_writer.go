package chainio

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewTaskSendBTC(
		ctx context.Context,
		destAddress string,
		amount *big.Int,
		quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
		quorumNumbers sdktypes.QuorumNums,
	) (cstaskmanager.IIncredibleSquaringTaskManagerTask, uint32, error)
	// RaiseChallenge(
	// 	ctx context.Context,
	// 	task cstaskmanager.IIncredibleSquaringTaskManagerTask,
	// 	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	// 	taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
	// 	pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
	// ) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		task cstaskmanager.IIncredibleSquaringTaskManagerTask,
		taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
		nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.Client
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.IncredibleSquaringRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}
func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewTaskSendBTC(ctx context.Context, destAddress string, amount *big.Int, quorumThresholdPercentage sdktypes.QuorumThresholdPercentage, quorumNumbers sdktypes.QuorumNums) (cstaskmanager.IIncredibleSquaringTaskManagerTask, uint32, error) {
	// GET TX OPTIONS
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}

	// GET WBTC CONTRACT
	wbtc, err := w.AvsContractBindings.GetWBTC(common.HexToAddress("0x5f3f1dBD7B74C6B46e8c44f98792A1dAf8d69154"))
	if err != nil {
		w.logger.Errorf("Error getting the WBTC contract :((")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}

	// MINT WBTC TO USER
	txMint, errMint := wbtc.Mint(txOpts, txOpts.From, big.NewInt(100))
	if err != nil {
		w.logger.Errorf("Error creating mint WBTC tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, errMint
	}
	_, errMintSend := w.TxMgr.Send(ctx, txMint)
	if err != nil {
		w.logger.Errorf("Error submitting WBTC tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, errMintSend
	}

	// CREATE NEW TASK
	tx, err := w.AvsContractBindings.TaskManager.CreateNewTask(txOpts, destAddress, amount, uint32(quorumThresholdPercentage), quorumNumbers.UnderlyingType())
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		fmt.Println(err)
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	//w.logger.Error("Aggregator failed to parse new task created event", "err", receipt)
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractIncredibleSquaringTaskManagerFilterer.ParseNewTaskCreated(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task cr", "err", err)
		fmt.Println("ssss")
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}

	return newTaskCreatedEvent.Task, newTaskCreatedEvent.TaskIndex, nil

	// // BURN WBTC
	// txBurn, errBurn := wbtc.Burn(txOpts, txOpts.From, big.NewInt(100), []byte{0,4,1})
	// if errBurn != nil {
	// 	w.logger.Errorf("Error creating burn WBTC tx")
	// 	return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, errBurn
	// }
	// _, errBurnSend := w.TxMgr.Send(ctx, txBurn)
	// if err != nil {
	// 	w.logger.Errorf("Error submitting WBTC tx")
	// 	return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, errBurnSend
	// }
}

func (w *AvsWriter) SendAggregatedResponse(
	ctx context.Context, task cstaskmanager.IIncredibleSquaringTaskManagerTask,
	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RespondToTask(txOpts, task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting respondToTask tx")
		return nil, err
	}
	return receipt, nil
}

// func (w *AvsWriter) RaiseChallenge(
// 	ctx context.Context,
// 	task cstaskmanager.IIncredibleSquaringTaskManagerTask,
// 	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
// 	taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
// 	pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
// ) (*types.Receipt, error) {
// 	txOpts, err := w.TxMgr.GetNoSendTxOpts()
// 	if err != nil {
// 		w.logger.Errorf("Error getting tx opts")
// 		return nil, err
// 	}
// 	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveChallenge(txOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
// 	if err != nil {
// 		w.logger.Errorf("Error assembling RaiseChallenge tx")
// 		return nil, err
// 	}
// 	receipt, err := w.TxMgr.Send(ctx, tx)
// 	if err != nil {
// 		w.logger.Errorf("Error submitting RaiseChallenge tx")
// 		return nil, err
// 	}
// 	return receipt, nil
// }
