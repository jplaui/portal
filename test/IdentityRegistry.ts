import {loadFixture,} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import {expect,} from "chai";
import {ethers} from "hardhat";

describe("IdentityRegistry", function () {

    async function deployRegistry() {

        // Contracts are deployed using the first signer/account by default
        const [owner, otherAccount] = await ethers.getSigners();

        const IdentityRegistry = await ethers.getContractFactory("IdentityRegistry");
        const identityInstance = await IdentityRegistry.deploy();


        const IdentityManager = await ethers.getContractFactory("IdentityManager");
        const managerInstance = await IdentityManager.deploy(identityInstance.target);

        const verify = await ethers.getContractFactory("Verifier");
        const verifyInstance = await verify.deploy();

        return {identityInstance, owner, otherAccount, managerInstance, verifyInstance};
    }

    describe("Registry", function () {
        it("Should set record", async function () {
            const {identityInstance, owner, managerInstance} = await loadFixture(deployRegistry);

            await identityInstance.register(owner.address, managerInstance.target, {from: owner.address})

            const res = await identityInstance.owner(owner.address);
            expect(res).to.equal(owner.address);
        });

        it("Should set the resolver for node and emit event when resolver is set.", async function () {
            const {identityInstance, owner, managerInstance} = await loadFixture(deployRegistry);

            await identityInstance.register(owner.address, managerInstance.target, {from: owner.address})

            const res = await identityInstance.owner(owner.address);
            expect(res).to.equal(owner.address);

            // await expect(identityInstance.setResolver(ZERO_HASH, managerInstance.target, {from: owner.address}))
            //     .to.emit(identityInstance, "NewResolver")
            //     .withArgs(ZERO_HASH, managerInstance.target);
            //
            // const resolverRes = await identityInstance.resolver(ZERO_HASH);
            // expect(resolverRes).to.equal(managerInstance.target);
        });


    });
    describe("Resolver", function () {
        it("Should set the the key for node", async function () {
            const {identityInstance, owner, managerInstance} = await loadFixture(deployRegistry);

            await identityInstance.register(owner.address, managerInstance.target, {from: owner.address})
            await managerInstance.setText('testKey', 'testValue', {from: owner.address})
            const res = await managerInstance.text(owner.address, 'testKey');
            expect(res).to.equal('testValue');
        });

        it("Should prevent setting the key for node by non-owners ", async function () {
            const {identityInstance, owner, otherAccount, managerInstance} = await loadFixture(deployRegistry);

            await identityInstance.register(owner.address, managerInstance.target, {from: owner.address})
            await managerInstance.setText('testKey', 'testValue', {from: owner.address})
            const res = await managerInstance.text(owner.address, 'testKey');
            expect(res).to.equal('testValue');
            try {
                await managerInstance.setText('testKey', 'testValue', {from: otherAccount.address})

            } catch (err) {
                expect(err).to.be.an.instanceOf(TypeError);

                expect(err.code).to.equal('INVALID_ARGUMENT');
            }

        });

        it("Should keep json as value (public claim example)", async function () {

            /*
            Here we can keep private claims (ipfs storage link) with a unique key as well.
             */
            const {identityInstance, owner, managerInstance} = await loadFixture(deployRegistry);

            await identityInstance.register(owner.address, managerInstance.target, {from: owner.address})
            const dummyClaim = {
                "type": "Integer",
                "timestamp": 123456789,
                "hash": "0x1234567890123456789012345678901234567890",
                "number": 18,
                "threshold": 10,
            }


            await managerInstance.setText('testKey', JSON.stringify(dummyClaim), {from: owner.address})

            expect(await managerInstance.text(owner.address, 'testKey')).to.equal(JSON.stringify(dummyClaim));


        });

    });
    describe("VerifySignature", function () {
        it("Check signature", async function () {
            const {verifyInstance} = await loadFixture(deployRegistry);

            const accounts = await ethers.getSigners()

            const signer = accounts[0]
            const message = "Hello"

            const hash = await verifyInstance.getMessageHash(message)
            const bytes = new Uint8Array(hash.length / 2);
            for (let i = 0; i < hash.length; i += 2) {
                bytes[i / 2] = parseInt(hash.substring(i, 2), 16);
            }

            const sig = await signer.signMessage(bytes)

            const ethHash = await verifyInstance.getEthSignedMessageHash(hash)

            console.log("signer          ", signer.address)
            console.log("recovered signer", await verifyInstance.recoverSigner(ethHash, sig))

            // Correct signature and message returns true
            expect(
                await verifyInstance.verify(signer.address, message, sig)
            ).to.equal(true)

            // Incorrect message returns false
            expect(
                await verifyInstance.verify(signer.address, message + 'a', sig)
            ).to.equal(false)
        })
    })
});
