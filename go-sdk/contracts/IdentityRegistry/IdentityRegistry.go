// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IdentityRegistry

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
)

// IdentityRegistryMetaData contains all meta data concerning the IdentityRegistry contract.
var IdentityRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"NewManager\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"circuits\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"deploymentType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"deploymentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"deregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_circuitId\",\"type\":\"string\"}],\"name\":\"getCircuit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"identities\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identityManagerContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_circuitId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_deploymentType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_deploymentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"name\":\"setCircuit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506119738061005c5f395ff3fe608060405234801561000f575f80fd5b50600436106100a7575f3560e01c80639ea2cd381161006f5780639ea2cd381461014d578063d0ebdbe71461017f578063d4d2e7f21461019b578063e4baeda3146101cb578063f653b81e146101fd578063f6a3d24e1461022d576100a7565b806314afd79e146100ab57806324b8fbf6146100db5780635208ef6a146100f757806384ac33ec146101135780638da5cb5b1461012f575b5f80fd5b6100c560048036038101906100c09190610eb0565b61025d565b6040516100d29190610eea565b60405180910390f35b6100f560048036038101906100f0919061103f565b6102c2565b005b610111600480360381019061010c9190611137565b6104d0565b005b61012d60048036038101906101289190610eb0565b6105ee565b005b6101376106e8565b6040516101449190610eea565b60405180910390f35b610167600480360381019061016291906111ef565b61070b565b604051610176939291906112b0565b60405180910390f35b61019960048036038101906101949190610eb0565b610874565b005b6101b560048036038101906101b09190610eb0565b6109d8565b6040516101c29190610eea565b60405180910390f35b6101e560048036038101906101e091906111ef565b610a3d565b6040516101f4939291906112b0565b60405180910390f35b61021760048036038101906102129190610eb0565b610be5565b6040516102249190610eea565b60405180910390f35b61024760048036038101906102429190610eb0565b610c15565b604051610254919061130d565b60405180910390f35b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461038c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610383906113bc565b60405180910390fd5b5f6103978333610ca9565b90506103c3815f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684610cdb565b610402576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f990611424565b60405180910390fd5b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b846040516104c39190610eea565b60405180910390a2505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610526575f80fd5b60405180606001604052808481526020018373ffffffffffffffffffffffffffffffffffffffff16815260200182815250600285604051610567919061147c565b90815260200160405180910390205f820151815f0190816105889190611695565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190816105e49190611695565b5090505050505050565b5f3390503373ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610685575f80fd5b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002818051602081018201805184825260208301602085012081835280955050505050505f91509050805f018054610742906114bf565b80601f016020809104026020016040519081016040528092919081815260200182805461076e906114bf565b80156107b95780601f10610790576101008083540402835291602001916107b9565b820191905f5260205f20905b81548152906001019060200180831161079c57829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020180546107f3906114bf565b80601f016020809104026020016040519081016040528092919081815260200182805461081f906114bf565b801561086a5780601f106108415761010080835404028352916020019161086a565b820191905f5260205f20905b81548152906001019060200180831161084d57829003601f168201915b5050505050905083565b5f3390503373ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461090b575f80fd5b8160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b836040516109cc9190610eea565b60405180910390a25050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60605f6060600284604051610a52919061147c565b90815260200160405180910390205f01600285604051610a72919061147c565b90815260200160405180910390206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600286604051610ab3919061147c565b9081526020016040518091039020600201828054610ad0906114bf565b80601f0160208091040260200160405190810160405280929190818152602001828054610afc906114bf565b8015610b475780601f10610b1e57610100808354040283529160200191610b47565b820191905f5260205f20905b815481529060010190602001808311610b2a57829003601f168201915b50505050509250808054610b5a906114bf565b80601f0160208091040260200160405190810160405280929190818152602001828054610b86906114bf565b8015610bd15780601f10610ba857610100808354040283529160200191610bd1565b820191905f5260205f20905b815481529060010190602001808311610bb457829003601f168201915b505050505090509250925092509193909250565b6001602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b5f8282604051602001610cbd9291906117a9565b60405160208183030381529060405280519060200120905092915050565b5f806040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090505f8186604051602001610d29929190611841565b6040516020818303038152906040528051906020012090505f805f610d4d87610de0565b9250925092505f6001858386866040515f8152602001604052604051610d769493929190611892565b6020604051602081039080840390855afa158015610d96573d5f803e3d5ffd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161496505050505050509392505050565b5f805f6041845114610e27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1e9061191f565b60405180910390fd5b602084015192506040840151915060608401515f1a90509193909250565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e7f82610e56565b9050919050565b610e8f81610e75565b8114610e99575f80fd5b50565b5f81359050610eaa81610e86565b92915050565b5f60208284031215610ec557610ec4610e4e565b5b5f610ed284828501610e9c565b91505092915050565b610ee481610e75565b82525050565b5f602082019050610efd5f830184610edb565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f5182610f0b565b810181811067ffffffffffffffff82111715610f7057610f6f610f1b565b5b80604052505050565b5f610f82610e45565b9050610f8e8282610f48565b919050565b5f67ffffffffffffffff821115610fad57610fac610f1b565b5b610fb682610f0b565b9050602081019050919050565b828183375f83830152505050565b5f610fe3610fde84610f93565b610f79565b905082815260208101848484011115610fff57610ffe610f07565b5b61100a848285610fc3565b509392505050565b5f82601f83011261102657611025610f03565b5b8135611036848260208601610fd1565b91505092915050565b5f806040838503121561105557611054610e4e565b5b5f61106285828601610e9c565b925050602083013567ffffffffffffffff81111561108357611082610e52565b5b61108f85828601611012565b9150509250929050565b5f67ffffffffffffffff8211156110b3576110b2610f1b565b5b6110bc82610f0b565b9050602081019050919050565b5f6110db6110d684611099565b610f79565b9050828152602081018484840111156110f7576110f6610f07565b5b611102848285610fc3565b509392505050565b5f82601f83011261111e5761111d610f03565b5b813561112e8482602086016110c9565b91505092915050565b5f805f806080858703121561114f5761114e610e4e565b5b5f85013567ffffffffffffffff81111561116c5761116b610e52565b5b6111788782880161110a565b945050602085013567ffffffffffffffff81111561119957611198610e52565b5b6111a58782880161110a565b93505060406111b687828801610e9c565b925050606085013567ffffffffffffffff8111156111d7576111d6610e52565b5b6111e38782880161110a565b91505092959194509250565b5f6020828403121561120457611203610e4e565b5b5f82013567ffffffffffffffff81111561122157611220610e52565b5b61122d8482850161110a565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561126d578082015181840152602081019050611252565b5f8484015250505050565b5f61128282611236565b61128c8185611240565b935061129c818560208601611250565b6112a581610f0b565b840191505092915050565b5f6060820190508181035f8301526112c88186611278565b90506112d76020830185610edb565b81810360408301526112e98184611278565b9050949350505050565b5f8115159050919050565b611307816112f3565b82525050565b5f6020820190506113205f8301846112fe565b92915050565b7f416c726561647920726567697374657265642e20557365207365744d616e61675f8201527f65722063616c6c20746f20757064617465206d616e616765722061646472657360208201527f732e000000000000000000000000000000000000000000000000000000000000604082015250565b5f6113a6604283611240565b91506113b182611326565b606082019050919050565b5f6020820190508181035f8301526113d38161139a565b9050919050565b7f496e76616c6964207369676e61747572650000000000000000000000000000005f82015250565b5f61140e601183611240565b9150611419826113da565b602082019050919050565b5f6020820190508181035f83015261143b81611402565b9050919050565b5f81905092915050565b5f61145682611236565b6114608185611442565b9350611470818560208601611250565b80840191505092915050565b5f611487828461144c565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806114d657607f821691505b6020821081036114e9576114e8611492565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261154b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611510565b6115558683611510565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61159961159461158f8461156d565b611576565b61156d565b9050919050565b5f819050919050565b6115b28361157f565b6115c66115be826115a0565b84845461151c565b825550505050565b5f90565b6115da6115ce565b6115e58184846115a9565b505050565b5b81811015611608576115fd5f826115d2565b6001810190506115eb565b5050565b601f82111561164d5761161e816114ef565b61162784611501565b81016020851015611636578190505b61164a61164285611501565b8301826115ea565b50505b505050565b5f82821c905092915050565b5f61166d5f1984600802611652565b1980831691505092915050565b5f611685838361165e565b9150826002028217905092915050565b61169e82611236565b67ffffffffffffffff8111156116b7576116b6610f1b565b5b6116c182546114bf565b6116cc82828561160c565b5f60209050601f8311600181146116fd575f84156116eb578287015190505b6116f5858261167a565b86555061175c565b601f19841661170b866114ef565b5f5b828110156117325784890151825560018201915060208501945060208101905061170d565b8683101561174f578489015161174b601f89168261165e565b8355505b6001600288020188555050505b505050505050565b5f8160601b9050919050565b5f61177a82611764565b9050919050565b5f61178b82611770565b9050919050565b6117a361179e82610e75565b611781565b82525050565b5f6117b48285611792565b6014820191506117c48284611792565b6014820191508190509392505050565b5f81519050919050565b5f81905092915050565b5f6117f2826117d4565b6117fc81856117de565b935061180c818560208601611250565b80840191505092915050565b5f819050919050565b5f819050919050565b61183b61183682611818565b611821565b82525050565b5f61184c82856117e8565b9150611858828461182a565b6020820191508190509392505050565b61187181611818565b82525050565b5f60ff82169050919050565b61188c81611877565b82525050565b5f6080820190506118a55f830187611868565b6118b26020830186611883565b6118bf6040830185611868565b6118cc6060830184611868565b95945050505050565b7f696e76616c6964207369676e6174757265206c656e67746800000000000000005f82015250565b5f611909601883611240565b9150611914826118d5565b602082019050919050565b5f6020820190508181035f830152611936816118fd565b905091905056fea2646970667358221220483374c5343093764d5c2734fcdfe03d67e9af9c4773c23fdf287a1854be5e9e64736f6c63430008140033",
}

// IdentityRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityRegistryMetaData.ABI instead.
var IdentityRegistryABI = IdentityRegistryMetaData.ABI

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityRegistryMetaData.Bin instead.
var IdentityRegistryBin = IdentityRegistryMetaData.Bin

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := IdentityRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// Circuits is a free data retrieval call binding the contract method 0x9ea2cd38.
//
// Solidity: function circuits(string ) view returns(string deploymentType, address deploymentAddress, string ipfsHash)
func (_IdentityRegistry *IdentityRegistryCaller) Circuits(opts *bind.CallOpts, arg0 string) (struct {
	DeploymentType    string
	DeploymentAddress common.Address
	IpfsHash          string
}, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "circuits", arg0)

	outstruct := new(struct {
		DeploymentType    string
		DeploymentAddress common.Address
		IpfsHash          string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DeploymentType = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DeploymentAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IpfsHash = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Circuits is a free data retrieval call binding the contract method 0x9ea2cd38.
//
// Solidity: function circuits(string ) view returns(string deploymentType, address deploymentAddress, string ipfsHash)
func (_IdentityRegistry *IdentityRegistrySession) Circuits(arg0 string) (struct {
	DeploymentType    string
	DeploymentAddress common.Address
	IpfsHash          string
}, error) {
	return _IdentityRegistry.Contract.Circuits(&_IdentityRegistry.CallOpts, arg0)
}

// Circuits is a free data retrieval call binding the contract method 0x9ea2cd38.
//
// Solidity: function circuits(string ) view returns(string deploymentType, address deploymentAddress, string ipfsHash)
func (_IdentityRegistry *IdentityRegistryCallerSession) Circuits(arg0 string) (struct {
	DeploymentType    string
	DeploymentAddress common.Address
	IpfsHash          string
}, error) {
	return _IdentityRegistry.Contract.Circuits(&_IdentityRegistry.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) Exists(opts *bind.CallOpts, node common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "exists", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) Exists(node common.Address) (bool, error) {
	return _IdentityRegistry.Contract.Exists(&_IdentityRegistry.CallOpts, node)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address node) view returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) Exists(node common.Address) (bool, error) {
	return _IdentityRegistry.Contract.Exists(&_IdentityRegistry.CallOpts, node)
}

// GetCircuit is a free data retrieval call binding the contract method 0xe4baeda3.
//
// Solidity: function getCircuit(string _circuitId) view returns(string, address, string)
func (_IdentityRegistry *IdentityRegistryCaller) GetCircuit(opts *bind.CallOpts, _circuitId string) (string, common.Address, string, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "getCircuit", _circuitId)

	if err != nil {
		return *new(string), *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetCircuit is a free data retrieval call binding the contract method 0xe4baeda3.
//
// Solidity: function getCircuit(string _circuitId) view returns(string, address, string)
func (_IdentityRegistry *IdentityRegistrySession) GetCircuit(_circuitId string) (string, common.Address, string, error) {
	return _IdentityRegistry.Contract.GetCircuit(&_IdentityRegistry.CallOpts, _circuitId)
}

// GetCircuit is a free data retrieval call binding the contract method 0xe4baeda3.
//
// Solidity: function getCircuit(string _circuitId) view returns(string, address, string)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetCircuit(_circuitId string) (string, common.Address, string, error) {
	return _IdentityRegistry.Contract.GetCircuit(&_IdentityRegistry.CallOpts, _circuitId)
}

// Identities is a free data retrieval call binding the contract method 0xf653b81e.
//
// Solidity: function identities(address ) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Identities(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "identities", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Identities is a free data retrieval call binding the contract method 0xf653b81e.
//
// Solidity: function identities(address ) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Identities(arg0 common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Identities(&_IdentityRegistry.CallOpts, arg0)
}

// Identities is a free data retrieval call binding the contract method 0xf653b81e.
//
// Solidity: function identities(address ) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Identities(arg0 common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Identities(&_IdentityRegistry.CallOpts, arg0)
}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Manager(opts *bind.CallOpts, node common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "manager", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Manager(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Manager(&_IdentityRegistry.CallOpts, node)
}

// Manager is a free data retrieval call binding the contract method 0xd4d2e7f2.
//
// Solidity: function manager(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Manager(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.Manager(&_IdentityRegistry.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) OwnerOf(opts *bind.CallOpts, node common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistry.contract.Call(opts, &out, "ownerOf", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistrySession) OwnerOf(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.OwnerOf(&_IdentityRegistry.CallOpts, node)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address node) view returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) OwnerOf(node common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.OwnerOf(&_IdentityRegistry.CallOpts, node)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address node) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Deregister(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "deregister", node)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address node) returns()
func (_IdentityRegistry *IdentityRegistrySession) Deregister(node common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Deregister(&_IdentityRegistry.TransactOpts, node)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(address node) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Deregister(node common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Deregister(&_IdentityRegistry.TransactOpts, node)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) Register(opts *bind.TransactOpts, identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "register", identityManagerContract, signature)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityRegistry *IdentityRegistrySession) Register(identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register(&_IdentityRegistry.TransactOpts, identityManagerContract, signature)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address identityManagerContract, bytes signature) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) Register(identityManagerContract common.Address, signature []byte) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.Register(&_IdentityRegistry.TransactOpts, identityManagerContract, signature)
}

