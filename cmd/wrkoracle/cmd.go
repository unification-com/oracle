package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/unification-com/mainchain/accounts"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/accounts/keystore"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/common/hexutil"
	wrkchainroot "github.com/unification-com/mainchain/contracts/wrkchainroot/contract"
	"github.com/unification-com/mainchain/core"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/ethclient"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
WRKChainRootContractAddress: Contract address, hard-coded into Mainchain's genesis block
DepositStorageAddress: Storage address in WRKChain Root contract contaiing the required UND deposit amount, in wei
DefaultMainchainTestnetRPC: Default UND Mainchain JSON RPC URL for testnet
DefaultMainchainMainnetRPC: Default UND Mainchain JSON RPC URL for maainnet
*/
const (
	WRKChainRootContractAddress = "0x0000000000000000000000000000000000000087"
	DepositStorageAddress       = "0x0000000000000000000000000000000000000000000000000000000000000000"
	DefaultMainchainTestnetRPC  = "https://rpc-testnet.unification.io"
	DefaultMainchainMainnetRPC  = "https://rpc-testnet.unification.io"
	WRKChainRootTax = 1
)

var (
	initCommand = cli.Command{
		Action:    initOracle,
		Name:      "init",
		Usage:     "Initialise the Oracle",
		ArgsUsage: "",
		Flags: []cli.Flag{
			PasswordPathFlag,
			PrivateKeyPathFlag,
			DataDirectoryFlag,
		},
		Category: "ORACLE COMMANDS",
		Description: `
The init command initialises the Oracle, creating a secure wallet for running.`,
	}

	registerCommand = cli.Command{
		Action:    registerWrkchain,
		Name:      "register",
		Usage:     "Register a WRKChain",
		ArgsUsage: "",
		Flags: []cli.Flag{
			AccountUnlockFlag,
			PasswordPathFlag,
			DataDirectoryFlag,
			GenesisPathFlag,
			AuthorisedAccountsFlag,
			MainchainJSONRPCFlag,
			UndTestnetFlag,
		},
		Category: "ORACLE COMMANDS",
		Description: `
The register command registers a new WRKChain on the UND Mainchain`,
	}

	recordCommand = cli.Command{
		Action:    recordWrkchainBlock,
		Name:      "record",
		Usage:     "Record WRKChain Block header hashes",
		ArgsUsage: "",
		Flags: []cli.Flag{
			AccountUnlockFlag,
			PasswordPathFlag,
			DataDirectoryFlag,
			MainchainJSONRPCFlag,
			UndTestnetFlag,
			WRKChainJSONRPCFlag,
			WriteFrequencyFlag,
			RecordParentHashFlag,
			RecordReceiptRootFlag,
			RecordTxRootFlag,
			RecordStateRootFlag,
		},
		Category: "ORACLE COMMANDS",
		Description: `
The record command runs the WRKChain Block Heaader Hash recorder and submits WRKChain hashes to Mainchain.
A WRKChain requires registering first, with the register command`,
	}

	lastPendingNonce uint64 = 0
)

