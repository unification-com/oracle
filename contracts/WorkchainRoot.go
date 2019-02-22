// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"math/big"
	"strings"

	ethereum "github.com/unification-com/mainchain"
	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getGenesis\",\"outputs\":[{\"name\":\"genesis_hash_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"name\":\"chain_id_\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chain_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_height\",\"type\":\"uint64\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_parent_hash\",\"type\":\"bytes32\"},{\"name\":\"_receipt_root\",\"type\":\"bytes32\"},{\"name\":\"_tx_root\",\"type\":\"bytes32\"},{\"name\":\"_state_root\",\"type\":\"bytes32\"},{\"name\":\"_sealer\",\"type\":\"address\"},{\"name\":\"_chain_id\",\"type\":\"uint64\"}],\"name\":\"recordHeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ev\",\"type\":\"address\"}],\"name\":\"isEv\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"block_headers\",\"outputs\":[{\"name\":\"height\",\"type\":\"uint64\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"parent_hash\",\"type\":\"bytes32\"},{\"name\":\"receipt_root\",\"type\":\"bytes32\"},{\"name\":\"tx_root\",\"type\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\"},{\"name\":\"sealer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"getHeader\",\"outputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"parent_hash\",\"type\":\"bytes32\"},{\"name\":\"receipt_root\",\"type\":\"bytes32\"},{\"name\":\"tx_root\",\"type\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\"},{\"name\":\"sealer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesis_hash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEvs\",\"outputs\":[{\"name\":\"current_evs_idx_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_new_evs\",\"type\":\"address[]\"}],\"name\":\"setEvs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_chain_id\",\"type\":\"uint64\"},{\"name\":\"_genesis_hash\",\"type\":\"bytes32\"},{\"name\":\"_evs\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"height\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"parent_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"receipt_root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"tx_root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"state_root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sealer\",\"type\":\"address\"}],\"name\":\"RecordHeader\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `60806040523480156200001157600080fd5b50604051620015df380380620015df833981018060405260608110156200003757600080fd5b81019080805190602001909291908051906020019092919080516401000000008111156200006457600080fd5b828101905060208101848111156200007b57600080fd5b81518560208202830111640100000000821117156200009957600080fd5b505092919050505060008367ffffffffffffffff1611151562000124576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f436861696e49442072657175697265640000000000000000000000000000000081525060200191505060405180910390fd5b6000602060ff16111515620001a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f47656e6573697320626c6f636b2072657175697265640000000000000000000081525060200191505060405180910390fd5b600081511115156200021b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f496e697469616c2045567320726571756972656400000000000000000000000081525060200191505060405180910390fd5b826000806101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508160018190555060008090505b8151811015620002d85760016003600084848151811015156200026f57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550808060010191505062000250565b508060049080519060200190620002f1929190620002fb565b50505050620003d0565b82805482825590600052602060002090810192821562000377579160200282015b82811115620003765782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906200031c565b5b5090506200038691906200038a565b5090565b620003cd91905b80821115620003c957600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010162000391565b5090565b90565b6111ff80620003e06000396000f3fe608060405234801561001057600080fd5b50600436106100bb576000357c010000000000000000000000000000000000000000000000000000000090048063a794c9fd11610083578063a794c9fd1461023d578063b203e111146102f3578063c366ea9e1461038e578063c5cb0a98146103ac578063db8297981461040b576100bb565b80631a43bcb5146100c05780633408e470146100de5780633af973b1146101105780633f10cc1e146101425780639e186c93146101e1575b600080fd5b6100c86104c3565b6040518082815260200191505060405180910390f35b6100e66104cd565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b6101186104ea565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b6101df600480360361010081101561015957600080fd5b81019080803567ffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803567ffffffffffffffff169060200190929190505050610503565b005b610223600480360360208110156101f757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c82565b604051808215151515815260200191505060405180910390f35b6102736004803603602081101561025357600080fd5b81019080803567ffffffffffffffff169060200190929190505050610cf0565b604051808867ffffffffffffffff1667ffffffffffffffff1681526020018781526020018681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200197505050505050505060405180910390f35b6103296004803603602081101561030957600080fd5b81019080803567ffffffffffffffff169060200190929190505050610d66565b604051808781526020018681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001965050505050505060405180910390f35b610396610dee565b6040518082815260200191505060405180910390f35b6103b4610df4565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156103f75780820151818401526020810190506103dc565b505050509050019250505060405180910390f35b6104c16004803603602081101561042157600080fd5b810190808035906020019064010000000081111561043e57600080fd5b82018360208201111561045057600080fd5b8035906020019184602083028401116401000000008311171561047257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610e82565b005b6000600154905090565b60008060009054906101000a900467ffffffffffffffff16905090565b6000809054906101000a900467ffffffffffffffff1681565b60011515600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151561056257600080fd5b6000602060ff161115156105de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f486173682072657175697265640000000000000000000000000000000000000081525060200191505060405180910390fd5b6000602060ff1611151561065a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f506172656e74204861736820726571756972656400000000000000000000000081525060200191505060405180910390fd5b6000602060ff161115156106d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f5265636569707420726f6f74207265717569726564000000000000000000000081525060200191505060405180910390fd5b6000602060ff16111515610752576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f547820726f6f742072657175697265640000000000000000000000000000000081525060200191505060405180910390fd5b6000602060ff161115156107ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f537461746520726f6f742072657175697265640000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610873576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f5365616c6572206164647265737320726571756972656400000000000000000081525060200191505060405180910390fd5b60011515600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151561093b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f5365616c6572206973206e6f7420612063757272656e7420455600000000000081525060200191505060405180910390fd5b6000809054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161415156109db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f436861696e20494420646f6573206e6f74206d6174636800000000000000000081525060200191505060405180910390fd5b6000600260008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002090508867ffffffffffffffff168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614151515610aaa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f426c6f636b2068656164657220616c726561647920657869737473000000000081525060200191505060405180910390fd5b60e0604051908101604052808a67ffffffffffffffff1681526020018981526020018881526020018781526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff16815250600260008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507fceec4c706f6e91cdd751ea972853f6862c309ad04ec22a41ca891868e495dd2c89898989898989604051808867ffffffffffffffff1667ffffffffffffffff1681526020018781526020018681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200197505050505050505060405180910390a1505050505050505050565b600060011515600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151415610ce65760019050610ceb565b600090505b919050565b60026020528060005260406000206000915090508060000160009054906101000a900467ffffffffffffffff16908060010154908060020154908060030154908060040154908060050154908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b6000806000806000806000600260008967ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020905080600101549650806002015495508060030154945080600401549350806005015492508060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505091939550919395565b60015481565b60606004805480602002602001604051908101604052809291908181526020018280548015610e7857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e2e575b5050505050905090565b60011515600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515610ee157600080fd5b60008151111515610f5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f455673207265717569726564000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008090505b60048054905081101561100b57600060036000600484815481101515610f8257fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050610f60565b5060008090505b815181101561109757600160036000848481518110151561102f57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080600101915050611012565b50600460006110a691906110c0565b80600490805190602001906110bc9291906110e1565b5050565b50805460008255906000526020600020908101906110de919061116b565b50565b82805482825590600052602060002090810192821561115a579160200282015b828111156111595782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611101565b5b5090506111679190611190565b5090565b61118d91905b80821115611189576000816000905550600101611171565b5090565b90565b6111d091905b808211156111cc57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101611196565b5090565b9056fea165627a7a723058207a468bf750311934769ef0f84845a163d98009dd2c09ba8a71539ab6228859c90029`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _chain_id uint64, _genesis_hash [32]byte, _evs []common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _chain_id, _genesis_hash, _evs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// BlockHeaders is a free data retrieval call binding the contract method 0xa794c9fd.