// SetCircuit is a paid mutator transaction binding the contract method 0x5208ef6a.
//
// Solidity: function setCircuit(string _circuitId, string _deploymentType, address _deploymentAddress, string _ipfsHash) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetCircuit(opts *bind.TransactOpts, _circuitId string, _deploymentType string, _deploymentAddress common.Address, _ipfsHash string) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setCircuit", _circuitId, _deploymentType, _deploymentAddress, _ipfsHash)
}

// SetCircuit is a paid mutator transaction binding the contract method 0x5208ef6a.
//
// Solidity: function setCircuit(string _circuitId, string _deploymentType, address _deploymentAddress, string _ipfsHash) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetCircuit(_circuitId string, _deploymentType string, _deploymentAddress common.Address, _ipfsHash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetCircuit(&_IdentityRegistry.TransactOpts, _circuitId, _deploymentType, _deploymentAddress, _ipfsHash)
}

// SetCircuit is a paid mutator transaction binding the contract method 0x5208ef6a.
//
// Solidity: function setCircuit(string _circuitId, string _deploymentType, address _deploymentAddress, string _ipfsHash) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetCircuit(_circuitId string, _deploymentType string, _deploymentAddress common.Address, _ipfsHash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetCircuit(&_IdentityRegistry.TransactOpts, _circuitId, _deploymentType, _deploymentAddress, _ipfsHash)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setManager", manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetManager(&_IdentityRegistry.TransactOpts, manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetManager(&_IdentityRegistry.TransactOpts, manager)
}

