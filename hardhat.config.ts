import {HardhatUserConfig} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";



const config: HardhatUserConfig = {
    networks: {
        ganache: {
            url: "http://127.0.0.1:8545",
            accounts: {
                mnemonic: "much repair shock carbon improve miss forget sock include bullet interest solution"
            }
        },
        hardhat: {
            chainId: 1337,
            accounts: {
                mnemonic: "much repair shock carbon improve miss forget sock include bullet interest solution"
            },
        },
        // hardhat: {
        //   accounts: [
        //     {
        //       balance: "10000000000000000000000",
        //       privateKey:
        //         "0xe87d780e4c31c953a68aef2763df56599c9cfe73df4740fc24c2d0f5acd21bae",
        //     },
        //   ],
        // },
    },

    solidity: {
        compilers: [
            {
                version: "0.8.20",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200,
                    },
                },
            },
        ],
    }

};
export default config;
