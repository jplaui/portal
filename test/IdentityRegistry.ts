import {loadFixture,} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import {expect,} from "chai";
import {ethers} from "hardhat";

describe("IdentityRegistry", function () {

    const ZERO_HASH = '0x0000000000000000000000000000000000000000000000000000000000000000'

    async function deployRegistry() {

        // Contracts are deployed using the first signer/account by default
        const [owner, otherAccount] = await ethers.getSigners();

        const IdentityRegistry = await ethers.getContractFactory("IdentityRegistry");
        const identityInstance = await IdentityRegistry.deploy();


        const IdentityResolver = await ethers.getContractFactory("IdentityResolver");
        const resolverInstance = await IdentityResolver.deploy(identityInstance.target);

        return {identityInstance, owner, otherAccount, resolverInstance};
    }

    describe("Registry", function () {
        it("Should set the owner for node", async function () {
            const {identityInstance, owner} = await loadFixture(deployRegistry);

            let addr = '0x0000000000000000000000000000000000001234'

            await identityInstance.setOwner(ZERO_HASH, addr, {from: owner.address})

            const res = await identityInstance.owner(ZERO_HASH);
            expect(res).to.equal(addr);
        });

        it("Should set the resolver for node and emit event when resolver is set.", async function () {
            const {identityInstance, owner, resolverInstance} = await loadFixture(deployRegistry);

            await identityInstance.setOwner(ZERO_HASH, owner.address, {from: owner.address})

            const res = await identityInstance.owner(ZERO_HASH);
            expect(res).to.equal(owner.address);

            await expect(identityInstance.setResolver(ZERO_HASH, resolverInstance.target, {from: owner.address}))
                .to.emit(identityInstance, "NewResolver")
                .withArgs(ZERO_HASH, resolverInstance.target);

            const resolverRes = await identityInstance.resolver(ZERO_HASH);
            expect(resolverRes).to.equal(resolverInstance.target);
        });


        it('should allow setting the TTL', async () => {

            const {identityInstance, owner} = await loadFixture(deployRegistry);

            await identityInstance.setOwner(ZERO_HASH, owner.address, {from: owner.address})

            const res = await identityInstance.owner(ZERO_HASH);
            expect(res).to.equal(owner.address);


            await identityInstance.setTTL(ZERO_HASH, 3600, {from: owner.address})
            let resp: bigint = await identityInstance.ttl(ZERO_HASH)
            expect(Number(resp)).to.equal(3600);
        });
        it('should prevent setting the TTL by non-owners', async () => {
            const {identityInstance, owner, otherAccount} = await loadFixture(deployRegistry);

            await identityInstance.setOwner(ZERO_HASH, owner.address, {from: owner.address})

            const res = await identityInstance.owner(ZERO_HASH);
            expect(res).to.equal(owner.address);
            console.log('owner', owner.address)
            console.log('otherAccount', otherAccount.address)
            try {
                await identityInstance.setTTL(ZERO_HASH, 3600, {from: otherAccount.address})
            } catch (err) {
                expect(err).to.be.an.instanceOf(TypeError);

                expect(err.code).to.equal('INVALID_ARGUMENT');
            }
        })
    });
    describe("Resolver", function () {
        it("Should set the the key for node", async function () {
            const {identityInstance, owner, resolverInstance} = await loadFixture(deployRegistry);

            await identityInstance.setOwner(ZERO_HASH, owner.address, {from: owner.address})
            await identityInstance.setResolver(ZERO_HASH, resolverInstance.target, {from: owner.address})

            await resolverInstance.setText(ZERO_HASH, 'testKey', 'testValue', {from: owner.address})
            const res = await resolverInstance.text(ZERO_HASH, 'testKey');
            expect(res).to.equal('testValue');
        });

        it("Should prevent setting the key for node by non-owners ", async function () {
            const {identityInstance, owner, resolverInstance, otherAccount} = await loadFixture(deployRegistry);

            await identityInstance.setOwner(ZERO_HASH, owner.address, {from: owner.address})
            await identityInstance.setResolver(ZERO_HASH, resolverInstance.target, {from: owner.address})
            try {
                await resolverInstance.setText(ZERO_HASH, 'testKey', 'testValue', {from: otherAccount.address})

            } catch (err) {
                expect(err).to.be.an.instanceOf(TypeError);

                expect(err.code).to.equal('INVALID_ARGUMENT');
            }

        });

        it("Should keep json as value (public claim example)", async function () {

            /*
            Here we can keep private claims (ipfs storage link) with a unique key as well.
             */
            const {identityInstance, owner, resolverInstance} = await loadFixture(deployRegistry);
            const dummyClaim = {
                "type": "Integer",
                "timestamp": 123456789,
                "hash": "0x1234567890123456789012345678901234567890",
                "number": 18,
                "threshold": 10,
            }
            await identityInstance.setOwner(ZERO_HASH, owner.address, {from: owner.address})
            await identityInstance.setResolver(ZERO_HASH, resolverInstance.target, {from: owner.address})

            await resolverInstance.setText(ZERO_HASH, 'testKey', JSON.stringify(dummyClaim), {from: owner.address})

            expect(await resolverInstance.text(ZERO_HASH, 'testKey')).to.equal(JSON.stringify(dummyClaim));


        });
        it("Should keep merkle root as contentHash", async function () {
            const {identityInstance, owner, resolverInstance} = await loadFixture(deployRegistry);

            await identityInstance.setOwner(ZERO_HASH, owner.address, {from: owner.address})
            await identityInstance.setResolver(ZERO_HASH, resolverInstance.target, {from: owner.address})

            await resolverInstance.setContenthash(ZERO_HASH, '0x1234567890123456789012345678901234567890', {from: owner.address})
            expect(await resolverInstance.contenthash(ZERO_HASH)).to.equal('0x1234567890123456789012345678901234567890');


        });
    });
});