// IdentityRegistryNewManagerIterator is returned from FilterNewManager and is used to iterate over the raw logs and unpacked data for NewManager events raised by the IdentityRegistry contract.
type IdentityRegistryNewManagerIterator struct {
	Event *IdentityRegistryNewManager // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryNewManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryNewManager)
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
		it.Event = new(IdentityRegistryNewManager)
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
func (it *IdentityRegistryNewManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryNewManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryNewManager represents a NewManager event raised by the IdentityRegistry contract.
type IdentityRegistryNewManager struct {
	Node    common.Address
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewManager is a free log retrieval operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterNewManager(opts *bind.FilterOpts, node []common.Address) (*IdentityRegistryNewManagerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "NewManager", nodeRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryNewManagerIterator{contract: _IdentityRegistry.contract, event: "NewManager", logs: logs, sub: sub}, nil
}

// WatchNewManager is a free log subscription operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchNewManager(opts *bind.WatchOpts, sink chan<- *IdentityRegistryNewManager, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "NewManager", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryNewManager)
				if err := _IdentityRegistry.contract.UnpackLog(event, "NewManager", log); err != nil {
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

// ParseNewManager is a log parse operation binding the contract event 0x770e6248a70b6ac757edf422766216da592c37e3112db900fe0da8984191831b.
//
// Solidity: event NewManager(address indexed node, address manager)
func (_IdentityRegistry *IdentityRegistryFilterer) ParseNewManager(log types.Log) (*IdentityRegistryNewManager, error) {
	event := new(IdentityRegistryNewManager)
	if err := _IdentityRegistry.contract.UnpackLog(event, "NewManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
