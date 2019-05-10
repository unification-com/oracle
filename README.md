# WRKChain Oracle

WRKChain Oracle allows WRKChain developers to register their WRKChains and submit
block header hashes to the Unification Mainchain.

## Installation

1. Install Golang - https://golang.org/doc/install

2. Download the dependencies:

```bash
go get gopkg.in/urfave/cli.v1
go get github.com/unification-com/mainchain
```

3. Grab the WRKChain Oracle

```bash
go get github.com/unification-com/oracle
go install github.com/unification-com/oracle/cmd/wrkoracle
```

4. Check it installed OK:

```bash
wrkoracle --version
```

Should output something like

```
[~]$ wrkoracle version 0.3.0-alpha
```

You can run:

```bash
wrkoracle help
```

to output the tool's commands and flags.

## Using WRKChain Oracle

### Initialising the Oracle with the `init` command

Since the WRKChain Oracle writes to the WRKChain Root smart contract on Mainchain, 
it requires a wallet funded with UND to run. It is recommended that the Oracle
use a dedicated wallet.

The `init` command takes a private key and password, and creates an encrypted wallet
file which the Oracle can use. This only needs to be run once per Oracle instance.

1. create a data directory:

```bash
mkdir ~/.wrkchain_oracle
```

2. Use a text editor such as `nano` to create and save the private key and
password files:

```bash
nano ~/.wrkchain_oracle/.password
nano ~/.wrkchain_oracle/.pkey
```

The files should only contain the selected password and private key respectively

3. Initialise the Oracle

```bash
wrkoracle init --password ~/.wrkchain_oracle/.password --key ~/.wrkchain_oracle/.pkey 
```

4. Delete the private key file - it's no longer required

```bash
rm -f ~/.wrkchain_oracle/.pkey
```

#### Available Flags

`--datadir`: _(optional)_ Optional flag specifying the path to store the wallet file, if different from `~/.wrkchain_oracle`  
`--key`: _(required)_ Path to the file containing the private key  
`--password`: _(required)_ Path to the file containing the password

### Registering your WRKChain with the `register` command

Before any WRKChain header hashes can be recorded, it requires registering with the
WRKChain Root smart contract on Mainchain.

In order to register, the wallet which will be used to run the Oracle will need
to send a small UND deposit along with the registration command. This deposit
will be refunded to the wallet after a number of WRKChain header hashes have been
recorded.

You will need your WRKChain's `genesis` JSON, along with a list of addresses
which will be authorised to write to the WRKChain Root smart contract.

Run:
```bash
wrkoracle register --password ~/.wrkchain_oracle/.password --account [oracle_wallet_address] --genesis [/path/to/wrkchain.genesis.json] --auth [auth_address1,auth_address2] --mainchain.rpc "http://[mainchain-rpc-url]:[port]"
```

for example:

```bash
wrkoracle register --password ~/.und_mainchain/.password --account 0x160b51e66e51327ac31c643f7675b8a9006aee1e --genesis ./test/wrkchain.genesis.test.json --auth 0x160B51e66e51327ac31C643f7675B8A9006aEE1E,0xbEc4127468c51fF89719DBcA5DC57F39C0049f06 --mainchain.rpc "http://67.231.18.141:8101"
```

You should see output similer to the following:

```
Registering WRKChain with:
WRKChain Genesis Hash: 0x37bc40d2d3ee49bf688a53010983b433f54d0d2f84d5fcb939de4cd5eaf635e2
WRKChain Network ID: 123456
Adding default authorised address: 0x160B51e66e51327ac31C643f7675B8A9006aEE1E
Adding authorised address: 0xbEc4127468c51fF89719DBcA5DC57F39C0049f06
auth.From: 0x160B51e66e51327ac31C643f7675B8A9006aEE1E
Connecting to Mainchain JSON RPC on http://172.25.0.5:8101
Balance for 0x160b51e66e51327ac31c643f7675b8a9006aee1e 5000000000000000000
NonceAt: 0
RegisterWrkChain tx sent: 0x49cee85afba7838e9cf1f8cd464eb7a0d530eaaec90b6f694903d3b4cd8e4d5d
```

#### Available Flags
			
`--account`: _(required)_ Wallet Address the WRKChain Oracle will use to register 
the WRKChain  
`--auth`: _(required)_ Comma separated list of wallet addresses authorised to write 
to the WRKChain Root smart contract  
`--datadir`: _(optional)_ Optional flag specifying the path to store the wallet file, 
if different from `~/.wrkchain_oracle`  
`--genesis`: _(required)_ Path to the `genesis` JSON file  
`--mainchain.rpc`: _(optional)_ HTTP endpoint for Mainchain's JSON RPC  
`--password`: _(required)_ Path to the file containing the password  

