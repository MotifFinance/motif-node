// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MotifServiceManager

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

// IRewardsCoordinatorRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// MotifServiceManagerMetaData contains all meta data concerning the MotifServiceManager contract.
var MotifServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_MotifRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardsCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmDeposit\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmWithdrawal\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"transaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardsInitiator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bitcoinPodManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsInitiator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBitcoinPodManager\",\"inputs\":[{\"name\":\"bitcoinPodManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsInitiator\",\"inputs\":[{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyBTCAddress\",\"inputs\":[{\"name\":\"btcAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"script\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyPSBTOutputs\",\"inputs\":[{\"name\":\"psbtBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"withdrawAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"withdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawBitcoinCompleteTx\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"completeTx\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBitcoinPSBT\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"psbtTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BTCAddressVerified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"btcAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BitcoinWithdrawalTransactionSigned\",\"inputs\":[{\"name\":\"pod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsInitiatorUpdated\",\"inputs\":[{\"name\":\"prevRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRewardsInitiator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b5060043610610132575f3560e01c80638da5cb5b116100b4578063c0c53b8b11610079578063c0c53b8b146102bb578063d5eff13d146102ce578063e481af9d146102f1578063f2fde38b146102f9578063fc299dee1461030c578063fce36c7d1461031f575f80fd5b80638da5cb5b1461025e5780639926ee7d1461026f578063a364f4da14610282578063a7d239b114610295578063a98fb355146102a8575f80fd5b806368304835116100fa57806368304835146101ad5780636b3aa72e146101ec578063715018a6146102135780637b6155411461021b578063818001a11461022e575f80fd5b806333cfb7b71461013657806334c9b14e1461015f5780633b1a4e44146101745780633b7d82ef146101875780633bc28c8c1461019a575b5f80fd5b6101496101443660046124e9565b610332565b604051610156919061250b565b60405180910390f35b61017261016d366004612594565b610343565b005b6101726101823660046125e4565b610610565b6101726101953660046125e4565b61095b565b6101726101a83660046124e9565b6109eb565b6101d47f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610156565b6101d47f000000000000000000000000000000000000000000000000000000000000000081565b6101726109ff565b610172610229366004612669565b610a12565b61017261023c3660046124e9565b609780546001600160a01b0319166001600160a01b0392909216919091179055565b6033546001600160a01b03166101d4565b61017261027d3660046127f6565b610d4c565b6101726102903660046124e9565b610da2565b6101726102a336600461289a565b610df3565b6101726102b636600461291e565b61127e565b6101726102c936600461294f565b61128f565b6102e16102dc366004612997565b6113ba565b6040519015158152602001610156565b61014961153d565b6101726103073660046124e9565b61154c565b6065546101d4906001600160a01b031681565b61017261032d366004612a02565b6115c2565b606061033d826115d4565b92915050565b336001600160a01b0316836001600160a01b031663e7f43c686040518163ffffffff1660e01b8152600401602060405180830381865afa158015610389573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ad9190612a70565b6001600160a01b0316146104245760405162461bcd60e51b815260206004820152603360248201527f4f6e6c79206f70657261746f722074686174206f776e732074686520706f642060448201527218d85b8818dbdb999a5c9b4819195c1bdcda5d606a1b60648201526084015b60405180910390fd5b609754604051638597447360e01b81526001600160a01b0385811660048301525f921690638597447390602401606060405180830381865afa15801561046c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104909190612a9f565b6020818101518251604080516bffffffffffffffffffffffff1960608b811b82168388015233901b16603482015260488101939093526068830191909152600160f81b6088830152805180830360690181526089909201905280519101209091505f6104fb826118ed565b90505f61053d8287878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061193f92505050565b90506001600160a01b03811633146105975760405162461bcd60e51b815260206004820152601a60248201527f496e76616c6964204f70657261746f72207369676e6174757265000000000000604482015260640161041b565b6097548451602086015160405163397ee69b60e01b81526001600160a01b038b811660048301526024820193909352604481019190915291169063397ee69b906064015f604051808303815f87803b1580156105f1575f80fd5b505af1158015610603573d5f803e3d5ffd5b5050505050505050505050565b336001600160a01b0316866001600160a01b031663e7f43c686040518163ffffffff1660e01b8152600401602060405180830381865afa158015610656573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061067a9190612a70565b6001600160a01b0316146106a05760405162461bcd60e51b815260040161041b90612add565b60975460405163180fdd4960e31b81526001600160a01b0388811660048301529091169063c07eea48906024015f60405180830381865afa1580156106e7573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261070e9190810190612b69565b51156107665760405162461bcd60e51b815260206004820152602160248201527f5769746864726177616c207265717565737420616c72656164792065786973746044820152607360f81b606482015260840161041b565b60975460405163180fdd4960e31b81526001600160a01b0388811660048301525f92169063c07eea48906024015f60405180830381865afa1580156107ad573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526107d49190810190612b69565b90505f87878787856040516020016107f0959493929190612bc4565b6040516020818303038152906040528051906020012090505f610812826118ed565b90505f6108548287878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061193f92505050565b90506001600160a01b03811633146108ae5760405162461bcd60e51b815260206004820152601a60248201527f496e76616c6964204f70657261746f72207369676e6174757265000000000000604482015260640161041b565b60975460405163d5062c7760e01b81526001600160a01b039091169063d5062c77906108e2908d908c908c90600401612c29565b5f604051808303815f87803b1580156108f9575f80fd5b505af115801561090b573d5f803e3d5ffd5b50506040518b81523392506001600160a01b038d1691507fce93a128bcccad6214b8d927323bca4acc32d14ee0ca93c4e647530c44306fc79060200160405180910390a350505050505050505050565b336001600160a01b0316866001600160a01b031663e7f43c686040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109a1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109c59190612a70565b6001600160a01b0316146107665760405162461bcd60e51b815260040161041b90612add565b6109f3611961565b6109fc816119bb565b50565b610a07611961565b610a105f611a24565b565b336001600160a01b0316856001600160a01b031663e7f43c686040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a58573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a7c9190612a70565b6001600160a01b031614610af15760405162461bcd60e51b815260206004820152603660248201527f4f6e6c79206f70657261746f722074686174206f776e732074686520706f642060448201527518d85b8818dbdb999a5c9b481dda5d1a191c985dd85b60521b606482015260840161041b565b60975460405163180fdd4960e31b81526001600160a01b0387811660048301525f92169063c07eea48906024015f60405180830381865afa158015610b38573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b5f9190810190612b69565b5111610bad5760405162461bcd60e51b815260206004820181905260248201527f4e6f207769746864726177616c207265717565737420746f20636f6e6669726d604482015260640161041b565b60975460405163180fdd4960e31b81526001600160a01b0387811660048301525f92169063c07eea48906024015f60405180830381865afa158015610bf4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610c1b9190810190612b69565b90505f86868684604051602001610c359493929190612c56565b6040516020818303038152906040528051906020012090505f610c57826118ed565b90505f610c998287878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061193f92505050565b90506001600160a01b0381163314610ce75760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b604482015260640161041b565b60975460405163d475ab0360e01b81526001600160a01b038b811660048301529091169063d475ab03906024015f604051808303815f87803b158015610d2b575f80fd5b505af1158015610d3d573d5f803e3d5ffd5b50505050505050505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d945760405162461bcd60e51b815260040161041b90612c8c565b610d9e8282611a75565b5050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610dea5760405162461bcd60e51b815260040161041b90612c8c565b6109fc81611af4565b60405163ec7fbb3160e01b81523360048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ec7fbb3190602401602060405180830381865afa158015610e58573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e7c9190612cfc565b610ec85760405162461bcd60e51b815260206004820152601b60248201527f4f70657261746f72206d75737420626520726567697374657265640000000000604482015260640161041b565b5f8073__",
}

// MotifServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MotifServiceManagerMetaData.ABI instead.
var MotifServiceManagerABI = MotifServiceManagerMetaData.ABI

// MotifServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MotifServiceManagerMetaData.Bin instead.
var MotifServiceManagerBin = MotifServiceManagerMetaData.Bin

// DeployMotifServiceManager deploys a new Ethereum contract, binding an instance of MotifServiceManager to it.
func DeployMotifServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _MotifRegistry common.Address, _rewardsCoordinator common.Address, _delegationManager common.Address) (common.Address, *types.Transaction, *MotifServiceManager, error) {
	parsed, err := MotifServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MotifServiceManagerBin), backend, _avsDirectory, _MotifRegistry, _rewardsCoordinator, _delegationManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MotifServiceManager{MotifServiceManagerCaller: MotifServiceManagerCaller{contract: contract}, MotifServiceManagerTransactor: MotifServiceManagerTransactor{contract: contract}, MotifServiceManagerFilterer: MotifServiceManagerFilterer{contract: contract}}, nil
}

// MotifServiceManager is an auto generated Go binding around an Ethereum contract.
type MotifServiceManager struct {
	MotifServiceManagerCaller     // Read-only binding to the contract
	MotifServiceManagerTransactor // Write-only binding to the contract
	MotifServiceManagerFilterer   // Log filterer for contract events
}

// MotifServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MotifServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MotifServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MotifServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MotifServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MotifServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MotifServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MotifServiceManagerSession struct {
	Contract     *MotifServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MotifServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MotifServiceManagerCallerSession struct {
	Contract *MotifServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MotifServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MotifServiceManagerTransactorSession struct {
	Contract     *MotifServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MotifServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MotifServiceManagerRaw struct {
	Contract *MotifServiceManager // Generic contract binding to access the raw methods on
}

// MotifServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MotifServiceManagerCallerRaw struct {
	Contract *MotifServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MotifServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MotifServiceManagerTransactorRaw struct {
	Contract *MotifServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMotifServiceManager creates a new instance of MotifServiceManager, bound to a specific deployed contract.
func NewMotifServiceManager(address common.Address, backend bind.ContractBackend) (*MotifServiceManager, error) {
	contract, err := bindMotifServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MotifServiceManager{MotifServiceManagerCaller: MotifServiceManagerCaller{contract: contract}, MotifServiceManagerTransactor: MotifServiceManagerTransactor{contract: contract}, MotifServiceManagerFilterer: MotifServiceManagerFilterer{contract: contract}}, nil
}

// NewMotifServiceManagerCaller creates a new read-only instance of MotifServiceManager, bound to a specific deployed contract.
func NewMotifServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*MotifServiceManagerCaller, error) {
	contract, err := bindMotifServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerCaller{contract: contract}, nil
}

// NewMotifServiceManagerTransactor creates a new write-only instance of MotifServiceManager, bound to a specific deployed contract.
func NewMotifServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MotifServiceManagerTransactor, error) {
	contract, err := bindMotifServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerTransactor{contract: contract}, nil
}

// NewMotifServiceManagerFilterer creates a new log filterer instance of MotifServiceManager, bound to a specific deployed contract.
func NewMotifServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MotifServiceManagerFilterer, error) {
	contract, err := bindMotifServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerFilterer{contract: contract}, nil
}

