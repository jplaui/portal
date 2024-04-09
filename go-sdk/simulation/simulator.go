package simulation

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"go-sdk/circuits"
	"go-sdk/circuits/merklecombined"
	"go-sdk/circuits/mimc"
	"go-sdk/client"
	"go-sdk/deployment"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var totalGasUsage uint64 = 0
var gasPrice = big.NewInt(12000000000) // // in wei (12 gwei), used to be 2000000000
var solidityPath = "circuits/contracts/%s.sol"

const totalClaim = 100

func SimulateIpfsOnChainPublicClaim(registryAddr string, userAccount common.Address, userSigner *client.Signer,
	reader *client.TxReader, ipfsConn string) {

	fmt.Println("Simulation type: ipfs-on-chain")
	regClient := client.NewRegistryClient(registryAddr, *reader)
	manager, err := client.GetManager(regClient, userAccount)
	if err != nil {
		fmt.Println("Error getting manager:", err)
		return
	}
	log.Printf("Manager address for user %s: %s\n", userAccount.Hex(), manager)

	manClient := client.NewManagerClient(manager.Hex(), *reader, ipfsConn)
	// start timer
	startTime := time.Now()
	for i := 0; i < totalClaim; i++ {
		jsonStr := GenerateDummyClaim()
		key := strconv.Itoa(i)
		rec, err := client.PublishClaim(manClient, userSigner, key, jsonStr)
		totalGasUsage += rec.GasUsed

		if err != nil {
			log.Fatal("Could not set claim", err)
			return
		}
	}

	// Calculate the execution time.
	elapsedTime := time.Since(startTime)

	fmt.Printf("Number of %d claims to set took %s\n.Average per claim:%s \n", totalClaim, elapsedTime, elapsedTime/totalClaim)

	fmt.Printf("Transaction Cost (Wei): %d \n", totalGasUsage)
	fmt.Printf("Transaction Cost (Ether): %s\n", weiToEther(totalGasUsage).String())
	total_cost := new(big.Int)
	total_gas := new(big.Int).SetUint64(totalGasUsage)
	total_cost.Mul(total_gas, gasPrice)
	fmt.Printf("Total Cost (Ether): %s\n", weiToEther(total_cost.Uint64()).String())
}

func SimulateOnChainPublicClaims(registryAddr string, userAccount common.Address, userSigner *client.Signer,
	reader *client.TxReader) {
	fmt.Println("Simulation type: onchain")
	regClient := client.NewRegistryClient(registryAddr, *reader)
	// check identity is registered
	manager, err := client.GetManager(regClient, userAccount)
	if err != nil {
		fmt.Println("Error getting manager:", err)
		return
	}
	log.Printf("Manager address for user %s: %s\n", userAccount.Hex(), manager)

	// now user can set public claim by it's own manager
	manClient := client.NewManagerClient(manager.Hex(), *reader, "")
	// start timer
	startTime := time.Now()
	for i := 0; i < totalClaim; i++ {
		jsonStr := GenerateDummyClaim()
		key := strconv.Itoa(i)
		rec, err := client.SetPublicClaim(manClient, userSigner, key, string(jsonStr), "age > 22")
		totalGasUsage += rec.GasUsed

		if err != nil {
			log.Fatal("Could not set claim", err)
			return
		}

	}

	// Calculate the execution time.
	elapsedTime := time.Since(startTime)

	fmt.Printf("Number of %d claims to set took %s\n.Average per claim:%s \n", totalClaim, elapsedTime, elapsedTime/totalClaim)

	fmt.Printf("Transaction Cost (Wei): %d \n", totalGasUsage)
	fmt.Printf("Transaction Cost (Ether): %s\n", weiToEther(totalGasUsage).String())
	total_cost := new(big.Int)
	total_gas := new(big.Int).SetUint64(totalGasUsage)
	total_cost.Mul(total_gas, gasPrice)
	fmt.Printf("Total Cost (Ether): %s\n", weiToEther(total_cost.Uint64()).String())
}