func initOracle(ctx *cli.Context) error {

	MkDataDir(ctx.String(DataDirectoryFlag.Name))

	// Grab the password
	if !ctx.IsSet(PasswordPathFlag.Name) {
		Fatalf("Path to password file required")
	}

	blob, err := ioutil.ReadFile(ctx.String(PasswordPathFlag.Name))

	if err != nil {
		Fatalf("Failed to read account password contents", "file", ctx.String(PasswordPathFlag.Name), "err", err)
	}
	pass := strings.TrimSpace(string(blob))

	// Grab the private key
	if !ctx.IsSet(PrivateKeyPathFlag.Name) {
		Fatalf("Path to private key file required")
	}

	blob, err = ioutil.ReadFile(ctx.String(PrivateKeyPathFlag.Name))

	if err != nil {
		Fatalf("Failed to read private key contents", "file", ctx.String(PrivateKeyPathFlag.Name), "err", err)
	}

	pkey := strings.TrimSpace(string(blob))

	privateKey, err := crypto.HexToECDSA(pkey)

	if err != nil {
		Fatalf("Failed to convert pkey", "err", err)
	}

	// Create a keystore for the account
	ks := keystore.NewKeyStore(filepath.Join(ctx.String(DataDirectoryFlag.Name), "keys"), keystore.StandardScryptN, keystore.StandardScryptP)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		Fatalf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	account := crypto.PubkeyToAddress(*publicKeyECDSA)

	if !ks.HasAddress(account) {
		_, err = ks.ImportECDSA(privateKey, pass)
		if err != nil {
			Fatalf("Failed to import Oracle signer account", "err", err)
		}

		fmt.Printf("Account %v created. You can now delete %v\n", account.Hex(), ctx.String(PrivateKeyPathFlag.Name))

	} else {
		fmt.Printf("Account %v already exists\n", account.Hex())
	}

	return nil
}

