package client

import (
	"bytes"
	"fmt"
	"go-sdk/contracts/IdentityRegistry"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type RegistryClient struct {
	Registry *IdentityRegistry.IdentityRegistry
}

type CircuitMeta struct {
	DeploymentType string
	Address        common.Address
	IpfsURI        string
}

func NewRegistryClient(contractAddr string, reader TxReader) *RegistryClient {

	contract := common.HexToAddress(contractAddr)
	instance, err := IdentityRegistry.NewIdentityRegistry(contract, reader.Client)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &RegistryClient{
		Registry: instance,
	}
}

func RegisterIdentity(client *RegistryClient, sender *Signer, owner *Signer, managerAddr common.Address) (*types.Receipt, error) {
	t1 := time.Now()
	auth := sender.BindTxOpts()
	senderAddr := sender.PublicKey
	// sign message by contract owner for integrity and authenticity check for registering a new identity
	_, signature, err := CalculateSignature(senderAddr, owner, managerAddr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//log.Printf("signature sending to contract: 0x%s\n", common.Bytes2Hex(signature))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tx, err := client.Registry.Register(auth, managerAddr, signature)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	receipt, err := sender.WaitForReceipt(tx.Hash())
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	fmt.Println("register client TX SIZE:", tx.Size())
	fmt.Println("register new client took", time.Since(t1))
	CalculateGasUsage("register user at registry", receipt)
	return receipt, nil
}
func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}
func CalculateSignature(senderAddr []byte, owner *Signer, managerAddr common.Address) ([]byte, []byte, error) {
	signDataInBytes := encodePacked(managerAddr.Bytes(), senderAddr)
	hash := crypto.Keccak256(signDataInBytes)

	log.Printf("MessageHash 0x%s\n", common.Bytes2Hex(hash))

	var buf bytes.Buffer
	buf.Write([]byte("\x19Ethereum Signed Message:\n32")) // For byte slice
	buf.Write(hash)

	finalHash := crypto.Keccak256Hash(buf.Bytes())
	fBytes := finalHash.Bytes()

	signature, err := owner.Sign(fBytes)
	signature[64] = signature[64] + 27

	// get v, r, s from signature
	//r := signature[:32]
	//s := signature[32:64]
	//v := signature[64:65]
	// verify signature locally
	//verified := crypto.VerifySignature(crypto.CompressPubkey(owner.PublicKey), fBytes, signature[:64])
	//fmt.Println("verified", verified)
	return fBytes, signature, err
}

func GetIdentity(client *RegistryClient, address common.Address) (common.Address, error) {
	read, err := client.Registry.OwnerOf(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return read, nil
}

func GetManager(client *RegistryClient, address common.Address) (common.Address, error) {

	read, err := client.Registry.Manager(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return read, nil
}

func RegisterCircuit(client *RegistryClient, signer *Signer, circuitId string, c *CircuitMeta) (*types.Receipt, error) {
	auth := signer.BindTxOpts()
	tx, err := client.Registry.SetCircuit(auth, circuitId, c.DeploymentType, c.Address, c.IpfsURI)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("reigster circuit TX SIZE:", tx.Size())
	receipt, err := signer.WaitForReceipt(tx.Hash())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return receipt, nil
}

func GetCircuit(client *RegistryClient, circuitId string) (*CircuitMeta, error) {
	read, err := client.Registry.Circuits(nil, circuitId)
	if err != nil {
		log.Fatal(err)
	}
	return &CircuitMeta{DeploymentType: read.DeploymentType, Address: read.DeploymentAddress, IpfsURI: read.IpfsHash}, nil
}