//
// Solidity: function block_headers(uint64 ) constant returns(uint64 height, bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreCaller) BlockHeaders(opts *bind.CallOpts, arg0 uint64) (struct {
	Height      uint64
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Sealer      common.Address
}, error) {
	ret := new(struct {
		Height      uint64
		Hash        [32]byte
		ParentHash  [32]byte
		ReceiptRoot [32]byte
		TxRoot      [32]byte
		StateRoot   [32]byte
		Sealer      common.Address
	})
	out := ret
	err := _Store.contract.Call(opts, out, "block_headers", arg0)
	return *ret, err
}

// BlockHeaders is a free data retrieval call binding the contract method 0xa794c9fd.
//
// Solidity: function block_headers(uint64 ) constant returns(uint64 height, bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreSession) BlockHeaders(arg0 uint64) (struct {
	Height      uint64
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Sealer      common.Address
}, error) {
	return _Store.Contract.BlockHeaders(&_Store.CallOpts, arg0)
}

// BlockHeaders is a free data retrieval call binding the contract method 0xa794c9fd.
//
// Solidity: function block_headers(uint64 ) constant returns(uint64 height, bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreCallerSession) BlockHeaders(arg0 uint64) (struct {
	Height      uint64
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Sealer      common.Address
}, error) {
	return _Store.Contract.BlockHeaders(&_Store.CallOpts, arg0)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint64)
