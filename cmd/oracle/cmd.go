package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/unification-com/mainchain/accounts"
	"github.com/unification-com/mainchain/accounts/abi/bind"

	//"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/accounts/keystore"
	"github.com/unification-com/mainchain/common"
	wrkchainroot "github.com/unification-com/mainchain/contracts/wrkchainroot"
	"github.com/unification-com/mainchain/core"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/ethclient"
	"gopkg.in/urfave/cli.v1"
	"io"
	"io/ioutil"
	//"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	WRKChainRootContractAddress = "0x0000000000000000000000000000000000000087"
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

	genesisHash := block.Hash().Hex()
	wrkchainNetworkId := genesis.Config.ChainId

	fmt.Println("genesisHash:", genesisHash)
	fmt.Println("wrkchainNetworkId:", wrkchainNetworkId)

	// Process authorised addresses
	if ! ctx.IsSet(AuthorisedAccountsFlag.Name) {
		Fatalf("List of Authorised addresses required required")
	}

	// add this account by default
	authAddresses := []common.Address{common.HexToAddress(strings.TrimSpace(ctx.String(AccountUnlockFlag.Name)))}

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

	// make connection to Mainchain
	wrkchainRootContract := initialiseConnnection(ctx)

	wrkChainList, err := wrkchainRootContract.WrkchainList(wrkchainNetworkId)

	if err != nil {
		Fatalf("Couldn't query WrkchainList", "err", err)
	}

	reggedGenHash := string(wrkChainList.GenesisHash[:])

	fmt.Printf("wrkChainList.GenesisHash: %v", reggedGenHash)

	return nil
}

func initialiseConnnection(ctx *cli.Context) *wrkchainroot.WrkchainRoot {

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

	fh, err := os.Open(thisAcc.URL.Path)
	if err != nil {
		Fatalf("Couldn't open Keystore", "err", err)
	}
	defer fh.Close()

	var reader io.Reader = fh
	txOpts, err := bind.NewTransactor(reader, pass)
	if err != nil {
		Fatalf("Couldn't bind transactor", "err", err)
	}

	// Todo: take RPC from flags
	mainchainClient, err := ethclient.Dial(DefaultMainchainTestnetRPC)
	if err != nil {
		Fatalf("Couldn't connect to Mainchain", "err", err)
	}

	balance, _ := mainchainClient.BalanceAt(context.Background(), common.HexToAddress(account), nil)
	fmt.Printf("Balance: %d\n",balance)

	// Todo - check balance is sufficient for registering

	wrkchainRootContract, err := wrkchainroot.NewWrkchainRoot(txOpts, common.HexToAddress(WRKChainRootContractAddress), mainchainClient)

	if err != nil {
		Fatalf("Couldn't bind WRKChainRoot contract", "err", err)
	}

	return wrkchainRootContract
}


// DefaultDataDir is the default data directory to use for the databases and other
// persistence requirements.
func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "WrkchainOracle")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "WrkchainOracle")
		} else {
			return filepath.Join(home, ".wrkchain_oracle")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

// Fatalf formats a message to standard error and exits the program.
// The message is also printed to standard output if standard error
// is redirected to a different file.
func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}

func MkDataDir(dirPath string) {
	if err := os.MkdirAll(dirPath, 0700); err != nil {
		Fatalf("Could not create datadir", "datadir", dirPath, "err", err)
	}
}