// bindMotifServiceManager binds a generic wrapper to an already deployed contract.
func bindMotifServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MotifServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MotifServiceManager *MotifServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MotifServiceManager.Contract.MotifServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MotifServiceManager *MotifServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.MotifServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MotifServiceManager *MotifServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.MotifServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MotifServiceManager *MotifServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MotifServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MotifServiceManager *MotifServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MotifServiceManager *MotifServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MotifServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_MotifServiceManager *MotifServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _MotifServiceManager.Contract.AvsDirectory(&_MotifServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _MotifServiceManager.Contract.AvsDirectory(&_MotifServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_MotifServiceManager *MotifServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _MotifServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", _operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_MotifServiceManager *MotifServiceManagerSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _MotifServiceManager.Contract.GetOperatorRestakedStrategies(&_MotifServiceManager.CallOpts, _operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_MotifServiceManager *MotifServiceManagerCallerSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _MotifServiceManager.Contract.GetOperatorRestakedStrategies(&_MotifServiceManager.CallOpts, _operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_MotifServiceManager *MotifServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MotifServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_MotifServiceManager *MotifServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _MotifServiceManager.Contract.GetRestakeableStrategies(&_MotifServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_MotifServiceManager *MotifServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _MotifServiceManager.Contract.GetRestakeableStrategies(&_MotifServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MotifServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MotifServiceManager *MotifServiceManagerSession) Owner() (common.Address, error) {
	return _MotifServiceManager.Contract.Owner(&_MotifServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCallerSession) Owner() (common.Address, error) {
	return _MotifServiceManager.Contract.Owner(&_MotifServiceManager.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCaller) RewardsInitiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MotifServiceManager.contract.Call(opts, &out, "rewardsInitiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_MotifServiceManager *MotifServiceManagerSession) RewardsInitiator() (common.Address, error) {
	return _MotifServiceManager.Contract.RewardsInitiator(&_MotifServiceManager.CallOpts)
}

// RewardsInitiator is a free data retrieval call binding the contract method 0xfc299dee.
//
// Solidity: function rewardsInitiator() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCallerSession) RewardsInitiator() (common.Address, error) {
	return _MotifServiceManager.Contract.RewardsInitiator(&_MotifServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MotifServiceManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_MotifServiceManager *MotifServiceManagerSession) StakeRegistry() (common.Address, error) {
	return _MotifServiceManager.Contract.StakeRegistry(&_MotifServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_MotifServiceManager *MotifServiceManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _MotifServiceManager.Contract.StakeRegistry(&_MotifServiceManager.CallOpts)
}

// VerifyPSBTOutputs is a free data retrieval call binding the contract method 0xd5eff13d.
//
// Solidity: function verifyPSBTOutputs(bytes psbtBytes, string withdrawAddress, uint256 withdrawAmount) pure returns(bool)
func (_MotifServiceManager *MotifServiceManagerCaller) VerifyPSBTOutputs(opts *bind.CallOpts, psbtBytes []byte, withdrawAddress string, withdrawAmount *big.Int) (bool, error) {
	var out []interface{}
	err := _MotifServiceManager.contract.Call(opts, &out, "verifyPSBTOutputs", psbtBytes, withdrawAddress, withdrawAmount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPSBTOutputs is a free data retrieval call binding the contract method 0xd5eff13d.
//
// Solidity: function verifyPSBTOutputs(bytes psbtBytes, string withdrawAddress, uint256 withdrawAmount) pure returns(bool)
func (_MotifServiceManager *MotifServiceManagerSession) VerifyPSBTOutputs(psbtBytes []byte, withdrawAddress string, withdrawAmount *big.Int) (bool, error) {
	return _MotifServiceManager.Contract.VerifyPSBTOutputs(&_MotifServiceManager.CallOpts, psbtBytes, withdrawAddress, withdrawAmount)
}

// VerifyPSBTOutputs is a free data retrieval call binding the contract method 0xd5eff13d.
//
// Solidity: function verifyPSBTOutputs(bytes psbtBytes, string withdrawAddress, uint256 withdrawAmount) pure returns(bool)
func (_MotifServiceManager *MotifServiceManagerCallerSession) VerifyPSBTOutputs(psbtBytes []byte, withdrawAddress string, withdrawAmount *big.Int) (bool, error) {
	return _MotifServiceManager.Contract.VerifyPSBTOutputs(&_MotifServiceManager.CallOpts, psbtBytes, withdrawAddress, withdrawAmount)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x34c9b14e.
//
// Solidity: function confirmDeposit(address pod, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) ConfirmDeposit(opts *bind.TransactOpts, pod common.Address, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "confirmDeposit", pod, signature)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x34c9b14e.
//
// Solidity: function confirmDeposit(address pod, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerSession) ConfirmDeposit(pod common.Address, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.ConfirmDeposit(&_MotifServiceManager.TransactOpts, pod, signature)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x34c9b14e.
//
// Solidity: function confirmDeposit(address pod, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) ConfirmDeposit(pod common.Address, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.ConfirmDeposit(&_MotifServiceManager.TransactOpts, pod, signature)
}

// ConfirmWithdrawal is a paid mutator transaction binding the contract method 0x7b615541.
//
// Solidity: function confirmWithdrawal(address pod, bytes transaction, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) ConfirmWithdrawal(opts *bind.TransactOpts, pod common.Address, transaction []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "confirmWithdrawal", pod, transaction, signature)
}

// ConfirmWithdrawal is a paid mutator transaction binding the contract method 0x7b615541.
//
// Solidity: function confirmWithdrawal(address pod, bytes transaction, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerSession) ConfirmWithdrawal(pod common.Address, transaction []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.ConfirmWithdrawal(&_MotifServiceManager.TransactOpts, pod, transaction, signature)
}

// ConfirmWithdrawal is a paid mutator transaction binding the contract method 0x7b615541.
//
// Solidity: function confirmWithdrawal(address pod, bytes transaction, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) ConfirmWithdrawal(pod common.Address, transaction []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.ConfirmWithdrawal(&_MotifServiceManager.TransactOpts, pod, transaction, signature)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_MotifServiceManager *MotifServiceManagerSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.CreateAVSRewardsSubmission(&_MotifServiceManager.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.CreateAVSRewardsSubmission(&_MotifServiceManager.TransactOpts, rewardsSubmissions)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_MotifServiceManager *MotifServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.DeregisterOperatorFromAVS(&_MotifServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.DeregisterOperatorFromAVS(&_MotifServiceManager.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _rewardsInitiator, address bitcoinPodManager) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _rewardsInitiator common.Address, bitcoinPodManager common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "initialize", _owner, _rewardsInitiator, bitcoinPodManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _rewardsInitiator, address bitcoinPodManager) returns()
func (_MotifServiceManager *MotifServiceManagerSession) Initialize(_owner common.Address, _rewardsInitiator common.Address, bitcoinPodManager common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.Initialize(&_MotifServiceManager.TransactOpts, _owner, _rewardsInitiator, bitcoinPodManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _owner, address _rewardsInitiator, address bitcoinPodManager) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) Initialize(_owner common.Address, _rewardsInitiator common.Address, bitcoinPodManager common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.Initialize(&_MotifServiceManager.TransactOpts, _owner, _rewardsInitiator, bitcoinPodManager)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_MotifServiceManager *MotifServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.RegisterOperatorToAVS(&_MotifServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.RegisterOperatorToAVS(&_MotifServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MotifServiceManager *MotifServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MotifServiceManager.Contract.RenounceOwnership(&_MotifServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MotifServiceManager.Contract.RenounceOwnership(&_MotifServiceManager.TransactOpts)
}

// SetBitcoinPodManager is a paid mutator transaction binding the contract method 0x818001a1.
//
// Solidity: function setBitcoinPodManager(address bitcoinPodManager) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) SetBitcoinPodManager(opts *bind.TransactOpts, bitcoinPodManager common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "setBitcoinPodManager", bitcoinPodManager)
}

// SetBitcoinPodManager is a paid mutator transaction binding the contract method 0x818001a1.
//
// Solidity: function setBitcoinPodManager(address bitcoinPodManager) returns()
func (_MotifServiceManager *MotifServiceManagerSession) SetBitcoinPodManager(bitcoinPodManager common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.SetBitcoinPodManager(&_MotifServiceManager.TransactOpts, bitcoinPodManager)
}

// SetBitcoinPodManager is a paid mutator transaction binding the contract method 0x818001a1.
//
// Solidity: function setBitcoinPodManager(address bitcoinPodManager) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) SetBitcoinPodManager(bitcoinPodManager common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.SetBitcoinPodManager(&_MotifServiceManager.TransactOpts, bitcoinPodManager)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) SetRewardsInitiator(opts *bind.TransactOpts, newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "setRewardsInitiator", newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_MotifServiceManager *MotifServiceManagerSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.SetRewardsInitiator(&_MotifServiceManager.TransactOpts, newRewardsInitiator)
}

// SetRewardsInitiator is a paid mutator transaction binding the contract method 0x3bc28c8c.
//
// Solidity: function setRewardsInitiator(address newRewardsInitiator) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) SetRewardsInitiator(newRewardsInitiator common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.SetRewardsInitiator(&_MotifServiceManager.TransactOpts, newRewardsInitiator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MotifServiceManager *MotifServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.TransferOwnership(&_MotifServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.TransferOwnership(&_MotifServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_MotifServiceManager *MotifServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.UpdateAVSMetadataURI(&_MotifServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.UpdateAVSMetadataURI(&_MotifServiceManager.TransactOpts, _metadataURI)
}

// VerifyBTCAddress is a paid mutator transaction binding the contract method 0xa7d239b1.
//
// Solidity: function verifyBTCAddress(string btcAddress, bytes script) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) VerifyBTCAddress(opts *bind.TransactOpts, btcAddress string, script []byte) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "verifyBTCAddress", btcAddress, script)
}

// VerifyBTCAddress is a paid mutator transaction binding the contract method 0xa7d239b1.
//
// Solidity: function verifyBTCAddress(string btcAddress, bytes script) returns()
func (_MotifServiceManager *MotifServiceManagerSession) VerifyBTCAddress(btcAddress string, script []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.VerifyBTCAddress(&_MotifServiceManager.TransactOpts, btcAddress, script)
}

// VerifyBTCAddress is a paid mutator transaction binding the contract method 0xa7d239b1.
//
// Solidity: function verifyBTCAddress(string btcAddress, bytes script) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) VerifyBTCAddress(btcAddress string, script []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.VerifyBTCAddress(&_MotifServiceManager.TransactOpts, btcAddress, script)
}

// WithdrawBitcoinCompleteTx is a paid mutator transaction binding the contract method 0x3b7d82ef.
//
// Solidity: function withdrawBitcoinCompleteTx(address pod, uint256 amount, bytes completeTx, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) WithdrawBitcoinCompleteTx(opts *bind.TransactOpts, pod common.Address, amount *big.Int, completeTx []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "withdrawBitcoinCompleteTx", pod, amount, completeTx, signature)
}

// WithdrawBitcoinCompleteTx is a paid mutator transaction binding the contract method 0x3b7d82ef.
//
// Solidity: function withdrawBitcoinCompleteTx(address pod, uint256 amount, bytes completeTx, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerSession) WithdrawBitcoinCompleteTx(pod common.Address, amount *big.Int, completeTx []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.WithdrawBitcoinCompleteTx(&_MotifServiceManager.TransactOpts, pod, amount, completeTx, signature)
}

