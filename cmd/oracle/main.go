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
		MainchainJsonRpcFlag,
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
		WRKChainJsonRPCFlag,
		WriteFrequencyFlag,
		RecordParentHashFlag,
		RecordReceiptRootFlag,
		RecordTxRootFlag,
		RecordStateRootFlag,
	}
)

func init() {

	app.Action = oracle
	app.Name = "oracle"
	app.Author = "Unification Foundation"
	app.Email = "hello@unification.com"
	app.Usage = "WRKChain Oracle"
	app.Version = Version
	app.Copyright = "Copyright (c) 2019 Unification Foundation"
	app.Commands = []cli.Command{
		initCommand,
		regCommand,
		recCommand,
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

func oracle(ctx *cli.Context) error {
	fmt.Println("RUN FOR THE HILLS!")
	fmt.Println("Or just run 'oracle help'")
	return nil
}
