package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"sort"
)

var (
	app = cli.NewApp()

	commonFlags = []cli.Flag{
		DataDirectoryFlag,
		UndTestnetFlag,
		MainchainJSONRPCFlag,
	}

	regFlags = []cli.Flag{
		GenesisPathFlag,
		AuthorisedAccountsFlag,
	}

	accFlags = []cli.Flag{
		PasswordPathFlag,
		PrivateKeyPathFlag,
		AccountUnlockFlag,
	}

	wrkchainFlags = []cli.Flag{
		WRKChainJSONRPCFlag,
		WriteFrequencyFlag,
		RecordParentHashFlag,
		RecordReceiptRootFlag,
		RecordTxRootFlag,
		RecordStateRootFlag,
	}
)

func init() {

	app.Action = wrkoracle
	app.Name = "wrkoracle"
	app.Author = "Unification Foundation"
	app.Email = "hello@unification.com"
	app.Usage = "WRKChain Oracle - registers a WRKCHain on UND Mainchain, and records a WRKChain's hashes on Mainchain"
	app.Version = Version
	app.Copyright = "Copyright (c) 2019 Unification Foundation"
	app.Commands = []cli.Command{
		initCommand,
		registerCommand,
		recordCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = append(app.Flags, commonFlags...)
	app.Flags = append(app.Flags, regFlags...)
	app.Flags = append(app.Flags, accFlags...)
	app.Flags = append(app.Flags, wrkchainFlags...)

	app.After = func(ctx *cli.Context) error {
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func wrkoracle(ctx *cli.Context) error {
	fmt.Println("RUN FOR THE HILLS!")
	fmt.Println("Or just run 'wrkoracle help'")
	return nil
}
