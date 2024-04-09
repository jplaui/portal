package deployment

import (
	"fmt"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go-sdk/client"
	mimcComp "go-sdk/contracts"
	"go-sdk/contracts/IdentityManager"
	"go-sdk/contracts/IdentityRegistry"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func DeployRegistry(signer *client.Signer, contractArgs ...interface{}) (common.Address, *types.Transaction) {

	auth := signer.BindTxOpts()
	address, tx, instance, err := IdentityRegistry.DeployIdentityRegistry(auth, signer.Client)
	if err != nil {
		log.Fatal(err)
	}

	_ = instance

	return address, tx

}

func DeployManager(signer *client.Signer, registryAddr common.Address) (common.Address, *types.Transaction) {

	auth := signer.BindTxOpts()
	managerAddr, tx, manager, err := IdentityManager.DeployIdentityManager(auth, signer.Client, registryAddr)
	if err != nil {
		log.Fatal(err)
	}

	_ = manager

	return managerAddr, tx

}
func CreateBindings(path string, contract string) {
	dir, err := os.Getwd()

	// compile solidity contract
	filePath := fmt.Sprintf("%s/%s.sol", path, contract)
	fullPath := filepath.Join(dir, filePath)

	cmd := exec.Command("solc", "--overwrite", "--abi", "--bin", fullPath, "-o", "build", "--include-path=../node_modules", "--base-path=..")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	assertNoError(err)

	abi := fmt.Sprintf("%s/%s.%s", "--abi=build", contract, "abi")
	bin := fmt.Sprintf("%s/%s.%s", "--bin=build", contract, "bin")
	pkg := fmt.Sprintf("%s%s", "--pkg=", contract)
	out := fmt.Sprintf("%s/%s.%s", "--out=contracts", contract, "go")
	cmd = exec.Command("abigen", abi, bin, pkg, out)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	assertNoError(err)

}

func CreateCircuitBindings(circuitName string) {
	//pkg := fmt.Sprintf("%s%s%s", "--pkg=", circuitName, "CircuitMeta")
	pkg := fmt.Sprintf("%s%s", "--pkg=", circuitName)
	out := fmt.Sprintf("%s/%s.%s", "--out=contracts", circuitName, "go")
	cmd := exec.Command("abigen", "--abi=build/PlonkVerifier.abi", "--bin=build/PlonkVerifier.bin", pkg, out)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	assertNoError(err)
}
func assertNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CircuitToSolidity(vk plonk.VerifyingKey, solidityPath string) {
	dir, err := os.Getwd()
	path := filepath.Join(dir, solidityPath)
	f, err := os.Create(path)
	err = vk.ExportSolidity(f)
	if err != nil {
		log.Fatal("vk.ExportSolidity ", err)
	}
}

func DeployCircuit(signer *client.Signer) (common.Address, *types.Transaction) {

	auth := signer.BindTxOpts()
	circuitContractAddr, tx, _, err := mimcComp.DeployMimcComp(auth, signer.Client)
	if err != nil {
		log.Fatal(err)
	}

	return circuitContractAddr, tx

}