func (_Store *StoreCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "chain_id")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint64)
func (_Store *StoreSession) ChainId() (uint64, error) {
	return _Store.Contract.ChainId(&_Store.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint64)
func (_Store *StoreCallerSession) ChainId() (uint64, error) {
	return _Store.Contract.ChainId(&_Store.CallOpts)
}

// GenesisHash is a free data retrieval call binding the contract method 0xc366ea9e.
//
// Solidity: function genesis_hash() constant returns(bytes32)
func (_Store *StoreCaller) GenesisHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "genesis_hash")
	return *ret0, err
}

// GenesisHash is a free data retrieval call binding the contract method 0xc366ea9e.
//
// Solidity: function genesis_hash() constant returns(bytes32)
func (_Store *StoreSession) GenesisHash() ([32]byte, error) {
	return _Store.Contract.GenesisHash(&_Store.CallOpts)
}

// GenesisHash is a free data retrieval call binding the contract method 0xc366ea9e.
//
// Solidity: function genesis_hash() constant returns(bytes32)
func (_Store *StoreCallerSession) GenesisHash() ([32]byte, error) {
	return _Store.Contract.GenesisHash(&_Store.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() constant returns(uint64 chain_id_)
func (_Store *StoreCaller) GetChainId(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getChainId")
	return *ret0, err
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() constant returns(uint64 chain_id_)
func (_Store *StoreSession) GetChainId() (uint64, error) {
	return _Store.Contract.GetChainId(&_Store.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() constant returns(uint64 chain_id_)
func (_Store *StoreCallerSession) GetChainId() (uint64, error) {
	return _Store.Contract.GetChainId(&_Store.CallOpts)
}

// GetEvs is a free data retrieval call binding the contract method 0xc5cb0a98.
//
// Solidity: function getEvs() constant returns(address[] current_evs_idx_)
func (_Store *StoreCaller) GetEvs(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getEvs")
	return *ret0, err
}

// GetEvs is a free data retrieval call binding the contract method 0xc5cb0a98.
//
// Solidity: function getEvs() constant returns(address[] current_evs_idx_)
func (_Store *StoreSession) GetEvs() ([]common.Address, error) {
	return _Store.Contract.GetEvs(&_Store.CallOpts)
}

// GetEvs is a free data retrieval call binding the contract method 0xc5cb0a98.
//
// Solidity: function getEvs() constant returns(address[] current_evs_idx_)
func (_Store *StoreCallerSession) GetEvs() ([]common.Address, error) {
	return _Store.Contract.GetEvs(&_Store.CallOpts)
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() constant returns(bytes32 genesis_hash_)
func (_Store *StoreCaller) GetGenesis(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "getGenesis")
	return *ret0, err
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() constant returns(bytes32 genesis_hash_)
func (_Store *StoreSession) GetGenesis() ([32]byte, error) {
	return _Store.Contract.GetGenesis(&_Store.CallOpts)
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() constant returns(bytes32 genesis_hash_)
func (_Store *StoreCallerSession) GetGenesis() ([32]byte, error) {
	return _Store.Contract.GetGenesis(&_Store.CallOpts)
}

// GetHeader is a free data retrieval call binding the contract method 0xb203e111.
//
// Solidity: function getHeader(uint64 _height) constant returns(bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreCaller) GetHeader(opts *bind.CallOpts, _height uint64) (struct {
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Sealer      common.Address
}, error) {
	ret := new(struct {
		Hash        [32]byte
		ParentHash  [32]byte
		ReceiptRoot [32]byte
		TxRoot      [32]byte
		StateRoot   [32]byte
		Sealer      common.Address
	})
	out := ret
	err := _Store.contract.Call(opts, out, "getHeader", _height)
	return *ret, err
}

// GetHeader is a free data retrieval call binding the contract method 0xb203e111.
//
// Solidity: function getHeader(uint64 _height) constant returns(bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreSession) GetHeader(_height uint64) (struct {
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Sealer      common.Address
}, error) {
	return _Store.Contract.GetHeader(&_Store.CallOpts, _height)
}

// GetHeader is a free data retrieval call binding the contract method 0xb203e111.
//
// Solidity: function getHeader(uint64 _height) constant returns(bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreCallerSession) GetHeader(_height uint64) (struct {
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Sealer      common.Address
}, error) {
	return _Store.Contract.GetHeader(&_Store.CallOpts, _height)
}

// IsEv is a free data retrieval call binding the contract method 0x9e186c93.
//
// Solidity: function isEv(address _ev) constant returns(bool)
func (_Store *StoreCaller) IsEv(opts *bind.CallOpts, _ev common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isEv", _ev)
	return *ret0, err
}

// IsEv is a free data retrieval call binding the contract method 0x9e186c93.
//
// Solidity: function isEv(address _ev) constant returns(bool)
func (_Store *StoreSession) IsEv(_ev common.Address) (bool, error) {
	return _Store.Contract.IsEv(&_Store.CallOpts, _ev)
}

// IsEv is a free data retrieval call binding the contract method 0x9e186c93.
//
// Solidity: function isEv(address _ev) constant returns(bool)
func (_Store *StoreCallerSession) IsEv(_ev common.Address) (bool, error) {
	return _Store.Contract.IsEv(&_Store.CallOpts, _ev)
}

// RecordHeader is a paid mutator transaction binding the contract method 0x3f10cc1e.
//
// Solidity: function recordHeader(uint64 _height, bytes32 _hash, bytes32 _parent_hash, bytes32 _receipt_root, bytes32 _tx_root, bytes32 _state_root, address _sealer, uint64 _chain_id) returns()
func (_Store *StoreTransactor) RecordHeader(opts *bind.TransactOpts, _height uint64, _hash [32]byte, _parent_hash [32]byte, _receipt_root [32]byte, _tx_root [32]byte, _state_root [32]byte, _sealer common.Address, _chain_id uint64) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "recordHeader", _height, _hash, _parent_hash, _receipt_root, _tx_root, _state_root, _sealer, _chain_id)
}