// WithdrawBitcoinCompleteTx is a paid mutator transaction binding the contract method 0x3b7d82ef.
//
// Solidity: function withdrawBitcoinCompleteTx(address pod, uint256 amount, bytes completeTx, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) WithdrawBitcoinCompleteTx(pod common.Address, amount *big.Int, completeTx []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.WithdrawBitcoinCompleteTx(&_MotifServiceManager.TransactOpts, pod, amount, completeTx, signature)
}

// WithdrawBitcoinPSBT is a paid mutator transaction binding the contract method 0x3b1a4e44.
//
// Solidity: function withdrawBitcoinPSBT(address pod, uint256 amount, bytes psbtTransaction, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactor) WithdrawBitcoinPSBT(opts *bind.TransactOpts, pod common.Address, amount *big.Int, psbtTransaction []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.contract.Transact(opts, "withdrawBitcoinPSBT", pod, amount, psbtTransaction, signature)
}

// WithdrawBitcoinPSBT is a paid mutator transaction binding the contract method 0x3b1a4e44.
//
// Solidity: function withdrawBitcoinPSBT(address pod, uint256 amount, bytes psbtTransaction, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerSession) WithdrawBitcoinPSBT(pod common.Address, amount *big.Int, psbtTransaction []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.WithdrawBitcoinPSBT(&_MotifServiceManager.TransactOpts, pod, amount, psbtTransaction, signature)
}

