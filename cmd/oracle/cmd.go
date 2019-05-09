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
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

const (
	WRKChainRootContractAddress = "0x0000000000000000000000000000000000000087"
	DepositStorageAddress       = "0x0000000000000000000000000000000000000000000000000000000000000000"
	DefaultMainchainTestnetRPC  = "http://67.231.18.141:8101"
	DefaultMainchainMainnetRPC  = "http://67.231.18.141:8101"
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

	regCommand = cli.Command{
		Action:    regWrkchain,
		Name:      "register",
		Usage:     "Register a WRKChain",
		ArgsUsage: "",
		Flags: []cli.Flag{
			AccountUnlockFlag,
			PasswordPathFlag,
			DataDirectoryFlag,
			GenesisPathFlag,
			AuthorisedAccountsFlag,
			MainchainJsonRpcFlag,
			UndTestnetFlag,
		},
		Category: "ORACLE COMMANDS",
		Description: `
The init command initialises the Oracle, creating a secure wallet for running.`,
	}
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
			Fatalf("Failed to import oracle signer account", "err", err)
		}

		fmt.Printf("Account %v created. You can now delete %v\n", account.Hex(), ctx.String(PrivateKeyPathFlag.Name))

	} else {
		fmt.Printf("Account %v already exists\n", account.Hex())
	}

	return nil
}

func regWrkchain(ctx *cli.Context) error {

	fmt.Println()
	ctxBg := context.Background()
	MkDataDir(ctx.String(DataDirectoryFlag.Name))

	// Process Genesis. Note - only geth based genesis blocks supported at this time
	if ! ctx.IsSet(GenesisPathFlag.Name) {
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

	genesisHash := block.Hash()
	wrkchainNetworkId := genesis.Config.ChainId

	fmt.Println("Registering WRKChain with:")
	fmt.Println("WRKChain Genesis Hash:", genesisHash.Hex())
	fmt.Println("WRKChain Network ID:", wrkchainNetworkId)

	// Process authorised addresses
	if ! ctx.IsSet(AuthorisedAccountsFlag.Name) {
		Fatalf("List of Authorised addresses required required")
	}

	thisAccount := common.HexToAddress(strings.TrimSpace(ctx.String(AccountUnlockFlag.Name)))

	// add this account by default
	authAddresses := []common.Address{thisAccount}

	addressParts := strings.Split(strings.TrimSpace(ctx.String(AuthorisedAccountsFlag.Name)), ",")

	for _, authAddr := range addressParts {
		authAddr = strings.TrimSpace(authAddr)
		if !common.IsHexAddress(authAddr) {
			Fatalf("Invalid address", authAddr)
		}
		authAddresses = append(authAddresses, common.StringToAddress(authAddr))
	}

	if len(authAddresses) == 0 {
		Fatalf("At least one valid authorised address required")
	}

	// Create a new WRKChainRoot Session
	wrkchainRootSession := NewWrkchainRootSession(ctx, ctxBg)

	// Connect
	fmt.Println("Connecting to Mainchain JSON RPC on", ctx.String(MainchainJsonRpcFlag.Name))
	mainchainClient, err := ethclient.Dial(strings.TrimSpace(ctx.String(MainchainJsonRpcFlag.Name)))
	if err != nil {
		Fatalf("Couldn't connect to Mainchain", "err", err)
	}

	balance, _ := mainchainClient.BalanceAt(ctxBg, thisAccount, nil)
	// ToDo: output balance in UND
	fmt.Println("Balance for", ctx.String(AccountUnlockFlag.Name), balance)

	wrkchainRootSession = LoadContract(wrkchainRootSession, mainchainClient)

	// Query RegisterWrkChain event to see if WRKChain has already been registered
	var filterOpts = new(bind.FilterOpts)
	filterOpts.Start = 0
	filterOpts.End = nil
	filterOpts.Context = ctxBg

	wrkchainIdFilterList := make([]*big.Int, 0)
	wrkchainIdFilterList = append(wrkchainIdFilterList, wrkchainNetworkId)

	registerWrkChainEvents, err := wrkchainRootSession.Contract.FilterRegisterWrkChain(filterOpts, wrkchainIdFilterList)
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
	deposit, err := mainchainClient.StorageAt(ctxBg, common.HexToAddress(WRKChainRootContractAddress), common.HexToHash(DepositStorageAddress), nil)
	depositAmount := big.NewInt(0).SetBytes(deposit)
	// ToDo: implement ethclient.estimateGas
	approxGas := uint64(300000)

	totalAmount := big.NewInt(0)
	totalAmount.Add(depositAmount, big.NewInt(0).SetUint64(approxGas))

	if balance.Cmp(totalAmount) == -1 {
		Fatalf("Not enough UND to register",
			"account",
			ctx.String(AccountUnlockFlag.Name),
			"balance",
			balance,
		)
	}

	nonce, _ := mainchainClient.NonceAt(ctxBg, thisAccount, nil)
	fmt.Printf("NonceAt: %v\n", nonce)

	gasPrice, err := mainchainClient.SuggestGasPrice(context.Background())
	if err != nil {
		Fatalf("Couldn't get gas price", "err", err)
	}

	wrkchainRootSession.TransactOpts.Value = depositAmount
	wrkchainRootSession.TransactOpts.Nonce = big.NewInt(int64(nonce))
	wrkchainRootSession.TransactOpts.GasLimit = approxGas
	wrkchainRootSession.TransactOpts.GasPrice = gasPrice

	tx, err := wrkchainRootSession.RegisterWrkChain(wrkchainNetworkId, authAddresses, genesisHash)

	if err != nil {
		Fatalf("Couldn't register WRKChain", "err", err)
	}

	fmt.Println("RegisterWrkChain tx sent:", tx.Hash().Hex())

	return nil
}


func NewWrkchainRootSession(ctx *cli.Context, bgCtx context.Context) (session wrkchainroot.WRKChainRootSession) {
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

func LoadContract(session wrkchainroot.WRKChainRootSession, client *ethclient.Client) wrkchainroot.WRKChainRootSession {
	addr := common.HexToAddress(WRKChainRootContractAddress)
	instance, err := wrkchainroot.NewWRKChainRoot(addr, client)
	if err != nil {
		Fatalf("could not load contract: %v\n", err)
	}
	session.Contract = instance
	return session
}

