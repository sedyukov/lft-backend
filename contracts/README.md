# Shared Wallet with allowance
Solidity smart contract that handle allowance control for a distributed budget.

## Contract
The contract files are at `./contracts` folder. Third-party contracts like Open Zeppelin are stored at `./contracts/vendor` folder,
while compiled contracts are stored at `./contracts/bin` folder

### Compile contracts
In order to compile the contracts the `ethereum/solc` docker image is used as following:

```shell
docker run --rm -v {WORK_DIR}/contracts:/contracts ethereum/solc:0.7.6 \
/contracts/solidity/LevelFiveToken.sol --bin --abi --optimize --overwrite -o /contracts/bin
```

example:
```shell
docker run --rm -v ~/Documents/training/go/lft-backend/contracts:/contracts ethereum/solc:0.7.6 \
/contracts/solidity/LevelFiveToken.sol --bin --abi --optimize --overwrite -o /contracts/bin
```
To generate a Go interface and bindings based on the contracts, the `ethereum/client-go:alltools-latest` docker image is used as following:

```shell
docker run --rm -v {WORK_DIR}/contracts:/contracts ethereum/client-go:alltools-latest \
abigen --abi /contracts/bin/LevelFiveToken.abi --pkg contracts --type Contract --out \
/contracts/interfaces/level-five-token.go  --bin /contracts/bin/LevelFiveToken.bin
```

example:
```shell
docker run --rm -v ~/Documents/training/go/lft-backend/contracts:/contracts ethereum/client-go:alltools-latest \
abigen --abi /contracts/bin/LevelFiveToken.abi --pkg contracts --type Contract --out \
/contracts/interfaces/level-five-token.go  --bin /contracts/bin/LevelFiveToken.bin
```

## Geth

Run Geth docker image as:
```shell
docker run --rm --name ethereum-node -v /Users/steven/code/eth/:/root -p 8545:8545 -p 30303:30303 ethereum/client-go --nodiscover
```

and connect to the instance as following in order to perform CLI operations:
```shell
docker exec -it <> /bin/sh

$ geth attach /root/.ethereum/geth.ipc
```