func registerWrkchain(ctx *cli.Context) error {

	fmt.Println()
	ctxBg := context.Background()
	MkDataDir(ctx.String(DataDirectoryFlag.Name))

	// Process Genesis. Note - only geth based genesis blocks supported at this time
	if !ctx.IsSet(GenesisPathFlag.Name) {
		Fatalf("Path to genesis JSON file required")
	}
	file, err := os.Open(strings.TrimSpace(ctx.String(GenesisPathFlag.Name)))

	if err != nil {
		Fatalf("Failed to read genesis file: %v", err)
	}

	defer file.Close()

	genesis := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		Fatalf("invalid genesis file: %v", err)
	}

	block := genesis.ToBlock(nil)

	genesisHash := block.Header().GoEthereumHash()
	wrkchainNetworkID := genesis.Config.ChainId

	fmt.Println("Registering WRKChain with:")
	fmt.Println("WRKChain Genesis Hash:", genesisHash.Hex())
	fmt.Println("WRKChain Network ID:", wrkchainNetworkID)

	// Process authorised addresses
	if !ctx.IsSet(AuthorisedAccountsFlag.Name) {
		Fatalf("List of Authorised addresses required required")
	}

	thisAccount := common.HexToAddress(strings.TrimSpace(ctx.String(AccountUnlockFlag.Name)))

	// add this account by default
	fmt.Println("Adding default authorised address:", thisAccount.Hex())
	authAddresses := []common.Address{thisAccount}

	addressParts := strings.Split(strings.TrimSpace(ctx.String(AuthorisedAccountsFlag.Name)), ",")

	for _, authAddr := range addressParts {
		authAddr = strings.TrimSpace(authAddr)
		if !common.IsHexAddress(authAddr) {
			Fatalf("Invalid address", authAddr)
		}
		authAddr := common.HexToAddress(authAddr)
		if authAddr != thisAccount {
			fmt.Println("Adding authorised address:", authAddr.Hex())
			authAddresses = append(authAddresses, authAddr)
		}
	}

	if len(authAddresses) == 0 {
		Fatalf("At least one valid authorised address required")
	}

	// Create a new WRKChainRoot Session
	wrkchainRootSession := NewWrkchainRootSession(ctxBg, ctx)

	// Connect
	fmt.Println("Connecting to Mainchain JSON RPC on", ctx.String(MainchainJSONRPCFlag.Name))
	mainchainClient, err := ethclient.Dial(strings.TrimSpace(ctx.String(MainchainJSONRPCFlag.Name)))
	if err != nil {
		Fatalf("Couldn't connect to Mainchain", "err", err)
	}

	balance, _ := mainchainClient.BalanceAt(ctxBg, thisAccount, nil)
	balanceFloat := new(big.Float)
	balanceFloat.SetString(balance.String())
	undValue := new(big.Float).Quo(balanceFloat, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance for", ctx.String(AccountUnlockFlag.Name), undValue, "UND")

	wrkchainRootSession = LoadContract(wrkchainRootSession, mainchainClient)

	// Query RegisterWrkChain event to see if WRKChain has already been registered
	var filterOpts = new(bind.FilterOpts)
	filterOpts.Start = 0
	filterOpts.End = nil
	filterOpts.Context = ctxBg

	wrkchainIDFilterList := make([]*big.Int, 0)
	wrkchainIDFilterList = append(wrkchainIDFilterList, wrkchainNetworkID)

	registerWrkChainEvents, err := wrkchainRootSession.Contract.FilterRegisterWrkChain(filterOpts, wrkchainIDFilterList)
	if err != nil {
		Fatalf("failed to filter for RegisterWrkChain events:", "err", err)
	}

	defer registerWrkChainEvents.Close()

	if registerWrkChainEvents.Next() {
		// already registered. Output info and exit
		fmt.Println("Found WRKChain ID:", registerWrkChainEvents.Event.ChainId.String())
		fmt.Println("with Genesis Hash", hexutil.Encode(registerWrkChainEvents.Event.GenesisHash[:]))
		Fatalf("WRKChain already registered in Tx", registerWrkChainEvents.Event.Raw.TxHash.Hex())
	}

	// gather up params for registering WRKChain
	// Required deposit amount held in first storage value in WRKChain Root contract
	deposit, _ := mainchainClient.StorageAt(ctxBg, common.HexToAddress(WRKChainRootContractAddress), common.HexToHash(DepositStorageAddress), nil)
	depositAmount := big.NewInt(0).SetBytes(deposit)

	fmt.Printf("depositAmount = %s\n", depositAmount.String())

	totalAmount := big.NewInt(0)
	totalAmount.Add(depositAmount, calcTax())

	if balance.Cmp(totalAmount) == -1 {
		Fatalf("Not enough UND to register",
			"account",
			ctx.String(AccountUnlockFlag.Name),
			"balance",
			undValue.String(),
			"UND",
		)
	}

	nonce, _ := mainchainClient.PendingNonceAt(ctxBg, thisAccount)
	fmt.Printf("PendingNonceAt: %v\n", nonce)

	wrkchainRootSession.TransactOpts.Value = depositAmount
	wrkchainRootSession.TransactOpts.Nonce = big.NewInt(int64(nonce))
	wrkchainRootSession.TransactOpts.GasLimit = 240000 // pseudo gas limit. Never consumed, but used to calculate block gas consumption

	tx, err := wrkchainRootSession.RegisterWrkChain(wrkchainNetworkID, authAddresses, genesisHash)

	if err != nil {
		Fatalf("Couldn't register WRKChain", "err", err)
	}

	fmt.Println("RegisterWrkChain tx sent:", tx.Hash().Hex())

	return nil
}

func recordWrkchainBlock(ctx *cli.Context) error {

	fmt.Println()
	ctxBg := context.Background()
	MkDataDir(ctx.String(DataDirectoryFlag.Name))

	if !ctx.IsSet(WRKChainJSONRPCFlag.Name) {
		Fatalf("WRKChainJSONRPCFlag not set")
	}

	// Create a new WRKChainRoot Session
	wrkchainRootSession := NewWrkchainRootSession(ctxBg, ctx)

	thisAccount := common.HexToAddress(strings.TrimSpace(ctx.String(AccountUnlockFlag.Name)))

	// Connect
	fmt.Println("Connecting to Mainchain JSON RPC on", ctx.String(MainchainJSONRPCFlag.Name))
	mainchainClient, err := ethclient.Dial(strings.TrimSpace(ctx.String(MainchainJSONRPCFlag.Name)))
	if err != nil {
		Fatalf("Couldn't connect to Mainchain", "err", err)
	}

	wrkchainRootSession = LoadContract(wrkchainRootSession, mainchainClient)

	fmt.Println("Connecting to WRKChain JSON RPC on", ctx.String(WRKChainJSONRPCFlag.Name))
	wrkChainClient, _ := ethclient.Dial(strings.TrimSpace(ctx.String(WRKChainJSONRPCFlag.Name)))

	wrkchainNetworkID, err := wrkChainClient.NetworkID(ctxBg)

	if err != nil {
		Fatalf("Could not get WRKChain Network ID: ", err)
	}

	pollWrkchain(ctx, mainchainClient, &wrkchainRootSession, wrkChainClient, wrkchainNetworkID, thisAccount)

	return nil
}