// RecordHeader is a paid mutator transaction binding the contract method 0x3f10cc1e.
//
// Solidity: function recordHeader(uint64 _height, bytes32 _hash, bytes32 _parent_hash, bytes32 _receipt_root, bytes32 _tx_root, bytes32 _state_root, address _sealer, uint64 _chain_id) returns()
func (_Store *StoreSession) RecordHeader(_height uint64, _hash [32]byte, _parent_hash [32]byte, _receipt_root [32]byte, _tx_root [32]byte, _state_root [32]byte, _sealer common.Address, _chain_id uint64) (*types.Transaction, error) {
	return _Store.Contract.RecordHeader(&_Store.TransactOpts, _height, _hash, _parent_hash, _receipt_root, _tx_root, _state_root, _sealer, _chain_id)
}

// RecordHeader is a paid mutator transaction binding the contract method 0x3f10cc1e.
//
// Solidity: function recordHeader(uint64 _height, bytes32 _hash, bytes32 _parent_hash, bytes32 _receipt_root, bytes32 _tx_root, bytes32 _state_root, address _sealer, uint64 _chain_id) returns()
func (_Store *StoreTransactorSession) RecordHeader(_height uint64, _hash [32]byte, _parent_hash [32]byte, _receipt_root [32]byte, _tx_root [32]byte, _state_root [32]byte, _sealer common.Address, _chain_id uint64) (*types.Transaction, error) {
	return _Store.Contract.RecordHeader(&_Store.TransactOpts, _height, _hash, _parent_hash, _receipt_root, _tx_root, _state_root, _sealer, _chain_id)
}

