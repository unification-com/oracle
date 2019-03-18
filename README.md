# Interacting with the Mainchain

## Connect to the Test-net

Install geth from Unification
```
go get github.com/unification-com/mainchain
go install github.com/unification-com/mainchain/cmd/geth
```

### Initialize the genesis block
```
geth init ~/.go/src/github.com/unification-com/haiku-core/native/mainchain/Docker/assets/genesis.json
```

### Copy in the static nodes
```
cp ~/.go/src/github.com/unification-com/haiku-core/native/mainchain/Docker/validator/bootnode_keys/static-nodes.json /Users/indika/Library/UndWorkchain/geth
```

### And start up your light client
```
geth --networkid "50001" --rpc --rpcaddr "127.0.0.1" --rpcapi "eth,web3,net,admin,debug,db,personal,miner" --rpccorsdomain "*" --rpcvhosts "*" --rpcport 8545  --verbosity=4 --nodiscover --nodekey="~/Library/Ethereum/bootnode.key"  --syncmode="fast" 
```

## Create a space
First, create a space on the Mainchain to deposit your data.
Supply your private key.
```
go run $GOPATH/src/github.com/unification-com/haiku-core/oracle/apply.go "http://localhost:8545" "07a4dc42e81b0e7c4e2e639741983940b4d33e1c358c3c41879370d25ed25216" 50001
```

Take note of the contract address that is provided to you.

## Run the Oracle
```
go run $GOPATH/src/github.com/unification-com/haiku-core/oracle/oracle.go "https://mainnet.infura.io" "http://localhost:8545" "07a4dc42e81b0e7c4e2e639741983940b4d33e1c358c3c41879370d25ed25216" "0xDCaA813d0537f59b211F6E5374fD1D85907b4B31" 600 50001
```