func pollWrkchain(
	ctx *cli.Context,
	mainchainClient *ethclient.Client,
	wrkchainRootSession *wrkchainroot.WRKChainRootSession,
	wrkChainClient *ethclient.Client,
	wrkchainNetworkID *big.Int,
	thisAccount common.Address,
) {

	fmt.Println("Start Polling")

	frequency := ctx.Int64(WriteFrequencyFlag.Name)

	for {

		// get UND Balance
		balance, _ := mainchainClient.BalanceAt(context.Background(), thisAccount, nil)

		balanceFloat := new(big.Float)
		balanceFloat.SetString(balance.String())
		undValue := new(big.Float).Quo(balanceFloat, big.NewFloat(math.Pow10(18)))
		fmt.Println("Balance for", thisAccount.Hex(), undValue, "UND")

		if balance.Cmp(calcTax()) == -1 {
			Fatalf("Not enough UND to record",
				"account",
				ctx.String(AccountUnlockFlag.Name),
				"balance",
				undValue.String(),
				"UND",
			)
		}

		nonce, _ := mainchainClient.PendingNonceAt(context.Background(), thisAccount)
		fmt.Printf("PendingNonceAt: %v\n", nonce)
		fmt.Printf("lastPendingNonce: %v\n", lastPendingNonce)
		if nonce == lastPendingNonce {
			nonce++
		}
		lastPendingNonce = nonce

		latestWrkchainHeader, err := wrkChainClient.HeaderByNumber(context.Background(), nil)

		if err != nil {
			Fatalf("Could not get latest WRKChain Block: ", err)
		}

		blockHash := latestWrkchainHeader.GoEthereumHash()
		parentHash := [32]byte{0}
		receiptHash := [32]byte{0}
		txHash := [32]byte{0}
		rootHash := [32]byte{0}
		blockHeight := latestWrkchainHeader.Number

		if ctx.IsSet(RecordParentHashFlag.Name) {
			parentHash = latestWrkchainHeader.ParentHash
		}

		if ctx.IsSet(RecordReceiptRootFlag.Name) {
			receiptHash = latestWrkchainHeader.ReceiptHash
		}

		if ctx.IsSet(RecordTxRootFlag.Name) {
			txHash = latestWrkchainHeader.TxHash
		}

		if ctx.IsSet(RecordStateRootFlag.Name) {
			rootHash = latestWrkchainHeader.Root
		}
		go record(
			wrkchainRootSession,
			wrkchainNetworkID,
			blockHeight,
			blockHash,
			parentHash,
			receiptHash,
			txHash,
			rootHash,
			thisAccount,
			frequency,
			nonce)

		<-time.After(time.Duration(frequency) * time.Second)
	}

}