### Recording WRKChain header hashes with the `record` command

Once the WRKChain is registered, any of the authorised wallets will be able to run
an Oracle to record hashes to Mainchain. The WRKChain Oracle will begin recording
the hashes from the latest WRKChain block.

To begin recording, run:

```bash
wrkoracle record --password ~/.und_mainchain/.password --account [oracle_wallet_address] --mainchain.rpc "http://[mainchain-rpc-url]:[port]" --wrkchain.rpc "http://[wrkchain-rpc-url]:[port]" [--hash.parent] [--hash.receipt] [--hash.tx] [--hash.state] --freq [seconds]
```

For example:

```bash
wrkoracle record --password ~/.und_mainchain/.password --account 0x160b51e66e51327ac31c643f7675b8a9006aee1e --mainchain.rpc "http://67.231.18.141:8101" --wrkchain.rpc "http://172.25.0.5:8101" --hash.parent --hash.receipt --hash.tx --hash.state --freq 60
```

You should begin to see output similar to:

```
auth.From: 0x160B51e66e51327ac31C643f7675B8A9006aEE1E
Connecting to Mainchain JSON RPC on http://172.25.0.5:8101
Connecting to WRKChain JSON RPC on http://172.25.0.5:8101
Start Polling
Balance for 0x160B51e66e51327ac31C643f7675B8A9006aEE1E 3999999999513847500
WRKChain Network ID: 50009
blockHeight 560
blockHash 0x69876f4b3646ecb0db218739b615b6bd3b2b7a00275359bd744b72090a02ab6b
parentHash 0x1c525169aef487de9a8ca09c5d5ade6772dcfa87a8162b47fc9168c0ab71084e
receiptHash 0x4ccd16c5d4907439955938f9393649aa42c9da116ae1e710fb2ea9aa55d378e5
txHash 0x2f52ff57360e18f8eff77168f4baeb08fc9213572ea85c6a9947bd47237edc70
rootHash 0xc709802cfaa1798395557b8565d0a413d7c6195c07c3bf3f8fa901ab6f6ee0bf
sealer 0x160B51e66e51327ac31C643f7675B8A9006aEE1E
Sending Tx to WRKChain Root on Mainchain
RecordHeader tx sent: 0xf02c0544eb71641bb4df611f5b486578c4f468a994b0332a3ab6b1d24e5d7fff
Waiting for 10 seconds
-------------------------------------
Balance for 0x160B51e66e51327ac31C643f7675B8A9006aEE1E 3999999999314450000
WRKChain Network ID: 50009
blockHeight 565
blockHash 0x8956f2cb47adcc19c66aad99e039959cd92f39e9aa0476436dab02803b432947
parentHash 0xc359694b2c11bb0c0f634ecbfd7c8c8396f8a5e0e25db6c4aa9c6baf8470e151
receiptHash 0xc88dacb820e3708b556f9f4c960a791058710ee8d58bb58350c53f7710f39fcf
txHash 0xdb90badc86f710ab76acabacaf14252a485c6d10376097497d15d924dda82abc
rootHash 0x9833f809e4220136af87d43b3efb03efa753beabb67684febebc369ff9a124db
sealer 0x160B51e66e51327ac31C643f7675B8A9006aEE1E
Sending Tx to WRKChain Root on Mainchain
RecordHeader tx sent: 0xdff523357b30f2f2b88c0fa2aca26d99ba30ac469fa5924fb72ad3529a45a279
Waiting for 10 seconds
-------------------------------------
```

#### Available Flags

`--account`: _(required)_ Wallet Address the WRKChain Oracle will use to register 
the WRKChain  
`--datadir`: _(optional)_ Optional flag specifying the path to store the wallet file, 
if different from `~/.wrkchain_oracle`  
`--freq`: _(optional)_ Frequency the WRKChain Oracle should write hashes to Mainchain, in seconds  
`--hash.parent`: _(optional)_ If set, the block's Parent Hash will also be recorded  
`--hash.receipt`: _(optional)_ If set, the block's Receipt Merkle Root Hash will also be recorded  
`--hash.state`: _(optional)_ If set, the block's State Merkle Root Hash will also be recorded  
`--hash.tx`: _(optional)_ If set, the block's Tx Merkle Root Hash will also be recorded  
`--mainchain.rpc`: _(optional)_ HTTP endpoint for Mainchain's JSON RPC  
`--password`: _(required)_ Path to the file containing the password  
`--wrkchain.rpc`: _(required)_ HTTP endpoint for *your WRKChain's* JSON RPC  
