package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/ethclient"

	store "./contracts"
)

func main() {

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	upstream := argsWithoutProg[0]
	downsteam := argsWithoutProg[1]
	private_key := argsWithoutProg[2]
	contract_address := argsWithoutProg[3]

	block_time, err := strconv.Atoi(argsWithoutProg[4])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	workchainChainId, err := strconv.ParseUint(argsWithoutProg[5], 10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("upstream: %s\n", upstream)
	fmt.Printf("downsteam: %s\n", downsteam)

	client, err := ethclient.Dial(downsteam)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection established to Downstream")

	client_upstream, err := ethclient.Dial(upstream)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection established to Upstream")

	for {

		fmt.Printf("Reading block header from Workchain RPC node: %s\n", upstream)

		block, err := client_upstream.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}

		privateKey, err := crypto.HexToECDSA(private_key)
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		fmt.Println(fromAddress)

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		address := common.HexToAddress(contract_address)
		instance, err := store.NewStore(address, client)
		if err != nil {
			log.Fatal(err)
		}

		Height := block.Number().Uint64()

		fmt.Printf("Sending hashes to Mainchain contract: %s via %s\n", contract_address, downsteam)

		fmt.Printf("chain ID: %d\n", workchainChainId)
		fmt.Printf("block: %d\n", Height)
		fmt.Printf("hash: %s\n",  block.Hash().Hex())
		fmt.Printf("parent hash: %s\n",  block.ParentHash().Hex())
		fmt.Printf("receipt hash: %s\n",  block.ReceiptHash().Hex())
		fmt.Printf("tx hash: %s\n",  block.TxHash().Hex())
		fmt.Printf("state root: %s\n",  block.Root().Hex())
		fmt.Printf("sender: %s\n",  fromAddress.Hex())

		tx, err := instance.RecordHeader(auth, Height, block.Hash(), block.ParentHash(), block.ReceiptHash(), block.TxHash(), block.Root(), fromAddress, workchainChainId)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

		result, err := instance.BlockHeaders(nil, Height)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)

		fmt.Println("Sleeping")
		time.Sleep(time.Duration(block_time) * time.Second)
	}

}
