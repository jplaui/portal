// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IdentityManager

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

// IdentityManagerMetaData contains all meta data concerning the IdentityManager contract.
var IdentityManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIdentityInterface\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"AttestationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NewAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPrivate\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"version\",\"type\":\"int256\"}],\"name\":\"NewClaim\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"attestations\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"attestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"claimId\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isPrivate\",\"type\":\"bool\"}],\"name\":\"deleteClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"deleteClaimURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ipfsClaims\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"privateClaims\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsCircuitMetadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"eventHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"version\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"publicClaims\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"version\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIdentityInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"revocations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"attestedTo\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"attestationId\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"attestedTo\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"attestationId\",\"type\":\"int256\"}],\"name\":\"revokeAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"attestor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"int256\",\"name\":\"claimId\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"setAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setClaimURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"eventHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"identifier\",\"type\":\"int256\"}],\"name\":\"setPrivateClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"statement\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"setPublicClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200299738038062002997833981810160405281019062000036919062000134565b3360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000164565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620000eb82620000c0565b9050919050565b5f620000fe82620000df565b9050919050565b6200011081620000f2565b81146200011b575f80fd5b50565b5f815190506200012e8162000105565b92915050565b5f602082840312156200014c576200014b620000bc565b5b5f6200015b848285016200011e565b91505092915050565b61282580620001725f395ff3fe608060405234801561000f575f80fd5b50600436106100e8575f3560e01c806384ce394c1161008a578063c87c6d4111610064578063c87c6d4114610266578063dbbd60d914610282578063ead05a341461029e578063ee10c7e2146102ba576100e8565b806384ce394c146102105780638da5cb5b1461022c578063c03bc8b51461024a576100e8565b80635d91ca9f116100c65780635d91ca9f1461016d578063755af686146101895780637b103999146101bf5780638107769b146101dd576100e8565b8063249d026b146100ec57806345ddaf5d146101215780634665b25814610151575b5f80fd5b61010660048036038101906101019190611697565b6102ee565b604051610118969594939291906117c7565b60405180910390f35b61013b60048036038101906101369190611697565b6103e3565b604051610148919061187f565b60405180910390f35b61016b60048036038101906101669190611931565b610496565b005b610187600480360381019061018291906119e2565b6105c7565b005b6101a3600480360381019061019e9190611697565b6107c2565b6040516101b69796959493929190611ab9565b60405180910390f35b6101c7610a30565b6040516101d49190611b9d565b60405180910390f35b6101f760048036038101906101f29190611697565b610a53565b6040516102079493929190611bb6565b60405180910390f35b61022a60048036038101906102259190611c00565b610b3c565b005b610234610bc3565b6040516102419190611c4b565b60405180910390f35b610264600480360381019061025f9190611c64565b610be8565b005b610280600480360381019061027b9190611d0c565b610cb7565b005b61029c60048036038101906102979190611e04565b610ed9565b005b6102b860048036038101906102b39190611ece565b6110e6565b005b6102d460048036038101906102cf9190611697565b61139c565b6040516102e5959493929190612007565b60405180910390f35b6006818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461032590612093565b80601f016020809104026020016040519081016040528092919081815260200182805461035190612093565b801561039c5780601f106103735761010080835404028352916020019161039c565b820191905f5260205f20905b81548152906001019060200180831161037f57829003601f168201915b505050505090806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154908060050154905086565b6004818051602081018201805184825260208301602085012081835280955050505050505f91509050805461041790612093565b80601f016020809104026020016040519081016040528092919081815260200182805461044390612093565b801561048e5780601f106104655761010080835404028352916020019161048e565b820191905f5260205f20905b81548152906001019060200180831161047157829003601f168201915b505050505081565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ee575f80fd5b801561056c57600283836040516105069291906120f1565b90815260200160405180910390205f8082015f61052391906114f2565b600182015f61053291906114f2565b600282015f61054191906114f2565b600382015f61055091906114f2565b600482015f9055600582015f9055600682015f905550506105c2565b6003838360405161057e9291906120f1565b90815260200160405180910390205f8082015f61059b91906114f2565b600182015f6105aa91906114f2565b600282015f9055600382015f9055600482015f905550505b505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461061f575f80fd5b5f600389896040516106329291906120f1565b90815260200160405180910390206004015490506040518060a0016040528088888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018481526020018381526020016001836106fc9190612136565b81525060038a8a6040516107119291906120f1565b90815260200160405180910390205f820151815f019081610732919061230b565b506020820151816001019081610748919061230b565b506040820151816002015560608201518160030155608082015181600401559050507f26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62898989895f60018761079d9190612136565b6040516107af96959493929190612415565b60405180910390a1505050505050505050565b6002818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546107f990612093565b80601f016020809104026020016040519081016040528092919081815260200182805461082590612093565b80156108705780601f1061084757610100808354040283529160200191610870565b820191905f5260205f20905b81548152906001019060200180831161085357829003601f168201915b50505050509080600101805461088590612093565b80601f01602080910402602001604051908101604052809291908181526020018280546108b190612093565b80156108fc5780601f106108d3576101008083540402835291602001916108fc565b820191905f5260205f20905b8154815290600101906020018083116108df57829003601f168201915b50505050509080600201805461091190612093565b80601f016020809104026020016040519081016040528092919081815260200182805461093d90612093565b80156109885780601f1061095f57610100808354040283529160200191610988565b820191905f5260205f20905b81548152906001019060200180831161096b57829003601f168201915b50505050509080600301805461099d90612093565b80601f01602080910402602001604051908101604052809291908181526020018280546109c990612093565b8015610a145780601f106109eb57610100808354040283529160200191610a14565b820191905f5260205f20905b8154815290600101906020018083116109f757829003601f168201915b5050505050908060040154908060050154908060060154905087565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005818051602081018201805184825260208301602085012081835280955050505050505f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054610ab590612093565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae190612093565b8015610b2c5780601f10610b0357610100808354040283529160200191610b2c565b820191905f5260205f20905b815481529060010190602001808311610b0f57829003601f168201915b5050505050908060030154905084565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b94575f80fd5b60048282604051610ba69291906120f1565b90815260200160405180910390205f610bbf91906114f2565b5050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c40575f80fd5b818160048686604051610c549291906120f1565b90815260200160405180910390209182610c6f929190612474565b507f26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62848484845f80604051610ca99695949392919061257a565b60405180910390a150505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d0f575f80fd5b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281525060058787604051610d9b9291906120f1565b90815260200160405180910390205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002019081610e0c919061230b565b506060820151816003015590505060068686604051610e2c9291906120f1565b90815260200160405180910390206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f613a02027bb8290f4136e1ca1f3f2d877cf32d609bbc9cfa4799ad69b8e2a08b60068888604051610ea69291906120f1565b9081526020016040518091039020600501548686604051610ec9939291906125cf565b60405180910390a2505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f31575f80fd5b5f60068989604051610f449291906120f1565b90815260200160405180910390206003015414610f5f575f80fd5b6040518060c0016040528085858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020014281526020018381526020018281525060068989604051610ff79291906120f1565b90815260200160405180910390205f820151815f0190816110189190612657565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015560a082015181600501559050508573ffffffffffffffffffffffffffffffffffffffff167f9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a8989886040516110d493929190612726565b60405180910390a25050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461113e575f80fd5b5f60028c60405161114f9190612786565b90815260200160405180910390206006015490506040518060e001604052808c8c8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018481526020018381526020016001836112ab9190612136565b81525060028d6040516112be9190612786565b90815260200160405180910390205f820151815f0190816112df919061230b565b5060208201518160010190816112f5919061230b565b50604082015181600201908161130b919061230b565b506060820151816003019081611321919061230b565b506080820151816004015560a0820151816005015560c082015181600601559050507f26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d628c8c8c600180866113759190612136565b60405161138695949392919061279c565b60405180910390a1505050505050505050505050565b6003818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546113d390612093565b80601f01602080910402602001604051908101604052809291908181526020018280546113ff90612093565b801561144a5780601f106114215761010080835404028352916020019161144a565b820191905f5260205f20905b81548152906001019060200180831161142d57829003601f168201915b50505050509080600101805461145f90612093565b80601f016020809104026020016040519081016040528092919081815260200182805461148b90612093565b80156114d65780601f106114ad576101008083540402835291602001916114d6565b820191905f5260205f20905b8154815290600101906020018083116114b957829003601f168201915b5050505050908060020154908060030154908060040154905085565b5080546114fe90612093565b5f825580601f1061150f575061152c565b601f0160209004905f5260205f209081019061152b919061152f565b5b50565b5b80821115611546575f815f905550600101611530565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6115a982611563565b810181811067ffffffffffffffff821117156115c8576115c7611573565b5b80604052505050565b5f6115da61154a565b90506115e682826115a0565b919050565b5f67ffffffffffffffff82111561160557611604611573565b5b61160e82611563565b9050602081019050919050565b828183375f83830152505050565b5f61163b611636846115eb565b6115d1565b9050828152602081018484840111156116575761165661155f565b5b61166284828561161b565b509392505050565b5f82601f83011261167e5761167d61155b565b5b813561168e848260208601611629565b91505092915050565b5f602082840312156116ac576116ab611553565b5b5f82013567ffffffffffffffff8111156116c9576116c8611557565b5b6116d58482850161166a565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156117155780820151818401526020810190506116fa565b5f8484015250505050565b5f61172a826116de565b61173481856116e8565b93506117448185602086016116f8565b61174d81611563565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61178182611758565b9050919050565b61179181611777565b82525050565b5f819050919050565b6117a981611797565b82525050565b5f819050919050565b6117c1816117af565b82525050565b5f60c0820190508181035f8301526117df8189611720565b90506117ee6020830188611788565b6117fb60408301876117a0565b61180860608301866117a0565b61181560808301856117b8565b61182260a08301846117b8565b979650505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f6118518261182d565b61185b8185611837565b935061186b8185602086016116f8565b61187481611563565b840191505092915050565b5f6020820190508181035f8301526118978184611847565b905092915050565b5f80fd5b5f80fd5b5f8083601f8401126118bc576118bb61155b565b5b8235905067ffffffffffffffff8111156118d9576118d861189f565b5b6020830191508360018202830111156118f5576118f46118a3565b5b9250929050565b5f8115159050919050565b611910816118fc565b811461191a575f80fd5b50565b5f8135905061192b81611907565b92915050565b5f805f6040848603121561194857611947611553565b5b5f84013567ffffffffffffffff81111561196557611964611557565b5b611971868287016118a7565b935093505060206119848682870161191d565b9150509250925092565b61199781611797565b81146119a1575f80fd5b50565b5f813590506119b28161198e565b92915050565b6119c1816117af565b81146119cb575f80fd5b50565b5f813590506119dc816119b8565b92915050565b5f805f805f805f8060a0898b0312156119fe576119fd611553565b5b5f89013567ffffffffffffffff811115611a1b57611a1a611557565b5b611a278b828c016118a7565b9850985050602089013567ffffffffffffffff811115611a4a57611a49611557565b5b611a568b828c016118a7565b9650965050604089013567ffffffffffffffff811115611a7957611a78611557565b5b611a858b828c016118a7565b94509450506060611a988b828c016119a4565b9250506080611aa98b828c016119ce565b9150509295985092959890939650565b5f60e0820190508181035f830152611ad1818a611847565b90508181036020830152611ae58189611847565b90508181036040830152611af98188611847565b90508181036060830152611b0d8187611847565b9050611b1c60808301866117a0565b611b2960a08301856117b8565b611b3660c08301846117b8565b98975050505050505050565b5f819050919050565b5f611b65611b60611b5b84611758565b611b42565b611758565b9050919050565b5f611b7682611b4b565b9050919050565b5f611b8782611b6c565b9050919050565b611b9781611b7d565b82525050565b5f602082019050611bb05f830184611b8e565b92915050565b5f608082019050611bc95f830187611788565b611bd660208301866117b8565b8181036040830152611be88185611847565b9050611bf760608301846117a0565b95945050505050565b5f8060208385031215611c1657611c15611553565b5b5f83013567ffffffffffffffff811115611c3357611c32611557565b5b611c3f858286016118a7565b92509250509250929050565b5f602082019050611c5e5f830184611788565b92915050565b5f805f8060408587031215611c7c57611c7b611553565b5b5f85013567ffffffffffffffff811115611c9957611c98611557565b5b611ca5878288016118a7565b9450945050602085013567ffffffffffffffff811115611cc857611cc7611557565b5b611cd4878288016118a7565b925092505092959194509250565b611ceb81611777565b8114611cf5575f80fd5b50565b5f81359050611d0681611ce2565b92915050565b5f805f805f8060808789031215611d2657611d25611553565b5b5f87013567ffffffffffffffff811115611d4357611d42611557565b5b611d4f89828a016118a7565b9650965050602087013567ffffffffffffffff811115611d7257611d71611557565b5b611d7e89828a016118a7565b94509450506040611d9189828a01611cf8565b9250506060611da289828a016119ce565b9150509295509295509295565b5f8083601f840112611dc457611dc361155b565b5b8235905067ffffffffffffffff811115611de157611de061189f565b5b602083019150836001820283011115611dfd57611dfc6118a3565b5b9250929050565b5f805f805f805f8060c0898b031215611e2057611e1f611553565b5b5f89013567ffffffffffffffff811115611e3d57611e3c611557565b5b611e498b828c016118a7565b98509850506020611e5c8b828c01611cf8565b9650506040611e6d8b828c016119a4565b955050606089013567ffffffffffffffff811115611e8e57611e8d611557565b5b611e9a8b828c01611daf565b94509450506080611ead8b828c016119ce565b92505060a0611ebe8b828c016119ce565b9150509295985092959890939650565b5f805f805f805f805f805f60e08c8e031215611eed57611eec611553565b5b5f8c013567ffffffffffffffff811115611f0a57611f09611557565b5b611f168e828f0161166a565b9b505060208c013567ffffffffffffffff811115611f3757611f36611557565b5b611f438e828f016118a7565b9a509a505060408c013567ffffffffffffffff811115611f6657611f65611557565b5b611f728e828f016118a7565b985098505060608c013567ffffffffffffffff811115611f9557611f94611557565b5b611fa18e828f016118a7565b965096505060808c013567ffffffffffffffff811115611fc457611fc3611557565b5b611fd08e828f016118a7565b945094505060a0611fe38e828f016119a4565b92505060c0611ff48e828f016119ce565b9150509295989b509295989b9093969950565b5f60a0820190508181035f83015261201f8188611847565b905081810360208301526120338187611847565b905061204260408301866117a0565b61204f60608301856117b8565b61205c60808301846117b8565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806120aa57607f821691505b6020821081036120bd576120bc612066565b5b50919050565b5f81905092915050565b5f6120d883856120c3565b93506120e583858461161b565b82840190509392505050565b5f6120fd8284866120cd565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612140826117af565b915061214b836117af565b92508282019050828112155f8312168382125f84121516171561217157612170612109565b5b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026121d37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612198565b6121dd8683612198565b95508019841693508086168417925050509392505050565b5f61220f61220a61220584611797565b611b42565b611797565b9050919050565b5f819050919050565b612228836121f5565b61223c61223482612216565b8484546121a4565b825550505050565b5f90565b612250612244565b61225b81848461221f565b505050565b5b8181101561227e576122735f82612248565b600181019050612261565b5050565b601f8211156122c35761229481612177565b61229d84612189565b810160208510156122ac578190505b6122c06122b885612189565b830182612260565b50505b505050565b5f82821c905092915050565b5f6122e35f19846008026122c8565b1980831691505092915050565b5f6122fb83836122d4565b9150826002028217905092915050565b6123148261182d565b67ffffffffffffffff81111561232d5761232c611573565b5b6123378254612093565b612342828285612282565b5f60209050601f831160018114612373575f8415612361578287015190505b61236b85826122f0565b8655506123d2565b601f19841661238186612177565b5f5b828110156123a857848901518255600182019150602085019450602081019050612383565b868310156123c557848901516123c1601f8916826122d4565b8355505b6001600288020188555050505b505050505050565b5f6123e58385611837565b93506123f283858461161b565b6123fb83611563565b840190509392505050565b61240f816118fc565b82525050565b5f6080820190508181035f83015261242e81888a6123da565b905081810360208301526124438186886123da565b90506124526040830185612406565b61245f60608301846117b8565b979650505050505050565b5f82905092915050565b61247e838361246a565b67ffffffffffffffff81111561249757612496611573565b5b6124a18254612093565b6124ac828285612282565b5f601f8311600181146124d9575f84156124c7578287013590505b6124d185826122f0565b865550612538565b601f1984166124e786612177565b5f5b8281101561250e578489013582556001820191506020850194506020810190506124e9565b8683101561252b5784890135612527601f8916826122d4565b8355505b6001600288020188555050505b50505050505050565b5f819050919050565b5f61256461255f61255a84612541565b611b42565b6117af565b9050919050565b6125748161254a565b82525050565b5f6080820190508181035f83015261259381888a6123da565b905081810360208301526125a88186886123da565b90506125b76040830185612406565b6125c4606083018461256b565b979650505050505050565b5f6040820190506125e25f8301866117b8565b81810360208301526125f58184866123da565b9050949350505050565b5f819050815f5260205f209050919050565b601f82111561265257612623816125ff565b61262c84612189565b8101602085101561263b578190505b61264f61264785612189565b830182612260565b50505b505050565b612660826116de565b67ffffffffffffffff81111561267957612678611573565b5b6126838254612093565b61268e828285612611565b5f60209050601f8311600181146126bf575f84156126ad578287015190505b6126b785826122f0565b86555061271e565b601f1984166126cd866125ff565b5f5b828110156126f4578489015182556001820191506020850194506020810190506126cf565b86831015612711578489015161270d601f8916826122d4565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f83015261273f8185876123da565b905061274e60208301846117a0565b949350505050565b5f6127608261182d565b61276a81856120c3565b935061277a8185602086016116f8565b80840191505092915050565b5f6127918284612756565b915081905092915050565b5f6080820190508181035f8301526127b48188611847565b905081810360208301526127c98186886123da565b90506127d86040830185612406565b6127e560608301846117b8565b969550505050505056fea26469706673582212208266b831bd047654d6ad419189dcdba0c530073843035942ff0a7034478d30f064736f6c63430008140033",
}

// IdentityManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityManagerMetaData.ABI instead.
var IdentityManagerABI = IdentityManagerMetaData.ABI

// IdentityManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityManagerMetaData.Bin instead.
var IdentityManagerBin = IdentityManagerMetaData.Bin

// DeployIdentityManager deploys a new Ethereum contract, binding an instance of IdentityManager to it.
func DeployIdentityManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *IdentityManager, error) {
	parsed, err := IdentityManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityManagerBin), backend, _registry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityManager{IdentityManagerCaller: IdentityManagerCaller{contract: contract}, IdentityManagerTransactor: IdentityManagerTransactor{contract: contract}, IdentityManagerFilterer: IdentityManagerFilterer{contract: contract}}, nil
}

// IdentityManager is an auto generated Go binding around an Ethereum contract.
type IdentityManager struct {
	IdentityManagerCaller     // Read-only binding to the contract
	IdentityManagerTransactor // Write-only binding to the contract
	IdentityManagerFilterer   // Log filterer for contract events
}

// IdentityManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityManagerSession struct {
	Contract     *IdentityManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityManagerCallerSession struct {
	Contract *IdentityManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentityManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityManagerTransactorSession struct {
	Contract     *IdentityManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentityManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityManagerRaw struct {
	Contract *IdentityManager // Generic contract binding to access the raw methods on
}

// IdentityManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityManagerCallerRaw struct {
	Contract *IdentityManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityManagerTransactorRaw struct {
	Contract *IdentityManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityManager creates a new instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManager(address common.Address, backend bind.ContractBackend) (*IdentityManager, error) {
	contract, err := bindIdentityManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityManager{IdentityManagerCaller: IdentityManagerCaller{contract: contract}, IdentityManagerTransactor: IdentityManagerTransactor{contract: contract}, IdentityManagerFilterer: IdentityManagerFilterer{contract: contract}}, nil
}

// NewIdentityManagerCaller creates a new read-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerCaller(address common.Address, caller bind.ContractCaller) (*IdentityManagerCaller, error) {
	contract, err := bindIdentityManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerCaller{contract: contract}, nil
}

// NewIdentityManagerTransactor creates a new write-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityManagerTransactor, error) {
	contract, err := bindIdentityManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerTransactor{contract: contract}, nil
}

// NewIdentityManagerFilterer creates a new log filterer instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityManagerFilterer, error) {
	contract, err := bindIdentityManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerFilterer{contract: contract}, nil
}

