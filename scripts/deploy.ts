import {ethers} from "hardhat";

async function main() {
    //await deployContracts();
    const registryAddr = '0xf3585FCD969502624c6A8ACf73721d1fce214E83';
    const resolverAddr = '0x2e144aF3Bde9B518C7C65FBE170c07c888f1fF1a';


    await register(resolverAddr);

}

async function deployContracts() {
    const identityRegistry = await ethers.deployContract("IdentityRegistry");
    await identityRegistry.waitForDeployment();
    console.log(
        `IdentityRegistry deployed to ${identityRegistry.target}`);

    const manager = await ethers.deployContract("IdentityManager", [identityRegistry.target]);

    await manager.waitForDeployment();
    console.log(
        `Resolver deployed to ${manager.target}`);

}

async function register(contractAddr: string) {
    const accs = await ethers.getSigners()

    const IdentityRegistry = await ethers.getContractAt("IdentityRegistry", contractAddr);
    console.log(accs[0].address)
    //console.log(IdentityRegistry);

    const res = await IdentityRegistry.register(accs[1].address, contractAddr, {from: accs[1].address})
    console.log(res);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});

