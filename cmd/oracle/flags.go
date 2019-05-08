package main

import (
	"flag"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/user"
	"path"
	"strings"
)

var (

	// Common flags
	DataDirectoryFlag = DirectoryFlag{
		Name:  "datadir",
		Usage: "Directory for the keystore and data",
		Value: DirectoryString{DefaultDataDir()},
	}
	MainchainJsonRpcFlag = cli.StringFlag{
		Name:  "mainchain.rpc",
		Usage: "Mainchain JSON RPC endpoint",
		Value: DefaultMainchainTestnetRPC,
	}
	UndTestnetFlag = cli.BoolFlag{
		Name:  "und-testnet",
		Usage: "und test network",
	}

	// Registration flags
	GenesisPathFlag = cli.StringFlag{
		Name:  "genesis",
		Usage: "Path to the WRKChain's genesis.json",
	}
	AuthorisedAccountsFlag = cli.StringFlag{
		Name:  "auth",
		Usage: "Comma separated list of addresses authorised to write to the WRKChain Root smart contract",
	}

	// Account flags
	PasswordPathFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Path to the account password",
	}
	PrivateKeyPathFlag = cli.StringFlag{
		Name:  "key",
		Usage: "Path to the private key",
	}
	AccountUnlockFlag = cli.StringFlag{
		Name:  "account",
		Usage: "Account to unlock - will be used tp write to the WRKChain Root smart contract",
	}

	// WRKChain flags
	WRKChainJsonRPCFlag  = cli.StringFlag{
		Name:  "wrkchain.rpc",
		Usage: "URI for the WRKChain's JSON RPC API, e.g. http://localhost:8101",
	}
	WriteFrequencyFlag = cli.IntFlag{
		Name:  "freq",
		Usage: "Frequency WRKChain block hashes are written, in seconds. Default 3600",
		Value: 3600,
	}
)


// Custom type which is registered in the flags library which cli uses for
// argument parsing. This allows us to expand Value to an absolute path when
// the argument is parsed
type DirectoryString struct {
	Value string
}

func (self *DirectoryString) String() string {
	return self.Value
}

func (self *DirectoryString) Set(value string) error {
	self.Value = expandPath(value)
	return nil
}

// Custom cli.Flag type which expand the received string to an absolute path.
// e.g. ~/.ethereum -> /home/username/.ethereum
type DirectoryFlag struct {
	Name  string
	Value DirectoryString
	Usage string
}

func (self DirectoryFlag) String() string {
	fmtString := "%s %v\t%v"
	if len(self.Value.Value) > 0 {
		fmtString = "%s \"%v\"\t%v"
	}
	return fmt.Sprintf(fmtString, prefixedNames(self.Name), self.Value.Value, self.Usage)
}

func (self DirectoryFlag) GetName() string {
	return self.Name
}

func (self *DirectoryFlag) Set(value string) {
	self.Value.Value = value
}


func eachName(longName string, fn func(string)) {
	parts := strings.Split(longName, ",")
	for _, name := range parts {
		name = strings.Trim(name, " ")
		fn(name)
	}
}

// called by cli library, grabs variable from environment (if in env)
// and adds variable to flag set for parsing.
func (self DirectoryFlag) Apply(set *flag.FlagSet) {
	eachName(self.Name, func(name string) {
		set.Var(&self.Value, self.Name, self.Usage)
	})
}

func prefixedNames(fullName string) (prefixed string) {
	parts := strings.Split(fullName, ",")
	for i, name := range parts {
		name = strings.Trim(name, " ")
		prefixed += prefixFor(name) + name
		if i < len(parts)-1 {
			prefixed += ", "
		}
	}
	return
}

func prefixFor(name string) (prefix string) {
	if len(name) == 1 {
		prefix = "-"
	} else {
		prefix = "--"
	}

	return
}

// Expands a file path
// 1. replace tilde with users home dir
// 2. expands embedded environment variables
// 3. cleans the path, e.g. /a/b/../c -> /a/c
// Note, it has limitations, e.g. ~someuser/tmp will not be expanded
func expandPath(p string) string {
	if strings.HasPrefix(p, "~/") || strings.HasPrefix(p, "~\\") {
		if home := homeDir(); home != "" {
			p = home + p[1:]
		}
	}
	return path.Clean(os.ExpandEnv(p))
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}