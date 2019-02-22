package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/ethclient"

	store "./contracts"
)

const genesis_hash = "GENESIS_STRING"

var (
	GenesisHash        = [32]byte{}
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	downstream := argsWithoutProg[0]
	upload_key := argsWithoutProg[1]

	workchainChainId, err := strconv.ParseUint(argsWithoutProg[2], 10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	client, err := ethclient.Dial(downstream)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection established")

	privateKey, err := crypto.HexToECDSA(upload_key)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice

	fmt.Printf("The chosen gas price is: %s wei\n", gasPrice)

	evs := make([]common.Address, 0, 1)
	evs = append(evs, fromAddress)

	copy(GenesisHash[:], []byte(genesis_hash))
	address, tx, _, err := store.DeployStore(auth, client, workchainChainId, GenesisHash, evs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contract address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

}
