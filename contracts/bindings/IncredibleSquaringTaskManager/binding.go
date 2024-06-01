// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSquaringTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IIncredibleSquaringTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTask struct {
	DestAddress               common.Address
	Amount                    *big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IIncredibleSquaringTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	TxHash             common.Address
	PublicKey          common.Address
}

// IIncredibleSquaringTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractIncredibleSquaringTaskManagerMetaData contains all meta data concerning the ContractIncredibleSquaringTaskManager contract.
var ContractIncredibleSquaringTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"destAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_wbtc\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"destAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"txHash\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKey\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wbtcAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"destAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"txHash\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKey\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005707380380620057078339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615417620002f060003960008181610290015281816105af0152612b54015260008181610578015261217a01526000818161041e015261235c0152600081816104580152818161253201526126f401526000818161047f01528181610f6a01528181611e4501528181611fdd015261221701526154176000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c806368304835116101255780638da5cb5b116100ad578063df5cf7231161007c578063df5cf72314610573578063f2fde38b1461059a578063f5c9899d146105ad578063f63c5bab146105d3578063fabc1cbc146105db57600080fd5b80638da5cb5b14610521578063b98d090814610532578063c9689f171461053f578063cefdc1d41461055257600080fd5b806372d18e8d116100f457806372d18e8d146104ca578063766c698b146104d85780637afa1eed146104eb578063886f1195146104fe5780638b00ce7c1461051157600080fd5b806368304835146104535780636d14a9871461047a5780636efb4636146104a1578063715018a6146104c257600080fd5b8063416c7e5e116101a85780635c155662116101775780635c155662146103ce5780635c975abb146103ee5780635decc3f5146103f65780635df459461461041957806361a2ac391461044057600080fd5b8063416c7e5e146103605780634f739f7414610373578063595c6a67146103935780635ac86ab71461039b57600080fd5b80631ad43189116101ef5780631ad431891461028b578063245a7bfc146102c75780632cb223d5146102f25780632d89f6fc146103205780633563b0d11461034057600080fd5b806310d67a2f14610221578063136439dd146102365780631459457a14610249578063171f1d5b1461025c575b600080fd5b61023461022f366004613faf565b6105ee565b005b610234610244366004613fcc565b6106aa565b610234610257366004613fe5565b6107e9565b61026f61026a3660046141bb565b610948565b6040805192151583529015156020830152015b60405180910390f35b6102b27f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610282565b60cd546102da906001600160a01b031681565b6040516001600160a01b039091168152602001610282565b610312610300366004614229565b60cb6020526000908152604090205481565b604051908152602001610282565b61031261032e366004614229565b60ca6020526000908152604090205481565b61035361034e366004614246565b610ad2565b60405161028291906143a1565b61023461036e3660046143c9565b610f68565b61038661038136600461442e565b6110dd565b6040516102829190614532565b610234611803565b6103be6103a93660046145fc565b606654600160ff9092169190911b9081161490565b6040519015158152602001610282565b6103e16103dc36600461463c565b6118ca565b60405161028291906146ed565b606654610312565b6103be610404366004614229565b60cc6020526000908152604090205460ff1681565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b60cf546102da906001600160a01b031681565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6104b46104af3660046149ad565b611a92565b604051610282929190614a6d565b6102346129a9565b60c95463ffffffff166102b2565b6102346104e6366004614ab6565b6129bd565b60ce546102da906001600160a01b031681565b6065546102da906001600160a01b031681565b60c9546102b29063ffffffff1681565b6033546001600160a01b03166102da565b6097546103be9060ff1681565b61023461054d366004614b3e565b612cec565b610565610560366004614bb0565b612e70565b604051610282929190614bf2565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6102346105a8366004613faf565b613002565b7f00000000000000000000000000000000000000000000000000000000000000006102b2565b6102b2606481565b6102346105e9366004613fcc565b613078565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610641573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106659190614c13565b6001600160a01b0316336001600160a01b03161461069e5760405162461bcd60e51b815260040161069590614c30565b60405180910390fd5b6106a7816131d4565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107169190614c7a565b6107325760405162461bcd60e51b815260040161069590614c97565b606654818116146107ab5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156108095750600054600160ff909116105b806108235750303b158015610823575060005460ff166001145b6108865760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610695565b6000805460ff1916600117905580156108a9576000805461ff0019166101001790555b6108b48660006132cb565b6108bd856133b5565b60cd80546001600160a01b038087166001600160a01b03199283161790925560ce805486841690831617905560cf8054928516929091169190911790558015610940576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061099057610990614cdf565b60200201518951600160200201518a602001516000600281106109b5576109b5614cdf565b60200201518b602001516001600281106109d1576109d1614cdf565b602090810291909101518c518d830151604051610a2e9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610a519190614cf5565b9050610ac4610a6a610a638884613407565b869061349e565b610a72613532565b610aba610aab85610aa5604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613407565b610ab48c6135f2565b9061349e565b886201d4c0613682565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b389190614c13565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9e9190614c13565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610be0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c049190614c13565b9050600086516001600160401b03811115610c2157610c21614056565b604051908082528060200260200182016040528015610c5457816020015b6060815260200190600190039081610c3f5790505b50905060005b8751811015610f5c576000888281518110610c7757610c77614cdf565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610cd8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d009190810190614d17565b905080516001600160401b03811115610d1b57610d1b614056565b604051908082528060200260200182016040528015610d6657816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610d395790505b50848481518110610d7957610d79614cdf565b602002602001018190525060005b8151811015610f46576040518060600160405280876001600160a01b03166347b314e8858581518110610dbc57610dbc614cdf565b60200260200101516040518263ffffffff1660e01b8152600401610de291815260200190565b602060405180830381865afa158015610dff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e239190614c13565b6001600160a01b03168152602001838381518110610e4357610e43614cdf565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610e7157610e71614cdf565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610ecd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef19190614da7565b6001600160601b0316815250858581518110610f0f57610f0f614cdf565b60200260200101518281518110610f2857610f28614cdf565b60200260200101819052508080610f3e90614de6565b915050610d87565b5050508080610f5490614de6565b915050610c5a565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fea9190614c13565b6001600160a01b0316336001600160a01b0316146110965760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610695565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6111086040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611148573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061116c9190614c13565b90506111996040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906111c9908b9089908990600401614e01565b600060405180830381865afa1580156111e6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261120e9190810190614e4b565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611240908b908b908b90600401614f02565b600060405180830381865afa15801561125d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112859190810190614e4b565b6040820152856001600160401b038111156112a2576112a2614056565b6040519080825280602002602001820160405280156112d557816020015b60608152602001906001900390816112c05790505b50606082015260005b60ff8116871115611714576000856001600160401b0381111561130357611303614056565b60405190808252806020026020018201604052801561132c578160200160208202803683370190505b5083606001518360ff168151811061134657611346614cdf565b602002602001018190525060005b868110156116145760008c6001600160a01b03166304ec63518a8a8581811061137f5761137f614cdf565b905060200201358e8860000151868151811061139d5761139d614cdf565b60200260200101516040518463ffffffff1660e01b81526004016113da9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156113f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141b9190614f2b565b90506001600160c01b0381166114bf5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610695565b8a8a8560ff168181106114d4576114d4614cdf565b6001600160c01b03841692013560f81c9190911c60019081161415905061160157856001600160a01b031663dd9846b98a8a8581811061151657611516614cdf565b905060200201358d8d8860ff1681811061153257611532614cdf565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611588573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ac9190614f54565b85606001518560ff16815181106115c5576115c5614cdf565b602002602001015184815181106115de576115de614cdf565b63ffffffff90921660209283029190910190910152826115fd81614de6565b9350505b508061160c81614de6565b915050611354565b506000816001600160401b0381111561162f5761162f614056565b604051908082528060200260200182016040528015611658578160200160208202803683370190505b50905060005b828110156116d95784606001518460ff168151811061167f5761167f614cdf565b6020026020010151818151811061169857611698614cdf565b60200260200101518282815181106116b2576116b2614cdf565b63ffffffff90921660209283029190910190910152806116d181614de6565b91505061165e565b508084606001518460ff16815181106116f4576116f4614cdf565b60200260200101819052505050808061170c90614f71565b9150506112de565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611755573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117799190614c13565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906117ac908b908b908e90600401614f91565b600060405180830381865afa1580156117c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117f19190810190614e4b565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561184b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061186f9190614c7a565b61188b5760405162461bcd60e51b815260040161069590614c97565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016118fc929190614fbb565b600060405180830381865afa158015611919573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119419190810190614e4b565b9050600084516001600160401b0381111561195e5761195e614056565b604051908082528060200260200182016040528015611987578160200160208202803683370190505b50905060005b8551811015611a8857866001600160a01b03166304ec63518783815181106119b7576119b7614cdf565b6020026020010151878685815181106119d2576119d2614cdf565b60200260200101516040518463ffffffff1660e01b8152600401611a0f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611a2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a509190614f2b565b6001600160c01b0316828281518110611a6b57611a6b614cdf565b602090810291909101015280611a8081614de6565b91505061198d565b5095945050505050565b6040805180820190915260608082526020820152600084611b095760405162461bcd60e51b815260206004820152603760248201526000805160206153c283398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610695565b60408301515185148015611b21575060a08301515185145b8015611b31575060c08301515185145b8015611b41575060e08301515185145b611bab5760405162461bcd60e51b815260206004820152604160248201526000805160206153c283398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610695565b82515160208401515114611c235760405162461bcd60e51b8152602060048201526044602482018190526000805160206153c2833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610695565b4363ffffffff168463ffffffff1610611c925760405162461bcd60e51b815260206004820152603c60248201526000805160206153c283398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610695565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611cd357611cd3614056565b604051908082528060200260200182016040528015611cfc578160200160208202803683370190505b506020820152866001600160401b03811115611d1a57611d1a614056565b604051908082528060200260200182016040528015611d43578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611d7757611d77614056565b604051908082528060200260200182016040528015611da0578160200160208202803683370190505b5081526020860151516001600160401b03811115611dc057611dc0614056565b604051908082528060200260200182016040528015611de9578160200160208202803683370190505b5081602001819052506000611ebb8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611e92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eb6919061500f565b6138a6565b905060005b87602001515181101561215657611f0588602001518281518110611ee657611ee6614cdf565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611f1b57611f1b614cdf565b60209081029190910101528015611fdb576020830151611f3c60018361502c565b81518110611f4c57611f4c614cdf565b602002602001015160001c83602001518281518110611f6d57611f6d614cdf565b602002602001015160001c11611fdb576040805162461bcd60e51b81526020600482015260248101919091526000805160206153c283398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061202057612020614cdf565b60200260200101518b8b60000151858151811061203f5761203f614cdf565b60200260200101516040518463ffffffff1660e01b815260040161207c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612099573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120bd9190614f2b565b6001600160c01b0316836000015182815181106120dc576120dc614cdf565b602002602001018181525050612142610a63612116848660000151858151811061210857612108614cdf565b602002602001015116613939565b8a60200151848151811061212c5761212c614cdf565b602002602001015161396490919063ffffffff16565b94508061214e81614de6565b915050611ec0565b505061216183613a48565b60975490935060ff166000816121785760006121fa565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121fa9190615043565b905060005b8a81101561287857821561235a578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061225657612256614cdf565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612296573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122ba9190615043565b6122c4919061505c565b1161235a5760405162461bcd60e51b815260206004820152606660248201526000805160206153c283398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061239b5761239b614cdf565b9050013560f81c60f81b60f81c8c8c60a0015185815181106123bf576123bf614cdf565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561241b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061243f9190615074565b6001600160401b0319166124628a604001518381518110611ee657611ee6614cdf565b67ffffffffffffffff1916146124fe5760405162461bcd60e51b815260206004820152606160248201526000805160206153c283398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610695565b61252e8960400151828151811061251757612517614cdf565b60200260200101518761349e90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061257157612571614cdf565b9050013560f81c60f81b60f81c8c8c60c00151858151811061259557612595614cdf565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156125f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126159190614da7565b8560200151828151811061262b5761262b614cdf565b6001600160601b0390921660209283029190910182015285015180518290811061265757612657614cdf565b60200260200101518560000151828151811061267557612675614cdf565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612863576126ed866000015182815181106126bf576126bf614cdf565b60200260200101518f8f868181106126d9576126d9614cdf565b600192013560f81c9290921c811614919050565b15612851577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061273357612733614cdf565b9050013560f81c60f81b60f81c8e8960200151858151811061275757612757614cdf565b60200260200101518f60e00151888151811061277557612775614cdf565b6020026020010151878151811061278e5761278e614cdf565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156127f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128169190614da7565b875180518590811061282a5761282a614cdf565b6020026020010181815161283e919061509f565b6001600160601b03169052506001909101905b8061285b81614de6565b915050612699565b5050808061287090614de6565b9150506121ff565b5050506000806128928c868a606001518b60800151610948565b91509150816129035760405162461bcd60e51b815260206004820152604360248201526000805160206153c283398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610695565b806129645760405162461bcd60e51b815260206004820152603960248201526000805160206153c283398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610695565b5050600087826020015160405160200161297f9291906150c7565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6129b1613ae3565b6129bb60006133b5565b565b60006129cf6060850160408601614229565b90503660006129e1606087018761510f565b909250905060006129f860a0880160808901614229565b905060ca6000612a0b6020890189614229565b63ffffffff1663ffffffff1681526020019081526020016000205487604051602001612a379190615155565b6040516020818303038152906040528051906020012014612ac05760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610695565b600060cb81612ad260208a018a614229565b63ffffffff1663ffffffff1681526020019081526020016000205414612b4f5760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610695565b612b797f000000000000000000000000000000000000000000000000000000000000000085615213565b63ffffffff164363ffffffff161115612bea5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610695565b600086604051602001612bfd9190615289565b604051602081830303815290604052805190602001209050600080612c258387878a8c611a92565b91509150600060405180604001604052804363ffffffff1681526020018381525090508981604051602001612c5b929190615297565b6040516020818303038152906040528051906020012060cb60008c6000016020810190612c889190614229565b63ffffffff1663ffffffff168152602001908152602001600020819055507f58ad2059a3fb85930eecebf2f8b821df273a847746148c59d08769ac458d27b88a82604051612cd7929190615297565b60405180910390a15050505050505050505050565b6040805160a0810182526060808201526001600160a01b0387168152602080820187905263ffffffff43811683850152861660808301528251601f85018290048202810182019093528383529091908490849081908401838280828437600092019190915250505050606082015260cf54604051632770a7eb60e21b8152336004820152602481018790526001600160a01b0390911690639dc29fac90604401600060405180830381600087803b158015612da657600080fd5b505af1158015612dba573d6000803e3d6000fd5b5050505080604051602001612dcf91906152c3565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907fd72b16a52516acce713de75ee6e7b0bebad08b40dc9f695e2471227e41223aba90612e329084906152c3565b60405180910390a260c954612e4e9063ffffffff166001615213565b60c9805463ffffffff191663ffffffff92909216919091179055505050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612eab57612eab614cdf565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612ee79088908690600401614fbb565b600060405180830381865afa158015612f04573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612f2c9190810190614e4b565b600081518110612f3e57612f3e614cdf565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612faa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fce9190614f2b565b6001600160c01b031690506000612fe482613b3d565b905081612ff28a838a610ad2565b9550955050505050935093915050565b61300a613ae3565b6001600160a01b03811661306f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610695565b6106a7816133b5565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130ef9190614c13565b6001600160a01b0316336001600160a01b03161461311f5760405162461bcd60e51b815260040161069590614c30565b60665419811960665419161461319d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107de565b6001600160a01b0381166132625760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610695565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b03161580156132ec57506001600160a01b03821615155b61336e5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26133b1826131d4565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613423613ec0565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561345657613458565bfe5b50806134965760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610695565b505092915050565b60408051808201909152600080825260208201526134ba613ede565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156134565750806134965760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610695565b61353a613efc565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806136226000805160206153a283398151915286614cf5565b90505b61362e81613c09565b90935091506000805160206153a2833981519152828309831415613668576040805180820190915290815260208101919091529392505050565b6000805160206153a2833981519152600182089050613625565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906136b4613f21565b60005b60028110156138795760006136cd826006615360565b90508482600281106136e1576136e1614cdf565b602002015151836136f383600061505c565b600c811061370357613703614cdf565b602002015284826002811061371a5761371a614cdf565b60200201516020015183826001613731919061505c565b600c811061374157613741614cdf565b602002015283826002811061375857613758614cdf565b602002015151518361376b83600261505c565b600c811061377b5761377b614cdf565b602002015283826002811061379257613792614cdf565b60200201515160016020020151836137ab83600361505c565b600c81106137bb576137bb614cdf565b60200201528382600281106137d2576137d2614cdf565b6020020151602001516000600281106137ed576137ed614cdf565b6020020151836137fe83600461505c565b600c811061380e5761380e614cdf565b602002015283826002811061382557613825614cdf565b60200201516020015160016002811061384057613840614cdf565b60200201518361385183600561505c565b600c811061386157613861614cdf565b6020020152508061387181614de6565b9150506136b7565b50613882613f40565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6000806138b284613c8b565b9050808360ff166001901b116139305760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610695565b90505b92915050565b6000805b82156139335761394e60018461502c565b909216918061395c8161537f565b91505061393d565b60408051808201909152600080825260208201526102008261ffff16106139c05760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610695565b8161ffff16600114156139d4575081613933565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613a3d57600161ffff871660ff83161c81161415613a2057613a1d848461349e565b93505b613a2a838461349e565b92506201fffe600192831b1691016139f0565b509195945050505050565b60408051808201909152600080825260208201528151158015613a6d57506020820151155b15613a8b575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206153a28339815191528460200151613abe9190614cf5565b613ad6906000805160206153a283398151915261502c565b905292915050565b919050565b6033546001600160a01b031633146129bb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610695565b6060600080613b4b84613939565b61ffff166001600160401b03811115613b6657613b66614056565b6040519080825280601f01601f191660200182016040528015613b90576020820181803683370190505b5090506000805b825182108015613ba8575061010081105b15613bff576001811b935085841615613bef578060f81b838381518110613bd157613bd1614cdf565b60200101906001600160f81b031916908160001a9053508160010191505b613bf881614de6565b9050613b97565b5090949350505050565b600080806000805160206153a283398151915260036000805160206153a2833981519152866000805160206153a2833981519152888909090890506000613c7f827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206153a2833981519152613e18565b91959194509092505050565b600061010082511115613d145760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610695565b8151613d2257506000919050565b60008083600081518110613d3857613d38614cdf565b0160200151600160f89190911c81901b92505b8451811015613e0f57848181518110613d6657613d66614cdf565b0160200151600160f89190911c1b9150828211613dfb5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610695565b91811791613e0881614de6565b9050613d4b565b50909392505050565b600080613e23613f40565b613e2b613f5e565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613456575082613eb55760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610695565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280613f0f613f7c565b8152602001613f1c613f7c565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106a757600080fd5b600060208284031215613fc157600080fd5b813561393081613f9a565b600060208284031215613fde57600080fd5b5035919050565b600080600080600060a08688031215613ffd57600080fd5b853561400881613f9a565b9450602086013561401881613f9a565b9350604086013561402881613f9a565b9250606086013561403881613f9a565b9150608086013561404881613f9a565b809150509295509295909350565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561408e5761408e614056565b60405290565b60405161010081016001600160401b038111828210171561408e5761408e614056565b604051601f8201601f191681016001600160401b03811182821017156140df576140df614056565b604052919050565b6000604082840312156140f957600080fd5b61410161406c565b9050813581526020820135602082015292915050565b600082601f83011261412857600080fd5b604051604081018181106001600160401b038211171561414a5761414a614056565b806040525080604084018581111561416157600080fd5b845b81811015613a3d578035835260209283019201614163565b60006080828403121561418d57600080fd5b61419561406c565b90506141a18383614117565b81526141b08360408401614117565b602082015292915050565b60008060008061012085870312156141d257600080fd5b843593506141e386602087016140e7565b92506141f2866060870161417b565b91506142018660e087016140e7565b905092959194509250565b63ffffffff811681146106a757600080fd5b8035613ade8161420c565b60006020828403121561423b57600080fd5b81356139308161420c565b60008060006060848603121561425b57600080fd5b833561426681613f9a565b92506020848101356001600160401b038082111561428357600080fd5b818701915087601f83011261429757600080fd5b8135818111156142a9576142a9614056565b6142bb601f8201601f191685016140b7565b915080825288848285010111156142d157600080fd5b80848401858401376000848284010152508094505050506142f46040850161421e565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614393578385038a52825180518087529087019087870190845b8181101561437e57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b0316908401529289019260609092019160010161433a565b50509a87019a9550509185019160010161431c565b509298975050505050505050565b6020815260006143b460208301846142fd565b9392505050565b80151581146106a757600080fd5b6000602082840312156143db57600080fd5b8135613930816143bb565b60008083601f8401126143f857600080fd5b5081356001600160401b0381111561440f57600080fd5b60208301915083602082850101111561442757600080fd5b9250929050565b6000806000806000806080878903121561444757600080fd5b863561445281613f9a565b955060208701356144628161420c565b945060408701356001600160401b038082111561447e57600080fd5b61448a8a838b016143e6565b909650945060608901359150808211156144a357600080fd5b818901915089601f8301126144b757600080fd5b8135818111156144c657600080fd5b8a60208260051b85010111156144db57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561452757815163ffffffff1687529582019590820190600101614505565b509495945050505050565b60006020808352835160808285015261454e60a08501826144f1565b905081850151601f198086840301604087015261456b83836144f1565b9250604087015191508086840301606087015261458883836144f1565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156145df57848783030184526145cd8287516144f1565b958801959388019391506001016145b3565b509998505050505050505050565b60ff811681146106a757600080fd5b60006020828403121561460e57600080fd5b8135613930816145ed565b60006001600160401b0382111561463257614632614056565b5060051b60200190565b60008060006060848603121561465157600080fd5b833561465c81613f9a565b92506020848101356001600160401b0381111561467857600080fd5b8501601f8101871361468957600080fd5b803561469c61469782614619565b6140b7565b81815260059190911b820183019083810190898311156146bb57600080fd5b928401925b828410156146d9578335825292840192908401906146c0565b80965050505050506142f46040850161421e565b6020808252825182820181905260009190848201906040850190845b8181101561472557835183529284019291840191600101614709565b50909695505050505050565b600082601f83011261474257600080fd5b8135602061475261469783614619565b82815260059290921b8401810191818101908684111561477157600080fd5b8286015b848110156147955780356147888161420c565b8352918301918301614775565b509695505050505050565b600082601f8301126147b157600080fd5b813560206147c161469783614619565b82815260069290921b840181019181810190868411156147e057600080fd5b8286015b84811015614795576147f688826140e7565b8352918301916040016147e4565b600082601f83011261481557600080fd5b8135602061482561469783614619565b82815260059290921b8401810191818101908684111561484457600080fd5b8286015b848110156147955780356001600160401b038111156148675760008081fd5b6148758986838b0101614731565b845250918301918301614848565b6000610180828403121561489657600080fd5b61489e614094565b905081356001600160401b03808211156148b757600080fd5b6148c385838601614731565b835260208401359150808211156148d957600080fd5b6148e5858386016147a0565b602084015260408401359150808211156148fe57600080fd5b61490a858386016147a0565b604084015261491c856060860161417b565b606084015261492e8560e086016140e7565b608084015261012084013591508082111561494857600080fd5b61495485838601614731565b60a084015261014084013591508082111561496e57600080fd5b61497a85838601614731565b60c084015261016084013591508082111561499457600080fd5b506149a184828501614804565b60e08301525092915050565b6000806000806000608086880312156149c557600080fd5b8535945060208601356001600160401b03808211156149e357600080fd5b6149ef89838a016143e6565b909650945060408801359150614a048261420c565b90925060608701359080821115614a1a57600080fd5b50614a2788828901614883565b9150509295509295909350565b600081518084526020808501945080840160005b838110156145275781516001600160601b031687529582019590820190600101614a48565b6040815260008351604080840152614a886080840182614a34565b90506020850151603f19848303016060850152614aa58282614a34565b925050508260208301529392505050565b600080600083850360a0811215614acc57600080fd5b84356001600160401b0380821115614ae357600080fd5b9086019060a08289031215614af757600080fd5b8195506060601f1984011215614b0c57600080fd5b6020870194506080870135925080831115614b2657600080fd5b5050614b3486828701614883565b9150509250925092565b600080600080600060808688031215614b5657600080fd5b8535614b6181613f9a565b9450602086013593506040860135614b788161420c565b925060608601356001600160401b03811115614b9357600080fd5b614b9f888289016143e6565b969995985093965092949392505050565b600080600060608486031215614bc557600080fd5b8335614bd081613f9a565b9250602084013591506040840135614be78161420c565b809150509250925092565b828152604060208201526000614c0b60408301846142fd565b949350505050565b600060208284031215614c2557600080fd5b815161393081613f9a565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614c8c57600080fd5b8151613930816143bb565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614d1257634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614d2a57600080fd5b82516001600160401b03811115614d4057600080fd5b8301601f81018513614d5157600080fd5b8051614d5f61469782614619565b81815260059190911b82018301908381019087831115614d7e57600080fd5b928401925b82841015614d9c57835182529284019290840190614d83565b979650505050505050565b600060208284031215614db957600080fd5b81516001600160601b038116811461393057600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415614dfa57614dfa614dd0565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115614e2e57600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215614e5e57600080fd5b82516001600160401b03811115614e7457600080fd5b8301601f81018513614e8557600080fd5b8051614e9361469782614619565b81815260059190911b82018301908381019087831115614eb257600080fd5b928401925b82841015614d9c578351614eca8161420c565b82529284019290840190614eb7565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000614f22604083018486614ed9565b95945050505050565b600060208284031215614f3d57600080fd5b81516001600160c01b038116811461393057600080fd5b600060208284031215614f6657600080fd5b81516139308161420c565b600060ff821660ff811415614f8857614f88614dd0565b60010192915050565b604081526000614fa5604083018587614ed9565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561500257845183529383019391830191600101614fe6565b5090979650505050505050565b60006020828403121561502157600080fd5b8151613930816145ed565b60008282101561503e5761503e614dd0565b500390565b60006020828403121561505557600080fd5b5051919050565b6000821982111561506f5761506f614dd0565b500190565b60006020828403121561508657600080fd5b815167ffffffffffffffff198116811461393057600080fd5b60006001600160601b03838116908316818110156150bf576150bf614dd0565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615102578151855293820193908201906001016150e6565b5092979650505050505050565b6000808335601e1984360301811261512657600080fd5b8301803591506001600160401b0382111561514057600080fd5b60200191503681900382131561442757600080fd5b602081526000823561516681613f9a565b60018060a01b03811660208401525060208301356040830152604083013561518d8161420c565b63ffffffff81166060840152506060830135601e198436030181126151b157600080fd5b830180356001600160401b038111156151c957600080fd5b8036038513156151d857600080fd5b60a060808501526151f060c085018260208501614ed9565b9150506151ff6080850161421e565b63ffffffff811660a0850152509392505050565b600063ffffffff80831681851680830382111561523257615232614dd0565b01949350505050565b80356152468161420c565b63ffffffff168252602081013561525c81613f9a565b6001600160a01b03908116602084015260408201359061527b82613f9a565b808216604085015250505050565b60608101613933828461523b565b60a081016152a5828561523b565b63ffffffff8351166060830152602083015160808301529392505050565b6000602080835260018060a01b038451168184015280840151604084015263ffffffff6040850151166060840152606084015160a0608085015280518060c086015260005b818110156153245782810184015186820160e001528301615308565b8181111561533657600060e083880101525b50608086015163ffffffff811660a08701529250601f01601f19169390930160e001949350505050565b600081600019048311821515161561537a5761537a614dd0565b500290565b600061ffff8083168181141561539757615397614dd0565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212208c8ec14953ae666c64fb40118b67dacfbf24f6fd8794aac270f5b28b4740952c64736f6c634300080c0033",
}

// ContractIncredibleSquaringTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.ABI instead.
var ContractIncredibleSquaringTaskManagerABI = ContractIncredibleSquaringTaskManagerMetaData.ABI

// ContractIncredibleSquaringTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.Bin instead.
var ContractIncredibleSquaringTaskManagerBin = ContractIncredibleSquaringTaskManagerMetaData.Bin

// DeployContractIncredibleSquaringTaskManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSquaringTaskManager to it.
func DeployContractIncredibleSquaringTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractIncredibleSquaringTaskManager, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSquaringTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSquaringTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManager struct {
	ContractIncredibleSquaringTaskManagerCaller     // Read-only binding to the contract
	ContractIncredibleSquaringTaskManagerTransactor // Write-only binding to the contract
	ContractIncredibleSquaringTaskManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSquaringTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSquaringTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSquaringTaskManagerSession struct {
	Contract     *ContractIncredibleSquaringTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSquaringTaskManagerCallerSession struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSquaringTaskManagerTransactorSession struct {
	Contract     *ContractIncredibleSquaringTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerRaw struct {
	Contract *ContractIncredibleSquaringTaskManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCallerRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactorRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSquaringTaskManager creates a new instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSquaringTaskManager, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSquaringTaskManagerCaller creates a new read-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSquaringTaskManagerCaller, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerTransactor creates a new write-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSquaringTaskManagerTransactor, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerFilterer creates a new log filterer instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSquaringTaskManagerFilterer, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSquaringTaskManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSquaringTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// WbtcAddress is a free data retrieval call binding the contract method 0x61a2ac39.
//
// Solidity: function wbtcAddress() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) WbtcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "wbtcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WbtcAddress is a free data retrieval call binding the contract method 0x61a2ac39.
//
// Solidity: function wbtcAddress() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) WbtcAddress() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.WbtcAddress(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// WbtcAddress is a free data retrieval call binding the contract method 0x61a2ac39.
//
// Solidity: function wbtcAddress() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) WbtcAddress() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.WbtcAddress(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xc9689f17.
//
// Solidity: function createNewTask(address destAddress, uint256 amount, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, destAddress common.Address, amount *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "createNewTask", destAddress, amount, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xc9689f17.
//
// Solidity: function createNewTask(address destAddress, uint256 amount, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CreateNewTask(destAddress common.Address, amount *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, destAddress, amount, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xc9689f17.
//
// Solidity: function createNewTask(address destAddress, uint256 amount, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) CreateNewTask(destAddress common.Address, amount *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, destAddress, amount, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _wbtc) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _wbtc common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator, _wbtc)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _wbtc) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _wbtc common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _wbtc)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _wbtc) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _wbtc common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _wbtc)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x766c698b.
//
// Solidity: function respondToTask((address,uint256,uint32,bytes,uint32) task, (uint32,address,address) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x766c698b.
//
// Solidity: function respondToTask((address,uint256,uint32,bytes,uint32) task, (uint32,address,address) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x766c698b.
//
// Solidity: function respondToTask((address,uint256,uint32,bytes,uint32) task, (uint32,address,address) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// ContractIncredibleSquaringTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitializedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerInitialized represents a Initialized event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerInitializedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerInitialized)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSquaringTaskManagerInitialized, error) {
	event := new(ContractIncredibleSquaringTaskManagerInitialized)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IIncredibleSquaringTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0xd72b16a52516acce713de75ee6e7b0bebad08b40dc9f695e2471227e41223aba.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (address,uint256,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0xd72b16a52516acce713de75ee6e7b0bebad08b40dc9f695e2471227e41223aba.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (address,uint256,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0xd72b16a52516acce713de75ee6e7b0bebad08b40dc9f695e2471227e41223aba.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (address,uint256,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractIncredibleSquaringTaskManagerNewTaskCreated, error) {
	event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSquaringTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSquaringTaskManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPaused represents a Paused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePaused(log types.Log) (*ContractIncredibleSquaringTaskManagerPaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerPaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractIncredibleSquaringTaskManagerPauserRegistrySet, error) {
	event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompletedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskCompletedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskCompleted, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskRespondedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskResponded represents a TaskResponded event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskResponded struct {
	TaskResponse         IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x58ad2059a3fb85930eecebf2f8b821df273a847746148c59d08769ac458d27b8.
//
// Solidity: event TaskResponded((uint32,address,address) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskRespondedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x58ad2059a3fb85930eecebf2f8b821df273a847746148c59d08769ac458d27b8.
//
// Solidity: event TaskResponded((uint32,address,address) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0x58ad2059a3fb85930eecebf2f8b821df273a847746148c59d08769ac458d27b8.
//
// Solidity: event TaskResponded((uint32,address,address) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskResponded, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerUnpaused represents a Unpaused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerUnpausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerUnpaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractIncredibleSquaringTaskManagerUnpaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerUnpaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