// bindIdentityManager binds a generic wrapper to an already deployed contract.
func bindIdentityManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.IdentityManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transact(opts, method, params...)
}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(bytes signature, address attestor, uint256 expires, uint256 timestamp, int256 claimId, int256 id)
func (_IdentityManager *IdentityManagerCaller) Attestations(opts *bind.CallOpts, arg0 string) (struct {
	Signature []byte
	Attestor  common.Address
	Expires   *big.Int
	Timestamp *big.Int
	ClaimId   *big.Int
	Id        *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "attestations", arg0)

	outstruct := new(struct {
		Signature []byte
		Attestor  common.Address
		Expires   *big.Int
		Timestamp *big.Int
		ClaimId   *big.Int
		Id        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signature = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Attestor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Expires = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ClaimId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(bytes signature, address attestor, uint256 expires, uint256 timestamp, int256 claimId, int256 id)
func (_IdentityManager *IdentityManagerSession) Attestations(arg0 string) (struct {
	Signature []byte
	Attestor  common.Address
	Expires   *big.Int
	Timestamp *big.Int
	ClaimId   *big.Int
	Id        *big.Int
}, error) {
	return _IdentityManager.Contract.Attestations(&_IdentityManager.CallOpts, arg0)
}

// Attestations is a free data retrieval call binding the contract method 0x249d026b.
//
// Solidity: function attestations(string ) view returns(bytes signature, address attestor, uint256 expires, uint256 timestamp, int256 claimId, int256 id)
func (_IdentityManager *IdentityManagerCallerSession) Attestations(arg0 string) (struct {
	Signature []byte
	Attestor  common.Address
	Expires   *big.Int
	Timestamp *big.Int
	ClaimId   *big.Int
	Id        *big.Int
}, error) {
	return _IdentityManager.Contract.Attestations(&_IdentityManager.CallOpts, arg0)
}

// IpfsClaims is a free data retrieval call binding the contract method 0x45ddaf5d.
//
// Solidity: function ipfsClaims(string ) view returns(string)
func (_IdentityManager *IdentityManagerCaller) IpfsClaims(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "ipfsClaims", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IpfsClaims is a free data retrieval call binding the contract method 0x45ddaf5d.
//
// Solidity: function ipfsClaims(string ) view returns(string)
func (_IdentityManager *IdentityManagerSession) IpfsClaims(arg0 string) (string, error) {
	return _IdentityManager.Contract.IpfsClaims(&_IdentityManager.CallOpts, arg0)
}

// IpfsClaims is a free data retrieval call binding the contract method 0x45ddaf5d.
//
// Solidity: function ipfsClaims(string ) view returns(string)
func (_IdentityManager *IdentityManagerCallerSession) IpfsClaims(arg0 string) (string, error) {
	return _IdentityManager.Contract.IpfsClaims(&_IdentityManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityManager *IdentityManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityManager *IdentityManagerSession) Owner() (common.Address, error) {
	return _IdentityManager.Contract.Owner(&_IdentityManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityManager *IdentityManagerCallerSession) Owner() (common.Address, error) {
	return _IdentityManager.Contract.Owner(&_IdentityManager.CallOpts)
}

// PrivateClaims is a free data retrieval call binding the contract method 0x755af686.
//
// Solidity: function privateClaims(string ) view returns(string value, string statement, string ipfsCircuitMetadata, string eventHash, uint256 timestamp, int256 id, int256 version)
func (_IdentityManager *IdentityManagerCaller) PrivateClaims(opts *bind.CallOpts, arg0 string) (struct {
	Value               string
	Statement           string
	IpfsCircuitMetadata string
	EventHash           string
	Timestamp           *big.Int
	Id                  *big.Int
	Version             *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "privateClaims", arg0)

	outstruct := new(struct {
		Value               string
		Statement           string
		IpfsCircuitMetadata string
		EventHash           string
		Timestamp           *big.Int
		Id                  *big.Int
		Version             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Statement = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IpfsCircuitMetadata = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.EventHash = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PrivateClaims is a free data retrieval call binding the contract method 0x755af686.
//
// Solidity: function privateClaims(string ) view returns(string value, string statement, string ipfsCircuitMetadata, string eventHash, uint256 timestamp, int256 id, int256 version)
func (_IdentityManager *IdentityManagerSession) PrivateClaims(arg0 string) (struct {
	Value               string
	Statement           string
	IpfsCircuitMetadata string
	EventHash           string
	Timestamp           *big.Int
	Id                  *big.Int
	Version             *big.Int
}, error) {
	return _IdentityManager.Contract.PrivateClaims(&_IdentityManager.CallOpts, arg0)
}

// PrivateClaims is a free data retrieval call binding the contract method 0x755af686.
//
// Solidity: function privateClaims(string ) view returns(string value, string statement, string ipfsCircuitMetadata, string eventHash, uint256 timestamp, int256 id, int256 version)
func (_IdentityManager *IdentityManagerCallerSession) PrivateClaims(arg0 string) (struct {
	Value               string
	Statement           string
	IpfsCircuitMetadata string
	EventHash           string
	Timestamp           *big.Int
	Id                  *big.Int
	Version             *big.Int
}, error) {
	return _IdentityManager.Contract.PrivateClaims(&_IdentityManager.CallOpts, arg0)
}

// PublicClaims is a free data retrieval call binding the contract method 0xee10c7e2.
//
// Solidity: function publicClaims(string ) view returns(string value, string statement, uint256 timestamp, int256 id, int256 version)
func (_IdentityManager *IdentityManagerCaller) PublicClaims(opts *bind.CallOpts, arg0 string) (struct {
	Value     string
	Statement string
	Timestamp *big.Int
	Id        *big.Int
	Version   *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "publicClaims", arg0)

	outstruct := new(struct {
		Value     string
		Statement string
		Timestamp *big.Int
		Id        *big.Int
		Version   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Statement = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PublicClaims is a free data retrieval call binding the contract method 0xee10c7e2.
//
// Solidity: function publicClaims(string ) view returns(string value, string statement, uint256 timestamp, int256 id, int256 version)
func (_IdentityManager *IdentityManagerSession) PublicClaims(arg0 string) (struct {
	Value     string
	Statement string
	Timestamp *big.Int
	Id        *big.Int
	Version   *big.Int
}, error) {
	return _IdentityManager.Contract.PublicClaims(&_IdentityManager.CallOpts, arg0)
}

// PublicClaims is a free data retrieval call binding the contract method 0xee10c7e2.
//
// Solidity: function publicClaims(string ) view returns(string value, string statement, uint256 timestamp, int256 id, int256 version)
func (_IdentityManager *IdentityManagerCallerSession) PublicClaims(arg0 string) (struct {
	Value     string
	Statement string
	Timestamp *big.Int
	Id        *big.Int
	Version   *big.Int
}, error) {
	return _IdentityManager.Contract.PublicClaims(&_IdentityManager.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityManager *IdentityManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityManager *IdentityManagerSession) Registry() (common.Address, error) {
	return _IdentityManager.Contract.Registry(&_IdentityManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityManager *IdentityManagerCallerSession) Registry() (common.Address, error) {
	return _IdentityManager.Contract.Registry(&_IdentityManager.CallOpts)
}

// Revocations is a free data retrieval call binding the contract method 0x8107769b.
//
// Solidity: function revocations(string ) view returns(address attestedTo, int256 attestationId, string status, uint256 timestamp)
func (_IdentityManager *IdentityManagerCaller) Revocations(opts *bind.CallOpts, arg0 string) (struct {
	AttestedTo    common.Address
	AttestationId *big.Int
	Status        string
	Timestamp     *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManager.contract.Call(opts, &out, "revocations", arg0)

	outstruct := new(struct {
		AttestedTo    common.Address
		AttestationId *big.Int
		Status        string
		Timestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttestedTo = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AttestationId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Revocations is a free data retrieval call binding the contract method 0x8107769b.
//
// Solidity: function revocations(string ) view returns(address attestedTo, int256 attestationId, string status, uint256 timestamp)
func (_IdentityManager *IdentityManagerSession) Revocations(arg0 string) (struct {
	AttestedTo    common.Address
	AttestationId *big.Int
	Status        string
	Timestamp     *big.Int
}, error) {
	return _IdentityManager.Contract.Revocations(&_IdentityManager.CallOpts, arg0)
}

// Revocations is a free data retrieval call binding the contract method 0x8107769b.
//
// Solidity: function revocations(string ) view returns(address attestedTo, int256 attestationId, string status, uint256 timestamp)
func (_IdentityManager *IdentityManagerCallerSession) Revocations(arg0 string) (struct {
	AttestedTo    common.Address
	AttestationId *big.Int
	Status        string
	Timestamp     *big.Int
}, error) {
	return _IdentityManager.Contract.Revocations(&_IdentityManager.CallOpts, arg0)
}

// DeleteClaim is a paid mutator transaction binding the contract method 0x4665b258.
//
// Solidity: function deleteClaim(string key, bool isPrivate) returns()
func (_IdentityManager *IdentityManagerTransactor) DeleteClaim(opts *bind.TransactOpts, key string, isPrivate bool) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "deleteClaim", key, isPrivate)
}

// DeleteClaim is a paid mutator transaction binding the contract method 0x4665b258.
//
// Solidity: function deleteClaim(string key, bool isPrivate) returns()
func (_IdentityManager *IdentityManagerSession) DeleteClaim(key string, isPrivate bool) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaim(&_IdentityManager.TransactOpts, key, isPrivate)
}

// DeleteClaim is a paid mutator transaction binding the contract method 0x4665b258.
//
// Solidity: function deleteClaim(string key, bool isPrivate) returns()
func (_IdentityManager *IdentityManagerTransactorSession) DeleteClaim(key string, isPrivate bool) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaim(&_IdentityManager.TransactOpts, key, isPrivate)
}

// DeleteClaimURI is a paid mutator transaction binding the contract method 0x84ce394c.
//
// Solidity: function deleteClaimURI(string key) returns()
func (_IdentityManager *IdentityManagerTransactor) DeleteClaimURI(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "deleteClaimURI", key)
}

// DeleteClaimURI is a paid mutator transaction binding the contract method 0x84ce394c.
//
// Solidity: function deleteClaimURI(string key) returns()
func (_IdentityManager *IdentityManagerSession) DeleteClaimURI(key string) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaimURI(&_IdentityManager.TransactOpts, key)
}

// DeleteClaimURI is a paid mutator transaction binding the contract method 0x84ce394c.
//
// Solidity: function deleteClaimURI(string key) returns()
func (_IdentityManager *IdentityManagerTransactorSession) DeleteClaimURI(key string) (*types.Transaction, error) {
	return _IdentityManager.Contract.DeleteClaimURI(&_IdentityManager.TransactOpts, key)
}

// RevokeAttestation is a paid mutator transaction binding the contract method 0xc87c6d41.
//
// Solidity: function revokeAttestation(string key, string reason, address attestedTo, int256 attestationId) returns()
func (_IdentityManager *IdentityManagerTransactor) RevokeAttestation(opts *bind.TransactOpts, key string, reason string, attestedTo common.Address, attestationId *big.Int) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "revokeAttestation", key, reason, attestedTo, attestationId)
}

// RevokeAttestation is a paid mutator transaction binding the contract method 0xc87c6d41.
//
// Solidity: function revokeAttestation(string key, string reason, address attestedTo, int256 attestationId) returns()
func (_IdentityManager *IdentityManagerSession) RevokeAttestation(key string, reason string, attestedTo common.Address, attestationId *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.RevokeAttestation(&_IdentityManager.TransactOpts, key, reason, attestedTo, attestationId)
}

// RevokeAttestation is a paid mutator transaction binding the contract method 0xc87c6d41.
//
// Solidity: function revokeAttestation(string key, string reason, address attestedTo, int256 attestationId) returns()
func (_IdentityManager *IdentityManagerTransactorSession) RevokeAttestation(key string, reason string, attestedTo common.Address, attestationId *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.RevokeAttestation(&_IdentityManager.TransactOpts, key, reason, attestedTo, attestationId)
}

// SetAttestation is a paid mutator transaction binding the contract method 0xdbbd60d9.
//
// Solidity: function setAttestation(string key, address attestor, uint256 expires, bytes signature, int256 claimId, int256 id) returns()
func (_IdentityManager *IdentityManagerTransactor) SetAttestation(opts *bind.TransactOpts, key string, attestor common.Address, expires *big.Int, signature []byte, claimId *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setAttestation", key, attestor, expires, signature, claimId, id)
}

// SetAttestation is a paid mutator transaction binding the contract method 0xdbbd60d9.
//
// Solidity: function setAttestation(string key, address attestor, uint256 expires, bytes signature, int256 claimId, int256 id) returns()
func (_IdentityManager *IdentityManagerSession) SetAttestation(key string, attestor common.Address, expires *big.Int, signature []byte, claimId *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetAttestation(&_IdentityManager.TransactOpts, key, attestor, expires, signature, claimId, id)
}

// SetAttestation is a paid mutator transaction binding the contract method 0xdbbd60d9.
//
// Solidity: function setAttestation(string key, address attestor, uint256 expires, bytes signature, int256 claimId, int256 id) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetAttestation(key string, attestor common.Address, expires *big.Int, signature []byte, claimId *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetAttestation(&_IdentityManager.TransactOpts, key, attestor, expires, signature, claimId, id)
}

// SetClaimURI is a paid mutator transaction binding the contract method 0xc03bc8b5.
//
// Solidity: function setClaimURI(string key, string value) returns()
func (_IdentityManager *IdentityManagerTransactor) SetClaimURI(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setClaimURI", key, value)
}

// SetClaimURI is a paid mutator transaction binding the contract method 0xc03bc8b5.
//
// Solidity: function setClaimURI(string key, string value) returns()
func (_IdentityManager *IdentityManagerSession) SetClaimURI(key string, value string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetClaimURI(&_IdentityManager.TransactOpts, key, value)
}

// SetClaimURI is a paid mutator transaction binding the contract method 0xc03bc8b5.
//
// Solidity: function setClaimURI(string key, string value) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetClaimURI(key string, value string) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetClaimURI(&_IdentityManager.TransactOpts, key, value)
}

// SetPrivateClaim is a paid mutator transaction binding the contract method 0xead05a34.
//
// Solidity: function setPrivateClaim(string key, string value, string statement, string ipfsURI, string eventHash, uint256 timestamp, int256 identifier) returns()
func (_IdentityManager *IdentityManagerTransactor) SetPrivateClaim(opts *bind.TransactOpts, key string, value string, statement string, ipfsURI string, eventHash string, timestamp *big.Int, identifier *big.Int) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setPrivateClaim", key, value, statement, ipfsURI, eventHash, timestamp, identifier)
}

// SetPrivateClaim is a paid mutator transaction binding the contract method 0xead05a34.
//
// Solidity: function setPrivateClaim(string key, string value, string statement, string ipfsURI, string eventHash, uint256 timestamp, int256 identifier) returns()
func (_IdentityManager *IdentityManagerSession) SetPrivateClaim(key string, value string, statement string, ipfsURI string, eventHash string, timestamp *big.Int, identifier *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPrivateClaim(&_IdentityManager.TransactOpts, key, value, statement, ipfsURI, eventHash, timestamp, identifier)
}

// SetPrivateClaim is a paid mutator transaction binding the contract method 0xead05a34.
//
// Solidity: function setPrivateClaim(string key, string value, string statement, string ipfsURI, string eventHash, uint256 timestamp, int256 identifier) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetPrivateClaim(key string, value string, statement string, ipfsURI string, eventHash string, timestamp *big.Int, identifier *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPrivateClaim(&_IdentityManager.TransactOpts, key, value, statement, ipfsURI, eventHash, timestamp, identifier)
}

// SetPublicClaim is a paid mutator transaction binding the contract method 0x5d91ca9f.
//
// Solidity: function setPublicClaim(string key, string value, string statement, uint256 timestamp, int256 id) returns()
func (_IdentityManager *IdentityManagerTransactor) SetPublicClaim(opts *bind.TransactOpts, key string, value string, statement string, timestamp *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "setPublicClaim", key, value, statement, timestamp, id)
}

// SetPublicClaim is a paid mutator transaction binding the contract method 0x5d91ca9f.
//
// Solidity: function setPublicClaim(string key, string value, string statement, uint256 timestamp, int256 id) returns()
func (_IdentityManager *IdentityManagerSession) SetPublicClaim(key string, value string, statement string, timestamp *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPublicClaim(&_IdentityManager.TransactOpts, key, value, statement, timestamp, id)
}

// SetPublicClaim is a paid mutator transaction binding the contract method 0x5d91ca9f.
//
// Solidity: function setPublicClaim(string key, string value, string statement, uint256 timestamp, int256 id) returns()
func (_IdentityManager *IdentityManagerTransactorSession) SetPublicClaim(key string, value string, statement string, timestamp *big.Int, id *big.Int) (*types.Transaction, error) {
	return _IdentityManager.Contract.SetPublicClaim(&_IdentityManager.TransactOpts, key, value, statement, timestamp, id)
}

// IdentityManagerAttestationRevokedIterator is returned from FilterAttestationRevoked and is used to iterate over the raw logs and unpacked data for AttestationRevoked events raised by the IdentityManager contract.
type IdentityManagerAttestationRevokedIterator struct {
	Event *IdentityManagerAttestationRevoked // Event containing the contract specifics and raw log

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
func (it *IdentityManagerAttestationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerAttestationRevoked)
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
		it.Event = new(IdentityManagerAttestationRevoked)
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
func (it *IdentityManagerAttestationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerAttestationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerAttestationRevoked represents a AttestationRevoked event raised by the IdentityManager contract.
type IdentityManagerAttestationRevoked struct {
	Identity common.Address
	Id       *big.Int
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttestationRevoked is a free log retrieval operation binding the contract event 0x613a02027bb8290f4136e1ca1f3f2d877cf32d609bbc9cfa4799ad69b8e2a08b.
//
// Solidity: event AttestationRevoked(address indexed identity, int256 id, string reason)
func (_IdentityManager *IdentityManagerFilterer) FilterAttestationRevoked(opts *bind.FilterOpts, identity []common.Address) (*IdentityManagerAttestationRevokedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "AttestationRevoked", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerAttestationRevokedIterator{contract: _IdentityManager.contract, event: "AttestationRevoked", logs: logs, sub: sub}, nil
}

// WatchAttestationRevoked is a free log subscription operation binding the contract event 0x613a02027bb8290f4136e1ca1f3f2d877cf32d609bbc9cfa4799ad69b8e2a08b.
//
// Solidity: event AttestationRevoked(address indexed identity, int256 id, string reason)
func (_IdentityManager *IdentityManagerFilterer) WatchAttestationRevoked(opts *bind.WatchOpts, sink chan<- *IdentityManagerAttestationRevoked, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "AttestationRevoked", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerAttestationRevoked)
				if err := _IdentityManager.contract.UnpackLog(event, "AttestationRevoked", log); err != nil {
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

// ParseAttestationRevoked is a log parse operation binding the contract event 0x613a02027bb8290f4136e1ca1f3f2d877cf32d609bbc9cfa4799ad69b8e2a08b.
//
// Solidity: event AttestationRevoked(address indexed identity, int256 id, string reason)
func (_IdentityManager *IdentityManagerFilterer) ParseAttestationRevoked(log types.Log) (*IdentityManagerAttestationRevoked, error) {
	event := new(IdentityManagerAttestationRevoked)
	if err := _IdentityManager.contract.UnpackLog(event, "AttestationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityManagerNewAttestationIterator is returned from FilterNewAttestation and is used to iterate over the raw logs and unpacked data for NewAttestation events raised by the IdentityManager contract.
type IdentityManagerNewAttestationIterator struct {
	Event *IdentityManagerNewAttestation // Event containing the contract specifics and raw log

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
func (it *IdentityManagerNewAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerNewAttestation)
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
		it.Event = new(IdentityManagerNewAttestation)
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
func (it *IdentityManagerNewAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerNewAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerNewAttestation represents a NewAttestation event raised by the IdentityManager contract.
type IdentityManagerNewAttestation struct {
	Identity  common.Address
	Key       string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewAttestation is a free log retrieval operation binding the contract event 0x9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a.
//
// Solidity: event NewAttestation(address indexed identity, string key, uint256 timestamp)
func (_IdentityManager *IdentityManagerFilterer) FilterNewAttestation(opts *bind.FilterOpts, identity []common.Address) (*IdentityManagerNewAttestationIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "NewAttestation", identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerNewAttestationIterator{contract: _IdentityManager.contract, event: "NewAttestation", logs: logs, sub: sub}, nil
}

// WatchNewAttestation is a free log subscription operation binding the contract event 0x9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a.
//
// Solidity: event NewAttestation(address indexed identity, string key, uint256 timestamp)
func (_IdentityManager *IdentityManagerFilterer) WatchNewAttestation(opts *bind.WatchOpts, sink chan<- *IdentityManagerNewAttestation, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "NewAttestation", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerNewAttestation)
				if err := _IdentityManager.contract.UnpackLog(event, "NewAttestation", log); err != nil {
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

// ParseNewAttestation is a log parse operation binding the contract event 0x9476865afb170ca0b75c734e92ad4be2fa603b741a4c4542361fe02f8ff9ae4a.
//
// Solidity: event NewAttestation(address indexed identity, string key, uint256 timestamp)
func (_IdentityManager *IdentityManagerFilterer) ParseNewAttestation(log types.Log) (*IdentityManagerNewAttestation, error) {
	event := new(IdentityManagerNewAttestation)
	if err := _IdentityManager.contract.UnpackLog(event, "NewAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityManagerNewClaimIterator is returned from FilterNewClaim and is used to iterate over the raw logs and unpacked data for NewClaim events raised by the IdentityManager contract.
type IdentityManagerNewClaimIterator struct {
	Event *IdentityManagerNewClaim // Event containing the contract specifics and raw log

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
func (it *IdentityManagerNewClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagerNewClaim)
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
		it.Event = new(IdentityManagerNewClaim)
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
func (it *IdentityManagerNewClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagerNewClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagerNewClaim represents a NewClaim event raised by the IdentityManager contract.
type IdentityManagerNewClaim struct {
	Key       string
	Value     string
	IsPrivate bool
	Version   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewClaim is a free log retrieval operation binding the contract event 0x26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62.
//
// Solidity: event NewClaim(string key, string value, bool isPrivate, int256 version)
func (_IdentityManager *IdentityManagerFilterer) FilterNewClaim(opts *bind.FilterOpts) (*IdentityManagerNewClaimIterator, error) {

	logs, sub, err := _IdentityManager.contract.FilterLogs(opts, "NewClaim")
	if err != nil {
		return nil, err
	}
	return &IdentityManagerNewClaimIterator{contract: _IdentityManager.contract, event: "NewClaim", logs: logs, sub: sub}, nil
}

// WatchNewClaim is a free log subscription operation binding the contract event 0x26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62.
//
// Solidity: event NewClaim(string key, string value, bool isPrivate, int256 version)
func (_IdentityManager *IdentityManagerFilterer) WatchNewClaim(opts *bind.WatchOpts, sink chan<- *IdentityManagerNewClaim) (event.Subscription, error) {

	logs, sub, err := _IdentityManager.contract.WatchLogs(opts, "NewClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagerNewClaim)
				if err := _IdentityManager.contract.UnpackLog(event, "NewClaim", log); err != nil {
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

// ParseNewClaim is a log parse operation binding the contract event 0x26f9b425c95c58e8bec57a1841c65b09f8b1b4408004758fcebf088adc4b3d62.
//
// Solidity: event NewClaim(string key, string value, bool isPrivate, int256 version)
func (_IdentityManager *IdentityManagerFilterer) ParseNewClaim(log types.Log) (*IdentityManagerNewClaim, error) {
	event := new(IdentityManagerNewClaim)
	if err := _IdentityManager.contract.UnpackLog(event, "NewClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