// WithdrawBitcoinPSBT is a paid mutator transaction binding the contract method 0x3b1a4e44.
//
// Solidity: function withdrawBitcoinPSBT(address pod, uint256 amount, bytes psbtTransaction, bytes signature) returns()
func (_MotifServiceManager *MotifServiceManagerTransactorSession) WithdrawBitcoinPSBT(pod common.Address, amount *big.Int, psbtTransaction []byte, signature []byte) (*types.Transaction, error) {
	return _MotifServiceManager.Contract.WithdrawBitcoinPSBT(&_MotifServiceManager.TransactOpts, pod, amount, psbtTransaction, signature)
}

// MotifServiceManagerBTCAddressVerifiedIterator is returned from FilterBTCAddressVerified and is used to iterate over the raw logs and unpacked data for BTCAddressVerified events raised by the MotifServiceManager contract.
type MotifServiceManagerBTCAddressVerifiedIterator struct {
	Event *MotifServiceManagerBTCAddressVerified // Event containing the contract specifics and raw log

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
func (it *MotifServiceManagerBTCAddressVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MotifServiceManagerBTCAddressVerified)
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
		it.Event = new(MotifServiceManagerBTCAddressVerified)
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
func (it *MotifServiceManagerBTCAddressVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MotifServiceManagerBTCAddressVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MotifServiceManagerBTCAddressVerified represents a BTCAddressVerified event raised by the MotifServiceManager contract.
type MotifServiceManagerBTCAddressVerified struct {
	Operator   common.Address
	BtcAddress string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBTCAddressVerified is a free log retrieval operation binding the contract event 0xabc3bc0a50a83331466b147d3ce6d2a71cc52bfa15d52467b4860bb8d08c206b.
//
// Solidity: event BTCAddressVerified(address indexed operator, string btcAddress)
func (_MotifServiceManager *MotifServiceManagerFilterer) FilterBTCAddressVerified(opts *bind.FilterOpts, operator []common.Address) (*MotifServiceManagerBTCAddressVerifiedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MotifServiceManager.contract.FilterLogs(opts, "BTCAddressVerified", operatorRule)
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerBTCAddressVerifiedIterator{contract: _MotifServiceManager.contract, event: "BTCAddressVerified", logs: logs, sub: sub}, nil
}

// WatchBTCAddressVerified is a free log subscription operation binding the contract event 0xabc3bc0a50a83331466b147d3ce6d2a71cc52bfa15d52467b4860bb8d08c206b.
//
// Solidity: event BTCAddressVerified(address indexed operator, string btcAddress)
func (_MotifServiceManager *MotifServiceManagerFilterer) WatchBTCAddressVerified(opts *bind.WatchOpts, sink chan<- *MotifServiceManagerBTCAddressVerified, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MotifServiceManager.contract.WatchLogs(opts, "BTCAddressVerified", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MotifServiceManagerBTCAddressVerified)
				if err := _MotifServiceManager.contract.UnpackLog(event, "BTCAddressVerified", log); err != nil {
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

// ParseBTCAddressVerified is a log parse operation binding the contract event 0xabc3bc0a50a83331466b147d3ce6d2a71cc52bfa15d52467b4860bb8d08c206b.
//
// Solidity: event BTCAddressVerified(address indexed operator, string btcAddress)
func (_MotifServiceManager *MotifServiceManagerFilterer) ParseBTCAddressVerified(log types.Log) (*MotifServiceManagerBTCAddressVerified, error) {
	event := new(MotifServiceManagerBTCAddressVerified)
	if err := _MotifServiceManager.contract.UnpackLog(event, "BTCAddressVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MotifServiceManagerBitcoinWithdrawalTransactionSignedIterator is returned from FilterBitcoinWithdrawalTransactionSigned and is used to iterate over the raw logs and unpacked data for BitcoinWithdrawalTransactionSigned events raised by the MotifServiceManager contract.
type MotifServiceManagerBitcoinWithdrawalTransactionSignedIterator struct {
	Event *MotifServiceManagerBitcoinWithdrawalTransactionSigned // Event containing the contract specifics and raw log

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
func (it *MotifServiceManagerBitcoinWithdrawalTransactionSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MotifServiceManagerBitcoinWithdrawalTransactionSigned)
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
		it.Event = new(MotifServiceManagerBitcoinWithdrawalTransactionSigned)
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
func (it *MotifServiceManagerBitcoinWithdrawalTransactionSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MotifServiceManagerBitcoinWithdrawalTransactionSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MotifServiceManagerBitcoinWithdrawalTransactionSigned represents a BitcoinWithdrawalTransactionSigned event raised by the MotifServiceManager contract.
type MotifServiceManagerBitcoinWithdrawalTransactionSigned struct {
	Pod      common.Address
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBitcoinWithdrawalTransactionSigned is a free log retrieval operation binding the contract event 0xce93a128bcccad6214b8d927323bca4acc32d14ee0ca93c4e647530c44306fc7.
//
// Solidity: event BitcoinWithdrawalTransactionSigned(address indexed pod, address indexed operator, uint256 amount)
func (_MotifServiceManager *MotifServiceManagerFilterer) FilterBitcoinWithdrawalTransactionSigned(opts *bind.FilterOpts, pod []common.Address, operator []common.Address) (*MotifServiceManagerBitcoinWithdrawalTransactionSignedIterator, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MotifServiceManager.contract.FilterLogs(opts, "BitcoinWithdrawalTransactionSigned", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerBitcoinWithdrawalTransactionSignedIterator{contract: _MotifServiceManager.contract, event: "BitcoinWithdrawalTransactionSigned", logs: logs, sub: sub}, nil
}

// WatchBitcoinWithdrawalTransactionSigned is a free log subscription operation binding the contract event 0xce93a128bcccad6214b8d927323bca4acc32d14ee0ca93c4e647530c44306fc7.
//
// Solidity: event BitcoinWithdrawalTransactionSigned(address indexed pod, address indexed operator, uint256 amount)
func (_MotifServiceManager *MotifServiceManagerFilterer) WatchBitcoinWithdrawalTransactionSigned(opts *bind.WatchOpts, sink chan<- *MotifServiceManagerBitcoinWithdrawalTransactionSigned, pod []common.Address, operator []common.Address) (event.Subscription, error) {

	var podRule []interface{}
	for _, podItem := range pod {
		podRule = append(podRule, podItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MotifServiceManager.contract.WatchLogs(opts, "BitcoinWithdrawalTransactionSigned", podRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MotifServiceManagerBitcoinWithdrawalTransactionSigned)
				if err := _MotifServiceManager.contract.UnpackLog(event, "BitcoinWithdrawalTransactionSigned", log); err != nil {
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

// ParseBitcoinWithdrawalTransactionSigned is a log parse operation binding the contract event 0xce93a128bcccad6214b8d927323bca4acc32d14ee0ca93c4e647530c44306fc7.
//
// Solidity: event BitcoinWithdrawalTransactionSigned(address indexed pod, address indexed operator, uint256 amount)
func (_MotifServiceManager *MotifServiceManagerFilterer) ParseBitcoinWithdrawalTransactionSigned(log types.Log) (*MotifServiceManagerBitcoinWithdrawalTransactionSigned, error) {
	event := new(MotifServiceManagerBitcoinWithdrawalTransactionSigned)
	if err := _MotifServiceManager.contract.UnpackLog(event, "BitcoinWithdrawalTransactionSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MotifServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MotifServiceManager contract.
type MotifServiceManagerInitializedIterator struct {
	Event *MotifServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *MotifServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MotifServiceManagerInitialized)
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
		it.Event = new(MotifServiceManagerInitialized)
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
func (it *MotifServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MotifServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MotifServiceManagerInitialized represents a Initialized event raised by the MotifServiceManager contract.
type MotifServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MotifServiceManager *MotifServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MotifServiceManagerInitializedIterator, error) {

	logs, sub, err := _MotifServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerInitializedIterator{contract: _MotifServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MotifServiceManager *MotifServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MotifServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _MotifServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MotifServiceManagerInitialized)
				if err := _MotifServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MotifServiceManager *MotifServiceManagerFilterer) ParseInitialized(log types.Log) (*MotifServiceManagerInitialized, error) {
	event := new(MotifServiceManagerInitialized)
	if err := _MotifServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MotifServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MotifServiceManager contract.
type MotifServiceManagerOwnershipTransferredIterator struct {
	Event *MotifServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MotifServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MotifServiceManagerOwnershipTransferred)
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
		it.Event = new(MotifServiceManagerOwnershipTransferred)
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
func (it *MotifServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MotifServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MotifServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MotifServiceManager contract.
type MotifServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MotifServiceManager *MotifServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MotifServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MotifServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerOwnershipTransferredIterator{contract: _MotifServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MotifServiceManager *MotifServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MotifServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MotifServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MotifServiceManagerOwnershipTransferred)
				if err := _MotifServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MotifServiceManager *MotifServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MotifServiceManagerOwnershipTransferred, error) {
	event := new(MotifServiceManagerOwnershipTransferred)
	if err := _MotifServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MotifServiceManagerRewardsInitiatorUpdatedIterator is returned from FilterRewardsInitiatorUpdated and is used to iterate over the raw logs and unpacked data for RewardsInitiatorUpdated events raised by the MotifServiceManager contract.
type MotifServiceManagerRewardsInitiatorUpdatedIterator struct {
	Event *MotifServiceManagerRewardsInitiatorUpdated // Event containing the contract specifics and raw log

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
func (it *MotifServiceManagerRewardsInitiatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MotifServiceManagerRewardsInitiatorUpdated)
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
		it.Event = new(MotifServiceManagerRewardsInitiatorUpdated)
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
func (it *MotifServiceManagerRewardsInitiatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MotifServiceManagerRewardsInitiatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MotifServiceManagerRewardsInitiatorUpdated represents a RewardsInitiatorUpdated event raised by the MotifServiceManager contract.
type MotifServiceManagerRewardsInitiatorUpdated struct {
	PrevRewardsInitiator common.Address
	NewRewardsInitiator  common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRewardsInitiatorUpdated is a free log retrieval operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_MotifServiceManager *MotifServiceManagerFilterer) FilterRewardsInitiatorUpdated(opts *bind.FilterOpts) (*MotifServiceManagerRewardsInitiatorUpdatedIterator, error) {

	logs, sub, err := _MotifServiceManager.contract.FilterLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return &MotifServiceManagerRewardsInitiatorUpdatedIterator{contract: _MotifServiceManager.contract, event: "RewardsInitiatorUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsInitiatorUpdated is a free log subscription operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_MotifServiceManager *MotifServiceManagerFilterer) WatchRewardsInitiatorUpdated(opts *bind.WatchOpts, sink chan<- *MotifServiceManagerRewardsInitiatorUpdated) (event.Subscription, error) {

	logs, sub, err := _MotifServiceManager.contract.WatchLogs(opts, "RewardsInitiatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MotifServiceManagerRewardsInitiatorUpdated)
				if err := _MotifServiceManager.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
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

// ParseRewardsInitiatorUpdated is a log parse operation binding the contract event 0xe11cddf1816a43318ca175bbc52cd0185436e9cbead7c83acc54a73e461717e3.
//
// Solidity: event RewardsInitiatorUpdated(address prevRewardsInitiator, address newRewardsInitiator)
func (_MotifServiceManager *MotifServiceManagerFilterer) ParseRewardsInitiatorUpdated(log types.Log) (*MotifServiceManagerRewardsInitiatorUpdated, error) {
	event := new(MotifServiceManagerRewardsInitiatorUpdated)
	if err := _MotifServiceManager.contract.UnpackLog(event, "RewardsInitiatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