/*
Tests setting public claim on ipfs only.
No contract interaction is involved.
*/
func SimulateIPFSConnection(ipfsConn string) {
	fmt.Println("Simulation type: ipfs")

	cli := client.NewIpfsClient(ipfsConn)
	// start timer
	startTime := time.Now()
	for i := 0; i < totalClaim; i++ {
		jsonStr := GenerateDummyClaim()

		_, err := cli.AddAndPublish(jsonStr)

		if err != nil {
			log.Fatal("Could not set claim", err)
			return
		}

	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Number of %d claims to set took %s\n.Average per claim:%s \n",
		totalClaim, elapsedTime, elapsedTime/totalClaim)
}

func SimulateContractDeployment(userSigner *client.Signer, ownerSigner *client.Signer) {

	fmt.Println("Simulation type: deployment")

	// ownerAccount deploys the registry
	startTime := time.Now()
	registryContract, tx := deployment.DeployRegistry(ownerSigner)
	elapsedTime := time.Since(startTime)

	fmt.Println("Registry contract deployed at:", registryContract)
	fmt.Printf("Deployment time: %s\n", elapsedTime)
	fmt.Println("registry contract deploy TX SIZE:", tx.Size())

	receipt, err := ownerSigner.WaitForReceipt(tx.Hash())
	if err != nil {
		log.Fatal("Couldn't get transaction receipt", err)
	}
	calculateGasUsage("Registry deployment", receipt)

	bytecode, err := userSigner.Client.CodeAt(context.Background(), registryContract, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytecode length for registry contract: %d\n", len(bytecode))

	// start timer
	startTime = time.Now()
	// userAccount deploys it's own manager
	managerAddr, tx := deployment.DeployManager(userSigner, registryContract)
	elapsedTime = time.Since(startTime)

	fmt.Println("identity contract deploy TX SIZE:", tx.Size())

	receipt, err = userSigner.WaitForReceipt(tx.Hash())
	if err != nil {
		log.Fatal("Couldn't get transaction receipt", err)
	}
	fmt.Println("Manager contract deployed at:", managerAddr)
	fmt.Printf("Deployment time: %s\n", elapsedTime)
	calculateGasUsage("Manager deployment", receipt)

	bytecode, err = userSigner.Client.CodeAt(context.Background(), managerAddr, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytecode length for identity contract: %d\n", len(bytecode))
}

/*
- circuit build
- circuit metadata generation
- ccs, vk,pk,proof serialization to ipfs
- circuit metadata registration on chain (registry)
- circuit metadata retrieval from chain
- circuit metadata retrieval from ipfs
- circuit metadata deserialization
- circuit verification
- circuit verification time
*/
func SimulateMerkleCircuitDeployment(userSigner *client.Signer, ipfsConn string, registryAddr string, reader *client.TxReader) {
	ipfsClient := client.NewIpfsClient(ipfsConn)
	// first build the circuit
	fmt.Println("Simulation type: circuit build")
	curve := ecc.BN254

	var (
		pre       = 100
		threshold = 22
		digest    = "20370067689261511688289967978544823130432235585709842144916192060767982363628"
	)
	numLeaves := 32
	segmentSize := 32
	mod := curve.ScalarField()
	modNbBytes := len(mod.Bytes())
	depth := 5

	startTime := time.Now()
	var merkleCircuit merklecombined.MerkleCombined
	merkleCircuit.Path = make([]frontend.Variable, depth+1)

	cc, err := frontend.Compile(curve.ScalarField(), scs.NewBuilder, &merkleCircuit)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	mimcGo := hash.MIMC_BN254.New()
	// convert pre to bytes with padding to 32 bytes
	inpBytes := make([]byte, 32)
	preBytes := []byte(strconv.Itoa(pre))
	copy(inpBytes[32-len(preBytes):], preBytes)

	mimcGo.Write(inpBytes)
	sum := mimcGo.Sum(nil)

	// normalize the input with mod BN254
	inputMod := new(big.Int).SetBytes(sum)
	inputMod.Mod(inputMod, curve.BaseField())

	sumModBytes := inputMod.Bytes()
	buf.Write(make([]byte, modNbBytes-len(sumModBytes)))
	buf.Write(sumModBytes)

	//// some random elements for tree
	for i := 1; i < numLeaves; i++ {
		leaf, _ := crand.Int(crand.Reader, mod)

		b := leaf.Bytes()

		buf.Write(make([]byte, modNbBytes-len(b)))
		buf.Write(b)
	}

	hGo := hash.MIMC_BN254.New()

	//merkleRoot, proofPath, index, leaves := tree.Prove()
	merkleRoot, proofPath, _, err := merkletree.BuildReaderProof(&buf, hGo, segmentSize, uint64(0))

	if err != nil {
		log.Fatal(err)
	}

	//// verify the proof in plain go (data part)
	//verified := merkletree.VerifyProof(hGo, merkleRoot, proofPath, 0, leaves)
	//if !verified {
	//	log.Print("The merkle proof in plain go should pass", err)
	//}

	// wt
	var merkleWt merklecombined.MerkleCombined
	merkleWt.Input = pre
	merkleWt.Threshold = threshold
	merkleWt.Digest = digest
	merkleWt.RootHash = merkleRoot
	merkleWt.Nonce = 0
	merkleWt.NoncePriv = 0
	merkleWt.Sender = 0
	merkleWt.SenderPriv = 0
	merkleWt.Leaf = 0
	merkleWt.Path = make([]frontend.Variable, depth+1)
	for i := 0; i < depth+1; i++ {
		merkleWt.Path[i] = proofPath[i]
	}

	w, err := frontend.NewWitness(&merkleWt, curve.ScalarField())
	if err != nil {
		log.Fatal("Couldn't set witness ", err)
	}

	srs, err := test.NewKZGSRS(cc)
	if err != nil {
		log.Fatal("test.NewKZGSRS(ccs)")
	}
	// groth16 zkSNARK: Setup
	t10 := time.Now()
	pk, vk, err := plonk.Setup(cc, srs)
	if err != nil {
		log.Fatal("plonk.Setup", err)
	}
	fmt.Println("plonk Setup took:", time.Since(t10))
	//// groth16: Prove & Verify
	t11 := time.Now()
	proof, err := plonk.Prove(cc, pk, w)
	if err != nil {
		log.Fatal("prove computation failed...", err)
	}
	fmt.Println("plonk Prove took:", time.Since(t11))

	publicWit, _ := w.Public()
	t12 := time.Now()
	err = plonk.Verify(proof, vk, publicWit)
	if err != nil {
		log.Fatal("Verify failed from local verifier", err)
		return
	}
	fmt.Println("plonk Verify took:", time.Since(t12))

	startTime = time.Now()
	circuitAddr, tx := deployment.DeployCircuit(userSigner)
	elapsedTime := time.Since(startTime)
	fmt.Printf("CircuitMeta deployed at: %s\n", circuitAddr)
	fmt.Printf("On-chain Deployment time: %s\n", elapsedTime)

	serialized := circuits.Serialize(cc, w, vk, pk, proof, "MerkleCombined", circuitAddr.Hex())
	t2 := time.Now()
	ipfs, err := circuits.PublishCircuitOnIpfs(ipfsClient, &serialized)
	if err != nil {
		log.Fatal("Could not publish circuit on ipfs", err)
		return
	}
	log.Printf("CircuitMeta published on ipfs: %s\n", ipfs)
	fmt.Println("ipfs upload of circuit meta took:", time.Since(t2))
	// elapsedTime = time.Since(startTime)
	// fmt.Printf("CircuitMeta build & publish time: %s\n", elapsedTime)
	// fmt.Println("Simulation type: deployment of circuit")

	// get transaction receipt
	receipt, err := userSigner.WaitForReceipt(tx.Hash())
	if err != nil {
		log.Fatal("Couldn't get transaction receipt", err)
	}
	calculateGasUsage("CircuitMeta deployment", receipt)
	// get total gas used

	bytecode, err := userSigner.Client.CodeAt(context.Background(), circuitAddr, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytecode len: %d\n", len(bytecode))

	//register circuit metadata on registry
	regClient := client.NewRegistryClient(registryAddr, *reader)
	startTime = time.Now()
	cMeta := client.CircuitMeta{
		DeploymentType: "ipfs",
		Address:        circuitAddr,
		IpfsURI:        ipfs,
	}
	setReceipt, err := client.RegisterCircuit(regClient, userSigner, "MerkleCombined", &cMeta)
	if err != nil {
		log.Fatal("Could not register circuit metadata on chain", err)
		return
	}
	elapsedTime = time.Since(startTime)
	// get total gas used
	calculateGasUsage("SetCircuit in registry", setReceipt)
	fmt.Printf("SetCircuit in registry time: %s\n", elapsedTime)
}
func SimulateMimcCircuitDeployment(userSigner *client.Signer, ipfsConn string, registryAddr string, reader *client.TxReader) {
	circuitName := "MimcComp"
	ipfsClient := client.NewIpfsClient(ipfsConn)
	// first build the circuit
	fmt.Println("Simulation type: circuit build")
	curve := ecc.BN254

	var (
		pre       = 100
		threshold = 22
		digest    = "20370067689261511688289967978544823130432235585709842144916192060767982363628"
	)

	startTime := time.Now()
	var mimcCircuit mimc.MimcComp

	cc, err := frontend.Compile(curve.ScalarField(), scs.NewBuilder, &mimcCircuit)
	if err != nil {
		log.Fatal(err)
	}

	//wt
	var mimcWt mimc.MimcComp
	mimcWt.Input = pre
	mimcWt.Threshold = threshold
	mimcWt.Digest = digest
	mimcWt.Nonce = 0
	mimcWt.NoncePriv = 0
	mimcWt.Sender = 0
	mimcWt.SenderPriv = 0

	w, err := frontend.NewWitness(&mimcWt, curve.ScalarField())
	if err != nil {
		log.Fatal("Couldn't set witness ", err)
	}

	srs, err := test.NewKZGSRS(cc)
	if err != nil {
		log.Fatal("test.NewKZGSRS(ccs)")
	}
	// groth16 zkSNARK: Setup
	pk, vk, err := plonk.Setup(cc, srs)
	if err != nil {
		log.Fatal("plonk.Setup", err)
	}
	path := fmt.Sprintf(solidityPath, circuitName)
	deployment.CircuitToSolidity(vk, path)
	//// groth16: Prove & Verify
	proof, err := plonk.Prove(cc, pk, w)
	if err != nil {
		log.Fatal("prove computation failed...", err)
	}

	startTime = time.Now()
	circuitAddr, tx := deployment.DeployCircuit(userSigner)
	elapsedTime := time.Since(startTime)

	fmt.Printf("CircuitMeta deployed at: %s\n", circuitAddr)
	serialized := circuits.Serialize(cc, w, vk, pk, proof, circuitName, circuitAddr.Hex())
	ipfs, err := circuits.PublishCircuitOnIpfs(ipfsClient, &serialized)
	if err != nil {
		log.Fatal("Could not publish circuit on ipfs", err)
		return
	}
	log.Printf("CircuitMeta published on ipfs: %s\n", ipfs)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta build & publish time: %s\n", elapsedTime)
	fmt.Println("Simulation type: deployment of circuit")

	// get transaction receipt
	receipt, err := userSigner.WaitForReceipt(tx.Hash())

	if err != nil {
		log.Fatal("Couldn't get transaction receipt", err)
	}
	calculateGasUsage("CircuitMeta deployment", receipt)

	bytecode, err := userSigner.Client.CodeAt(context.Background(), circuitAddr, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytecode len: %d\n", len(bytecode))

	//register circuit metadata on registry
	startTime = time.Now()

	regClient := client.NewRegistryClient(registryAddr, *reader)
	cMeta := client.CircuitMeta{
		DeploymentType: "ipfs",
		Address:        circuitAddr,
		IpfsURI:        ipfs,
	}
	setReceipt, err := client.RegisterCircuit(regClient, userSigner, circuitName, &cMeta)
	if err != nil {
		log.Fatal("Could not register circuit metadata on chain", err)
		return
	}
	elapsedTime = time.Since(startTime)
	// get total gas used
	calculateGasUsage("SetCircuit in registry", setReceipt)
	fmt.Printf("SetCircuit in registry time: %s\n", elapsedTime)
	startTime = time.Now()
	circuitData, _ := client.GetCircuit(regClient, circuitName)
	elapsedTime = time.Since(startTime)
	fmt.Printf("GetCircuit from registry time: %s\n", elapsedTime)
	fmt.Printf("CircuitMeta data from registry: %s\n", circuitData)
}

func SimulateAttestation(claimOwner *client.Signer, attestor *client.Signer, reader *client.TxReader,
	ipfsConn string, registryAddr string, userIdentityAddr string) {
	fmt.Println("Simulation type: attestation")

	// register circuit metadata on registry
	regClient := client.NewRegistryClient(registryAddr, *reader)
	manClient := client.NewManagerClient(userIdentityAddr, *reader, ipfsConn)

	//check clients are valid
	if regClient.Registry == nil || manClient.Manager == nil {
		log.Fatal("Check contract addresses")
		return
	}
	///* simulation pre-requirements:
	//1. IdentityManager should be deployed for both user and attestor
	//*/
	//// ownerAccount deploys the registry
	//startTime := time.Now()
	//registryContract, _ := deployment.DeployRegistry(attestor)
	//registryAddr := registryContract.String()
	//elapsedTime := time.Since(startTime)
	//fmt.Printf("Registry deployment took: %s\n", elapsedTime)
	//fmt.Println("Registry contract deployed at:", registryContract)
	//
	//// userAccount deploys it's own manager
	//startTime = time.Now()
	//userManager, _ := deployment.DeployManager(claimOwner, registryContract)
	//elapsedTime = time.Since(startTime)
	//fmt.Printf("Manager deployment took: %s\n", elapsedTime)
	//fmt.Println("Manager contract for user deployed at:", userManager)
	//
	//// attestorAccount deploys it's own manager
	//attestorManager, _ := deployment.DeployManager(attestor, registryContract)
	//fmt.Println("Manager contract for attestor user deployed at:", attestorManager)
	//
	//startTime = time.Now()
	//regClient := client.NewRegistryClient(registryAddr, *reader)
	//receipt, err := client.RegisterIdentity(regClient, claimOwner, attestor, userManager)
	//if err != nil {
	//	log.Fatal("Could not register identity", err)
	//}
	//elapsedTime = time.Since(startTime)
	//fmt.Printf("RegisterIdentity took: %s\n", elapsedTime)
	//// check receipt status reverted or not
	//if receipt.Status == 0 {
	//	log.Fatal("Transaction reverted by contract", receipt.TxHash.Hex())
	//	return
	//}
	//fmt.Println("Identity registered for", common.Bytes2Hex(claimOwner.PublicKey))
	//receipt, err = client.RegisterIdentity(regClient, attestor, attestor, attestorManager)
	//if err != nil {
	//	log.Fatal("Could not register identity", err)
	//}
	//// check receipt status reverted or not
	//if receipt.Status == 0 {
	//	log.Fatal("Transaction reverted by contract")
	//	return
	//}
	//fmt.Println("Identity registered for", common.Bytes2Hex(attestor.PublicKey))

	//2. Private claim should be set for user => circuit should be deployed and registered on chain)
	// generate private claim (proof) and deploy verifier circuit
	circuitName := "MimcComp"
	ipfsClient := client.NewIpfsClient(ipfsConn)
	// first build the circuit

	curve := ecc.BN254
	startTime := time.Now()
	var (
		pre       = 100
		threshold = 22
		digest    = "20370067689261511688289967978544823130432235585709842144916192060767982363628"
	)

	var mimcCircuit mimc.MimcComp

	cc, err := frontend.Compile(curve.ScalarField(), scs.NewBuilder, &mimcCircuit)
	if err != nil {
		log.Fatal(err)
	}

	//wt
	var mimcWt mimc.MimcComp
	mimcWt.Input = pre
	mimcWt.Threshold = threshold
	mimcWt.Digest = digest
	mimcWt.Nonce = 0
	mimcWt.NoncePriv = 0
	mimcWt.Sender = claimOwner.PublicKey
	mimcWt.SenderPriv = claimOwner.PublicKey

	w, err := frontend.NewWitness(&mimcWt, curve.ScalarField())
	if err != nil {
		log.Fatal("Couldn't set witness ", err)
	}

	srs, err := test.NewKZGSRS(cc)
	if err != nil {
		log.Fatal("test.NewKZGSRS(ccs)")
	}
	// groth16 zkSNARK: Setup
	pk, vk, err := plonk.Setup(cc, srs)
	if err != nil {
		log.Fatal("plonk.Setup", err)
	}
	path := fmt.Sprintf(solidityPath, circuitName)
	elapsedTime := time.Since(startTime)
	fmt.Printf("CircuitMeta setup took: %s\n", elapsedTime)

	//// groth16: Prove & Verify
	proof, err := plonk.Prove(cc, pk, w)
	if err != nil {
		log.Fatal("prove computation failed...", err)
	}

	startTime = time.Now()
	deployment.CircuitToSolidity(vk, path)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta solidity generation took: %s\n", elapsedTime)

	startTime = time.Now()
	deployment.CreateCircuitBindings(circuitName)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta bindings generation took: %s\n", elapsedTime)

	startTime = time.Now()
	circuitAddr, _ := deployment.DeployCircuit(attestor)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta deployment took: %s\n", elapsedTime)
	fmt.Printf("CircuitMeta deployed at: %s\n", circuitAddr)

	startTime = time.Now()
	serialized := circuits.Serialize(cc, w, vk, pk, proof, circuitName, circuitAddr.Hex())
	ipfs, err := circuits.PublishCircuitOnIpfs(ipfsClient, &serialized)
	if err != nil {
		log.Fatal("Could not publish circuit on ipfs", err)
		return
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta metadata ipfs-publish time: %s\n", elapsedTime)
	log.Printf("CircuitMeta published on ipfs: %s\n", ipfs)

	startTime = time.Now()
	cMeta := client.CircuitMeta{
		DeploymentType: "ipfs",
		Address:        circuitAddr,
		IpfsURI:        ipfs,
	}
	receipt, err := client.RegisterCircuit(regClient, attestor, circuitName, &cMeta)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta metadata registration on chain took: %s\n", elapsedTime)
	calculateGasUsage("CircuitMeta registration", receipt)

	startTime = time.Now()
	var buf bytes.Buffer
	_, err = proof.WriteTo(&buf)
	commitment := string(buf.Bytes()[:32])
	fmt.Println("length of proof:", len(commitment))
	// set private claim for claimOwner
	receipt, err = client.SetPrivateClaim(manClient, claimOwner, "age", commitment, "age > 22", ipfs, "0x0")
	if err != nil {
		log.Fatal("Could not set private claim", err)
		return
	}

	elapsedTime = time.Since(startTime)
	fmt.Printf("Private claim set took: %s\n", elapsedTime)
	calculateGasUsage("Private claim set", receipt)

	startTime = time.Now()
	// attestor read claim from user contract
	privateClaim, err := client.GetPrivateClaim(manClient, "age")
	if err != nil {
		log.Fatal("Could not get private claim", err)
		return
	}
	// sign the claim with attestor {t, claim={commitment/public_value}, statement}
	jsonMeta, err := json.Marshal(privateClaim.ClaimMeta)
	if err != nil {
		return
	}

	fmt.Println("size privClaim 32", len(jsonMeta))

	claimSig, err := attestor.Sign(jsonMeta) // we assume signature is sent to user in a secure channel
	expires := new(big.Int).SetUint64(1000000000000000000)
	id := new(big.Int).SetUint64(uint64(privateClaim.ClaimMeta.Id))
	// user gets signature related with the claim and sets attestation in usersigner manager contract
	receipt, err = client.SetAttestation(manClient, claimOwner, "age", common.BytesToAddress(attestor.PublicKey),
		claimSig, *expires, *id)

	if err != nil {
		log.Fatal("Could not set attestation", err)
		return
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("Attestation set took: %s\n", elapsedTime)
	calculateGasUsage("Set attestation", receipt)

	// attestor reads attestation from user contract
	att, err := client.GetAttestation(manClient, "age")
	if err != nil {
		log.Fatal("Could not get attestation", err)
		return
	}

	fmt.Println("size of attestation:", att.Size())

}
func SimulateRevocation(claimOwner *client.Signer, attestor *client.Signer, reader *client.TxReader, ipfsConn string,
	registryAddr string, userIdentityAddr string, attestorIdentityAddr string) {
	fmt.Println("Simulation type: revocation")
	regClient := client.NewRegistryClient(registryAddr, *reader)
	userIdentityClient := client.NewManagerClient(userIdentityAddr, *reader, ipfsConn)
	attestorIdentityClient := client.NewManagerClient(attestorIdentityAddr, *reader, ipfsConn)

	circuitName := "MimcComp"
	ipfsClient := client.NewIpfsClient(ipfsConn)
	// first build the circuit

	curve := ecc.BN254
	startTime := time.Now()
	var (
		pre       = 100
		threshold = 22
		digest    = "20370067689261511688289967978544823130432235585709842144916192060767982363628"
	)

	var mimcCircuit mimc.MimcComp

	cc, err := frontend.Compile(curve.ScalarField(), scs.NewBuilder, &mimcCircuit)
	if err != nil {
		log.Fatal(err)
	}

	//wt
	var mimcWt mimc.MimcComp
	mimcWt.Input = pre
	mimcWt.Threshold = threshold
	mimcWt.Digest = digest
	mimcWt.Nonce = 0
	mimcWt.NoncePriv = 0
	mimcWt.Sender = claimOwner.PublicKey
	mimcWt.SenderPriv = claimOwner.PublicKey

	w, err := frontend.NewWitness(&mimcWt, curve.ScalarField())
	if err != nil {
		log.Fatal("Couldn't set witness ", err)
	}

	srs, err := test.NewKZGSRS(cc)
	if err != nil {
		log.Fatal("test.NewKZGSRS(ccs)")
	}
	// groth16 zkSNARK: Setup
	pk, vk, err := plonk.Setup(cc, srs)
	if err != nil {
		log.Fatal("plonk.Setup", err)
	}
	path := fmt.Sprintf(solidityPath, circuitName)
	elapsedTime := time.Since(startTime)
	fmt.Printf("CircuitMeta setup took: %s\n", elapsedTime)

	//// groth16: Prove & Verify
	proof, err := plonk.Prove(cc, pk, w)
	if err != nil {
		log.Fatal("prove computation failed...", err)
	}

	startTime = time.Now()
	deployment.CircuitToSolidity(vk, path)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta solidity generation took: %s\n", elapsedTime)

	startTime = time.Now()
	deployment.CreateCircuitBindings(circuitName)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta bindings generation took: %s\n", elapsedTime)

	startTime = time.Now()
	circuitAddr, _ := deployment.DeployCircuit(attestor)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta deployment took: %s\n", elapsedTime)
	fmt.Printf("CircuitMeta deployed at: %s\n", circuitAddr)

	startTime = time.Now()
	serialized := circuits.Serialize(cc, w, vk, pk, proof, circuitName, circuitAddr.Hex())
	ipfs, err := circuits.PublishCircuitOnIpfs(ipfsClient, &serialized)
	if err != nil {
		log.Fatal("Could not publish circuit on ipfs", err)
		return
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta metadata ipfs-publish time: %s\n", elapsedTime)
	log.Printf("CircuitMeta published on ipfs: %s\n", ipfs)

	// register circuit metadata on registry
	startTime = time.Now()
	cMeta := client.CircuitMeta{
		DeploymentType: "ipfs",
		Address:        circuitAddr,
		IpfsURI:        ipfs,
	}
	receipt, err := client.RegisterCircuit(regClient, attestor, circuitName, &cMeta)
	receipt, err = attestor.WaitForReceipt(receipt.TxHash)
	elapsedTime = time.Since(startTime)
	fmt.Printf("CircuitMeta metadata registration on chain took: %s\n", elapsedTime)
	calculateGasUsage("CircuitMeta registration", receipt)

	startTime = time.Now()
	var buf bytes.Buffer
	_, err = proof.WriteTo(&buf)
	commitment := string(buf.Bytes())
	// set private claim for claimOwner
	//manClient := client.NewManagerClient(userManager.Hex(), *reader, ipfsConn)
	//attestorManClient := client.NewManagerClient(attestorManager.Hex(), *reader, ipfsConn)

	receipt, err = client.SetPrivateClaim(userIdentityClient, claimOwner, "age", commitment, "age > 22", ipfs, "0x0")
	if err != nil {
		log.Fatal("Could not set private claim", err)
		return
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("Private claim set took: %s\n", elapsedTime)
	calculateGasUsage("Private claim set", receipt)

	startTime = time.Now()
	// attestor read claim from user contract
	privateClaim, err := client.GetPrivateClaim(userIdentityClient, "age")
	if err != nil {
		log.Fatal("Could not get private claim", err)
		return
	}
	// sign the claim with attestor {t, claim={commitment/public_value}, statement}
	jsonMeta, err := json.Marshal(privateClaim.ClaimMeta)
	if err != nil {
		return
	}
	// fmt.Println("size privateclaim", len(jsonMeta))

	claimSig, err := attestor.Sign(jsonMeta) // we assume signature is sent to user in a secure channel
	expires := new(big.Int).SetUint64(1000000000000000000)
	id := new(big.Int).SetUint64(uint64(privateClaim.ClaimMeta.Id))
	// user gets signature related with the claim and sets attestation in usersigner manager contract
	receipt, err = client.SetAttestation(userIdentityClient, claimOwner, "ageAttestation", common.BytesToAddress(attestor.PublicKey),
		claimSig, *expires, *id)

	if err != nil {
		log.Fatal("Could not set attestation", err)
		return
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("Attestation set took: %s\n", elapsedTime)
	calculateGasUsage("Set attestation", receipt)

	// attestor reads attestation from user contract
	t1 := time.Now()
	attestation, err := client.GetAttestation(userIdentityClient, "age")
	if err != nil {
		log.Fatal("Could not get attestation", err)
		return
	}
	fmt.Println("getAttestation from chain took:", time.Since(t1))
	fmt.Println("Pre-requirements are set up successfully. Running actual simulation...")
	fmt.Println(strings.Repeat("*", 100))
	//3. Attestor should revoke the claim
	startTime = time.Now()
	id = new(big.Int).SetInt64(int64(attestation.Id))
	receipt, err = client.Revoke(attestorIdentityClient, attestor, "ageAttestation",
		common.BytesToAddress(claimOwner.PublicKey), id, "Invalid age")
	if err != nil {
		log.Fatal("Could not revoke claim", err)
		return
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("Revocation took: %s\n", elapsedTime)
	calculateGasUsage("Revocation", receipt)

	//4. User should be able to see revoked attestations
	startTime = time.Now()
	revocation, err := client.GetRevocation(attestorIdentityClient, "ageAttestation")
	if err != nil {
		log.Fatal("Could not get attestation", err)
		return
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("Revocation from contract: %s %s %d %d\n", revocation.AttestedTo, revocation.Status,
		revocation.AttestationId, revocation.Timestamp)
	fmt.Printf("Get revocation took: %s\n", elapsedTime)

	fmt.Println("size revocation:", revocation.Size())
}

func sizeASDF(s []client.Attestation) int {
	size := 0
	s = s[:cap(s)]
	size += cap(s) * int(unsafe.Sizeof(s))
	for i := range s {
		size += (&s[i]).Size()
	}
	return size
}

func GetCircuitInfoFromRegistry(registryAddr string, circuitName string, reader *client.TxReader) {
	regClient := client.NewRegistryClient(registryAddr, *reader)
	circuitInfo, err := client.GetCircuit(regClient, circuitName)
	if err != nil {
		log.Fatal("Could not get circuit info", err)
		return
	}
	fmt.Printf("CircuitMeta info: %s %s %s %s\n", circuitInfo.Address, circuitInfo.DeploymentType, circuitInfo.IpfsURI)
}

func weiToEther(wei uint64) *big.Float {
	weiInEther := new(big.Float).SetUint64(wei)
	ether := new(big.Float).Quo(weiInEther, big.NewFloat(1e18))
	return ether
}

func calculateGasUsage(operation string, receipt *types.Receipt) {
	fmt.Printf("%s Transaction Cost (Wei): %d \n", operation, receipt.GasUsed)
	fmt.Printf("%s Transaction Cost (Ether): %s\n", operation, weiToEther(receipt.GasUsed).String())
	total_cost := new(big.Int)
	total_gas := new(big.Int).SetUint64(receipt.GasUsed)
	total_cost.Mul(total_gas, receipt.EffectiveGasPrice)
	fmt.Printf("Total Cost (Ether): %s\n", weiToEther(total_cost.Uint64()).String())
}

func GenerateDummyClaim() []byte {
	claimDict := make(map[string]string)
	claimDict["age"] = strconv.Itoa(rand.Int())
	claimDict["timestamp"] = strconv.Itoa(rand.Int())
	claimDict["revoked"] = "false"
	claimDict["subject"] = "0x2e144aF3Bde9B518C7C65FBE170c07c888f1fF1a"
	jsonStr, _ := json.Marshal(claimDict)
	fmt.Println("dummy claim size:", len(jsonStr))
	return jsonStr
}