// SetEvs is a paid mutator transaction binding the contract method 0xdb829798.
//
// Solidity: function setEvs(address[] _new_evs) returns()
func (_Store *StoreTransactor) SetEvs(opts *bind.TransactOpts, _new_evs []common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setEvs", _new_evs)
}

// SetEvs is a paid mutator transaction binding the contract method 0xdb829798.
//
// Solidity: function setEvs(address[] _new_evs) returns()
func (_Store *StoreSession) SetEvs(_new_evs []common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetEvs(&_Store.TransactOpts, _new_evs)
}

// SetEvs is a paid mutator transaction binding the contract method 0xdb829798.
//
// Solidity: function setEvs(address[] _new_evs) returns()
func (_Store *StoreTransactorSession) SetEvs(_new_evs []common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetEvs(&_Store.TransactOpts, _new_evs)
}

// StoreRecordHeaderIterator is returned from FilterRecordHeader and is used to iterate over the raw logs and unpacked data for RecordHeader events raised by the Store contract.
type StoreRecordHeaderIterator struct {
	Event *StoreRecordHeader // Event containing the contract specifics and raw log

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
func (it *StoreRecordHeaderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRecordHeader)
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
		it.Event = new(StoreRecordHeader)
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
func (it *StoreRecordHeaderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRecordHeaderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRecordHeader represents a RecordHeader event raised by the Store contract.
type StoreRecordHeader struct {
	Height      uint64
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Sealer      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecordHeader is a free log retrieval operation binding the contract event 0xceec4c706f6e91cdd751ea972853f6862c309ad04ec22a41ca891868e495dd2c.
//
// Solidity: event RecordHeader(uint64 height, bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreFilterer) FilterRecordHeader(opts *bind.FilterOpts) (*StoreRecordHeaderIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RecordHeader")
	if err != nil {
		return nil, err
	}
	return &StoreRecordHeaderIterator{contract: _Store.contract, event: "RecordHeader", logs: logs, sub: sub}, nil
}

// WatchRecordHeader is a free log subscription operation binding the contract event 0xceec4c706f6e91cdd751ea972853f6862c309ad04ec22a41ca891868e495dd2c.
//
// Solidity: event RecordHeader(uint64 height, bytes32 hash, bytes32 parent_hash, bytes32 receipt_root, bytes32 tx_root, bytes32 state_root, address sealer)
func (_Store *StoreFilterer) WatchRecordHeader(opts *bind.WatchOpts, sink chan<- *StoreRecordHeader) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "RecordHeader")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRecordHeader)
				if err := _Store.contract.UnpackLog(event, "RecordHeader", log); err != nil {
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