func record(
	wrkchainRootSession *wrkchainroot.WRKChainRootSession,
	wrkchainNetworkID *big.Int,
	blockHeight *big.Int,
	blockHash [32]byte,
	parentHash [32]byte,
	receiptHash [32]byte,
	txHash [32]byte,
	rootHash [32]byte,
	sealer common.Address,
	frequency int64,
	nonce uint64) {

	fmt.Println("WRKChain Network ID:", wrkchainNetworkID)
	fmt.Println("blockHeight", blockHeight)
	fmt.Println("blockHash", common.ToHex(blockHash[:]))
	fmt.Println("parentHash", common.ToHex(parentHash[:]))
	fmt.Println("receiptHash", common.ToHex(receiptHash[:]))
	fmt.Println("txHash", common.ToHex(txHash[:]))
	fmt.Println("rootHash", common.ToHex(rootHash[:]))
	fmt.Println("sealer", sealer.Hex())
	fmt.Println("nonce", nonce)

	fmt.Println("Sending Tx to WRKChain Root on Mainchain")

	wrkchainRootSession.TransactOpts.Value = big.NewInt(0)
	wrkchainRootSession.TransactOpts.Nonce = big.NewInt(int64(nonce))
	wrkchainRootSession.TransactOpts.GasLimit = 240000 // pseudo gas limit. Never consumed, but used to calculate block gas consumption

	tx, err := wrkchainRootSession.RecordHeader(wrkchainNetworkID, blockHeight, blockHash, parentHash, receiptHash, txHash, rootHash, sealer)

	if err != nil {
		Fatalf("Could not record WRKChain Header:", err)
	}

	fmt.Println("RecordHeader tx sent:", tx.Hash().Hex())

	// ToDo: Check tx receipt for success/failure and report

	fmt.Println("Waiting for", frequency, "seconds")
	fmt.Println("-------------------------------------")

}

// NewWrkchainRootSession Create a new session for the WRKChain Root smart contract
func NewWrkchainRootSession(bgCtx context.Context, ctx *cli.Context) (session wrkchainroot.WRKChainRootSession) {
	// Grab the password
	if !ctx.IsSet(PasswordPathFlag.Name) {
		Fatalf("Path to password file required")
	}

	blob, err := ioutil.ReadFile(ctx.String(PasswordPathFlag.Name))

	if err != nil {
		Fatalf("Failed to read account password contents", "file", ctx.String(PasswordPathFlag.Name), "err", err)
	}
	pass := strings.TrimSpace(string(blob))

	// grab account to unlock
	if !ctx.IsSet(AccountUnlockFlag.Name) {
		Fatalf("Account to unlock required")
	}

	account := strings.TrimSpace(ctx.String(AccountUnlockFlag.Name))
	if !common.IsHexAddress(account) {
		Fatalf("Account not in common hex format, e.g. 0xabd123...")
	}
	acc := accounts.Account{Address: common.HexToAddress(account)}

	ks := keystore.NewKeyStore(filepath.Join(ctx.String(DataDirectoryFlag.Name), "keys"), keystore.StandardScryptN, keystore.StandardScryptP)

	thisAcc, err := ks.Find(acc)
	if err != nil {
		Fatalf("Could not find account. Did you init first?: ", "err", err)
	}

	keystore, err := os.Open(thisAcc.URL.Path)
	if err != nil {
		Fatalf("Couldn't read Keystore", "err", err)
	}
	defer keystore.Close()

	auth, err := bind.NewTransactor(keystore, pass)
	if err != nil {
		Fatalf("Couldn't bind transactor", "err", err)
	}

	fmt.Printf("auth.From: %v\n", auth.From.Hex())

	return wrkchainroot.WRKChainRootSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    auth.From,
			Context: bgCtx,
		},
	}

}

// LoadContract Load the WRKChain Root smart contract into the WRKChain Root Session
func LoadContract(session wrkchainroot.WRKChainRootSession, client *ethclient.Client) wrkchainroot.WRKChainRootSession {
	addr := common.HexToAddress(WRKChainRootContractAddress)
	instance, err := wrkchainroot.NewWRKChainRoot(addr, client)
	if err != nil {
		Fatalf("could not load contract: %v\n", err)
	}
	session.Contract = instance
	return session
}

func calcTax() *big.Int {
	tax := big.NewInt(WRKChainRootTax)
	tax.Mul(tax, big.NewInt(1e18))
	return tax
}
