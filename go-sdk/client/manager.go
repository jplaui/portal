package client

import (
	"fmt"
	"go-sdk/contracts/IdentityManager"
	"log"
	"math/big"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ManagerClient struct {
	Manager    *IdentityManager.IdentityManager
	IpfsClient *IpfsClient
}

type ClaimMeta struct {
	Value     string  `json:"value"`
	Statement string  `json:"statement"`
	Timestamp big.Int `json:"timestamp"`
	Id        int     `json:"id"`
	Version   int     `json:"version"`
}

func (s *ClaimMeta) Size() int {
	size := int(unsafe.Sizeof(*s))
	// size += len(s.F)
	return size
}

type PublicClaim struct {
	ClaimMeta ClaimMeta
}

func NewPublicClaim(claimMeta ClaimMeta) *PublicClaim {
	return &PublicClaim{ClaimMeta: claimMeta}
}

type PrivateClaim struct {
	ClaimMeta ClaimMeta
	IpfsURI   string
	EventHash string
}

func NewPrivateClaim(claimMeta ClaimMeta, ipfsURI string, eventHash string) *PrivateClaim {
	return &PrivateClaim{ClaimMeta: claimMeta, IpfsURI: ipfsURI, EventHash: eventHash}
}

type Attestation struct {
	Signature []byte
	Attestor  string
	Timestamp big.Int
	ExpiresAt big.Int
	Id        int
	ClaimId   int
}

func (s *Attestation) Size() int {
	size := int(unsafe.Sizeof(*s))
	// size += len(s.F)
	return size
}
func (s *Revocation) Size() int {
	size := int(unsafe.Sizeof(*s))
	// size += len(s.F)
	return size
}

type Revocation struct {
	AttestedTo    common.Address
	AttestationId *big.Int
	Status        string
	Timestamp     *big.Int
}

func NewAttestation(signature []byte, attestor string, timestamp big.Int, expiresAt big.Int, id int, claimId int) *Attestation {
	return &Attestation{Signature: signature, Attestor: attestor, Timestamp: timestamp, ExpiresAt: expiresAt, Id: id, ClaimId: claimId}
}

func NewManagerClient(contractAddr string, reader TxReader, ipfsConn string) *ManagerClient {
	contract := common.HexToAddress(contractAddr)
	instance, err := IdentityManager.NewIdentityManager(contract, reader.Client)
	if err != nil {
		log.Fatal(err)
	}

	ipfsClient := NewIpfsClient(ipfsConn)
	return &ManagerClient{
		Manager:    instance,
		IpfsClient: ipfsClient,
	}
}

var gasPrice = big.NewInt(2800000000)

func WeiToEther(wei uint64) *big.Float {
	weiInEther := new(big.Float).SetUint64(wei)
	ether := new(big.Float).Quo(weiInEther, big.NewFloat(1e18))
	return ether
}
func CalculateGasUsage(operation string, receipt *types.Receipt) {
	fmt.Printf("%s Transaction Cost (Wei): %d \n", operation, receipt.GasUsed)
	fmt.Printf("%s Transaction Cost (Ether): %s\n", operation, WeiToEther(receipt.GasUsed).String())
	total_cost := new(big.Int)
	total_gas := new(big.Int).SetUint64(receipt.GasUsed)
	total_cost.Mul(total_gas, gasPrice)
	fmt.Printf("Total Cost (Ether): %s\n", WeiToEther(total_cost.Uint64()).String())
}

func SetPublicClaim(client *ManagerClient, signer *Signer, key string, value string, statement string) (*types.Receipt, error) {
	auth := signer.BindTxOpts()
	// convert time now to unix timestamp

	timestamp := time.Now().Unix()
	timeInBig := new(big.Int).SetInt64(timestamp)
	id := timeInBig
	t1 := time.Now()
	tx, err := client.Manager.SetPublicClaim(auth, key, value, statement, timeInBig, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("setPublicClaim tx size:", tx.Size())

	receipt, err := signer.WaitForReceipt(tx.Hash())
	if err != nil {
		return nil, err
	}

	fmt.Println("setting Public Claim took:", time.Since(t1))
	CalculateGasUsage("Plublic claim deployment", receipt)

	return receipt, nil
}

func GetPublicClaim(client *ManagerClient, key string) (*PublicClaim, error) {
	t1 := time.Now()
	claim, err := client.Manager.PublicClaims(&bind.CallOpts{}, key)
	if err != nil {
		return nil, err
	}
	fmt.Println("gettting public claim took:", time.Since(t1))
	return NewPublicClaim(ClaimMeta{
		Value:     claim.Value,
		Statement: claim.Statement,
		Timestamp: *claim.Timestamp,
		Id:        int(claim.Id.Int64()),
	}), nil
}

func SetPrivateClaim(client *ManagerClient, signer *Signer, key string, value string, statement string, ipfsURI string,
	eventHash string) (*types.Receipt, error) {
	t1 := time.Now()
	auth := signer.BindTxOpts()
	timestamp := time.Now().Unix()
	timeInBig := new(big.Int).SetInt64(timestamp)
	id := timeInBig
	tx, err := client.Manager.SetPrivateClaim(auth, key, value, statement, ipfsURI, eventHash, timeInBig, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("setPrivateClaim tx size:", tx.Size())

	receipt, err := signer.WaitForReceipt(tx.Hash())
	if err != nil {
		return nil, err
	}
	fmt.Println("getting private claim took:", time.Since(t1))
	return receipt, nil
}
func GetPrivateClaim(client *ManagerClient, key string) (*PrivateClaim, error) {
	claim, err := client.Manager.PrivateClaims(&bind.CallOpts{}, key)
	if err != nil {
		return nil, err
	}
	return NewPrivateClaim(ClaimMeta{
		Value:     claim.Value,
		Statement: claim.Statement,
		Timestamp: *claim.Timestamp,
		Id:        int(claim.Id.Int64()),
		Version:   int(claim.Version.Int64()),
	}, claim.IpfsCircuitMetadata, claim.EventHash), nil
}

//func DeleteClaim(client *ManagerClient, signer *Signer, key string) (*types.Receipt, error) {
//	auth := signer.BindTxOpts()
//
//	tx, err := client.Manager.DeleteClaim(auth, key)
//	if err != nil {
//		return nil, err
//	}
//
//	receipt, err := signer.WaitForReceipt( tx.Hash())
//	if err != nil {
//		return nil, err
//	}
//
//	return receipt, nil
//}

func GetClaimURI(client *ManagerClient, claimHash string) (string, error) {
	t1 := time.Now()
	val, err := client.Manager.IpfsClaims(nil, claimHash)
	if err != nil {
		return "", err
	}

	fmt.Println("getting claim from IPFS took;", time.Since(t1))
	return val, nil
}

func setClaimURI(client *ManagerClient, signer *Signer, claimHash string, uri string) (*types.Receipt, error) {
	auth := signer.BindTxOpts()

	tx, err := client.Manager.SetClaimURI(auth, claimHash, uri)
	if err != nil {
		return nil, err
	}

	receipt, err := signer.WaitForReceipt(tx.Hash())
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func PublishClaim(client *ManagerClient, signer *Signer, key string, value []byte) (*types.Receipt, error) {
	//write claim first on ipfs
	t1 := time.Now()
	resp, err := client.IpfsClient.AddAndPublish(value)
	if err != nil {
		return nil, err
	}
	ipnsPath := resp["Name"]
	fmt.Println("setting claim on IPFS took", time.Since(t1))
	return setClaimURI(client, signer, key, ipnsPath)
}

func SetAttestation(client *ManagerClient, signer *Signer, key string, attestor common.Address,
	signature []byte, expires big.Int, claimId big.Int) (*types.Receipt, error) {
	auth := signer.BindTxOpts()
	tx, err := client.Manager.SetAttestation(auth, key, attestor, &expires, signature, &claimId, &claimId)
	if err != nil {
		return nil, err
	}
	fmt.Println("attetstation tx size:", tx.Size())

	receipt, err := signer.WaitForReceipt(tx.Hash())
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func GetAttestation(client *ManagerClient, key string) (*Attestation, error) {
	attestation, err := client.Manager.Attestations(&bind.CallOpts{}, key)
	if err != nil {
		return nil, err
	}

	return NewAttestation(
		attestation.Signature,
		attestation.Attestor.Hex(),
		*attestation.Timestamp,
		*attestation.Expires,
		int(attestation.Id.Int64()),
		int(attestation.ClaimId.Int64())), nil
}

func Revoke(client *ManagerClient, signer *Signer, key string, attestedTo common.Address,
	id *big.Int, reason string) (*types.Receipt, error) {
	auth := signer.BindTxOpts()
	tx, err := client.Manager.RevokeAttestation(auth, key, reason, attestedTo, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("revocation tx size:", tx.Size())

	receipt, err := signer.WaitForReceipt(tx.Hash())
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func GetRevocation(client *ManagerClient, key string) (Revocation, error) {
	revocation, err := client.Manager.Revocations(&bind.CallOpts{}, key)
	if err != nil {
		return Revocation{}, err
	}

	return revocation, nil
}
