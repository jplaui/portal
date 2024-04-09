## Identity registry and resolver contracts


### Prerequisites
1. ganache v7.8.0 (@ganache/cli: 0.9.0, @ganache/core: 0.9.0)
2. go version go1.20.5 darwin/arm64
3. solc Version: 0.8.20+commit.a1b79de6.Darwin.appleclang 
4. abigen version 1.10.16-unstable
5. IPFS-kubo (for testing ipfs claims)
   - `docker pull ipfs/kubo`
   -  Configure the ipfs node:
     - Set local volume mappings
         - `export ipfs_staging=</absolute/path/to/>`
         - `export ipfs_data=</absolute/path/to/>`
     - Run the docker container by mapping volumes and ports
         - 4001 -> P2P TCP/QUIC transports,
         - 5001 -> RPC API,
         - 8080->Gateway

   - `docker run -d --name ipfs_host -v $ipfs_staging:/export -v $ipfs_data:/data/ipfs -p 4001:4001 -p 4001:4001/udp -p 127.0.0.1:8080:8080 -p 127.0.0.1:5001:5001 ipfs/kubo:latest`
   - Check the status of the ipfs node by running the following command: ```docker logs -f ipfs_host```
   ```
   # Expected logs when the system is started.
   RPC API server listening on /ip4/0.0.0.0/tcp/5001
   WebUI: http://0.0.0.0:5001/webui
   Gateway server listening on /ip4/0.0.0.0/tcp/8080
   Daemon is ready
   ```
6. Connect to peers:```docker exec ipfs_host ipfs swarm peers```
7. Stop container after the test: ```docker stop ipfs_host```
8. Update `ipfsConn` constant in main.go to your ipfs node address.

### Go-sdk setup
1. Run ganache network first `ganache -v -m "much repair shock carbon improve miss forget sock include bullet interest solution"`
2. Run go bindings (if not exist under go-sdk/contracts repository): `go run main.go -bindings true`
3. Overwrite go bindings under `go-sdk/contracts` in packages `go-sdk/contracts/IdentityInterface/` `go-sdk/contracts/IdentityManager/` `go-sdk/contracts/IdentityRegistry/` (make sure to move in vscode and overwrite the three .go files /contracts/IdentityInterface.go,IdentityManager.go,IdentityRegistry.go) and leave MimcComp.go where it is...
4. Deploy contracts with: `go run main.go -deploy true`

command will output something like:
```
Registry contract deployed at: 0xf3585FCD969502624c6A8ACf73721d1fce214E83
Manager contract deployed at: 0x2e144aF3Bde9B518C7C65FBE170c07c888f1fF1a
```
5. Update `managerAddr` and `registryAddr` in main.go constants with the deployment command output addresses.
6. Register identity: ```go run main.go -register true```
7. Set public claim for user: ```go run main.go -claim true```
8. Test individual ipfs claim: ```go run main.go -ipfs true```
9. Test on-chain ipfs claims: ```go run main.go -ipfs-on-chain true```
10. Test zero knowledge proove & verify from ipfs serialized circuits: ```go run main.go -circuits true```

### How to run simulations:
1. Run ganache network first `ganache -v -m "much repair shock carbon improve miss forget sock include bullet interest solution"`
2. Run go bindings `go run main.go -bindings true`
3. Overwrite go bindings under `go-sdk/contracts` in packages `go-sdk/contracts/IdentityInterface/` `go-sdk/contracts/IdentityManager/` `go-sdk/contracts/IdentityRegistry/`
4. Deploy contracts with: `go run main.go -deploy true`
5. Run `go run main.go -sim {SIM_TYPE}`. Restarting ganache server per each simulation is recommended. Available simulation types: `ipfs, ipfs-on-chain, onchain, deployment`

### Setup
1. Install dependencies: run `npm install` in main project directory.
2. `npx hardhat test` to run tests. Tests cover the identity registry and resolver contracts cases:
    ```
   1. Deployment cases
   2. Identity registry registration 
   3. Identity registry resolvers
   4. Resolver set for public claims, merkle root and possibly (in future) private claims
   ```

```shell

npx hardhat test
npx hardhat run scripts/deploy.ts
```
