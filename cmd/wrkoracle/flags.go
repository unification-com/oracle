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

/*
CommandHelpTemplate: Template for help output
*/
var (
	CommandHelpTemplate = `{{.cmd.Name}}{{if .cmd.Subcommands}} command{{end}}{{if .cmd.Flags}} [command options]{{end}} [arguments...]
{{if .cmd.Description}}{{.cmd.Description}}
{{end}}{{if .cmd.Subcommands}}
SUBCOMMANDS:
	{{range .cmd.Subcommands}}{{.cmd.Name}}{{with .cmd.ShortName}}, {{.cmd}}{{end}}{{ "\t" }}{{.cmd.Usage}}
	{{end}}{{end}}{{if .categorizedFlags}}
{{range $idx, $categorized := .categorizedFlags}}{{$categorized.Name}} OPTIONS:
{{range $categorized.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}{{end}}`
)

func init() {
	cli.AppHelpTemplate = `{{.Name}} command{{if .Flags}} [command options]{{end}} [arguments...]

USAGE:
   {{.Usage}}

VERSION:
   {{.Version}}

AUTHOR:
   {{.Author}}
   {{.Email}}
   {{.Copyright}}

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

	cli.CommandHelpTemplate = CommandHelpTemplate
}

var (
	// Common flags

	// DataDirectoryFlag Directory for the keystore and data
	DataDirectoryFlag = DirectoryFlag{
		Name:  "datadir",
		Usage: "Directory for the keystore and data",
		Value: DirectoryString{DefaultDataDir()},
	}
	// MainchainJSONRPCFlag Mainchain JSON RPC endpoint
	MainchainJSONRPCFlag = cli.StringFlag{
		Name:  "mainchain.rpc",
		Usage: "Mainchain JSON RPC endpoint",
		Value: DefaultMainchainTestnetRPC,
	}
	// UndTestnetFlag configure for und test network
	UndTestnetFlag = cli.BoolFlag{
		Name:  "und-testnet",
		Usage: "configure for und test network",
	}

	// Registration flags

	// GenesisPathFlag Full path to the WRKChain's genesis.json. E.g.: /path/to/genesis.json
	GenesisPathFlag = cli.StringFlag{
		Name:  "genesis",
		Usage: "Full path to the WRKChain's genesis.json. E.g.: /path/to/genesis.json",
	}
	// AuthorisedAccountsFlag Comma separated list of addresses authorised to write to the WRKChain Root smart contract
	AuthorisedAccountsFlag = cli.StringFlag{
		Name:  "auth",
		Usage: "Comma separated list of addresses authorised to write to the WRKChain Root smart contract. No spaces. E.g.: 0x160B51e66e51327ac31C643f7675B8A9006aEE1E,0xbEc4127468c51fF89719DBcA5DC57F39C0049f06",
	}

	// Account flags

	// PasswordPathFlag Full path to the account password file
	PasswordPathFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Full path to the account password file. E.g. /path/to/.password",
	}
	// PrivateKeyPathFlag Full path to the private key file
	PrivateKeyPathFlag = cli.StringFlag{
		Name:  "key",
		Usage: "Full path to the private key file. E.g. /path/to/.private_key",
	}
	// AccountUnlockFlag Account to unlock
	AccountUnlockFlag = cli.StringFlag{
		Name:  "account",
		Usage: "Account to unlock - will be used tp write to the WRKChain Root smart contract when register and record commands are run. E.g. 0x160B51e66e51327ac31C643f7675B8A9006aEE1E",
	}

	// WRKChain flags

	// WRKChainJSONRPCFlag URI for the WRKChain's JSON RPC API
	WRKChainJSONRPCFlag = cli.StringFlag{
		Name:  "wrkchain.rpc",
		Usage: "URI for the WRKChain's JSON RPC API, e.g. http://localhost:8101",
	}
	// WriteFrequencyFlag Frequency WRKChain block hashes are written, in seconds
	WriteFrequencyFlag = cli.IntFlag{
		Name:  "freq",
		Usage: "Frequency WRKChain block hashes are written, in seconds. Default 3600",
		Value: 3600,
	}
	// RecordParentHashFlag If set, WRKChain Oracle will submit the WRKChain's parent hash
	RecordParentHashFlag = cli.BoolFlag{
		Name:  "hash.parent",
		Usage: "If set, WRKChain Oracle will submit the WRKChain's parent hash",
	}
	// RecordReceiptRootFlag If set, WRKChain Oracle will submit the WRKChain's Receipt Root hash
	RecordReceiptRootFlag = cli.BoolFlag{
		Name:  "hash.receipt",
		Usage: "If set, WRKChain Oracle will submit the WRKChain's Receipt Root hash",
	}
	// RecordTxRootFlag If set, WRKChain Oracle will submit the WRKChain's Tx Root hash
	RecordTxRootFlag = cli.BoolFlag{
		Name:  "hash.tx",
		Usage: "If set, WRKChain Oracle will submit the WRKChain's Tx Root hash",
	}
	// RecordStateRootFlag If set, WRKChain Oracle will submit the WRKChain's State Root hash
	RecordStateRootFlag = cli.BoolFlag{
		Name:  "hash.state",
		Usage: "If set, WRKChain Oracle will submit the WRKChain's State Root hash",
	}
)

// DirectoryString Custom type which is registered in the flags library which cli uses for
// argument parsing. This allows us to expand Value to an absolute path when
// the argument is parsed
type DirectoryString struct {
	Value string
}

// String to string
func (directoryString *DirectoryString) String() string {
	return directoryString.Value
}

// Set Set the DirectoryString value
func (directoryString *DirectoryString) Set(value string) error {
	directoryString.Value = expandPath(value)
	return nil
}

// DirectoryFlag Custom cli.Flag type which expand the received string to an absolute path.
// e.g. ~/.ethereum -> /home/username/.ethereum
type DirectoryFlag struct {
	Name  string
	Value DirectoryString
	Usage string
}

// String to string
func (directoryFlag DirectoryFlag) String() string {
	fmtString := "%s %v\t%v"
	if len(directoryFlag.Value.Value) > 0 {
		fmtString = "%s \"%v\"\t%v"
	}
	return fmt.Sprintf(fmtString, prefixedNames(directoryFlag.Name), directoryFlag.Value.Value, directoryFlag.Usage)
}

// GetName return Name attribute
func (directoryFlag DirectoryFlag) GetName() string {
	return directoryFlag.Name
}

// Set Set the DirectoryFlag value
func (directoryFlag *DirectoryFlag) Set(value string) {
	directoryFlag.Value.Value = value
}

func eachName(longName string, fn func(string)) {
	parts := strings.Split(longName, ",")
	for _, name := range parts {
		name = strings.Trim(name, " ")
		fn(name)
	}
}

// Apply called by cli library, grabs variable from environment (if in env)
// and adds variable to flag set for parsing.
func (directoryFlag DirectoryFlag) Apply(set *flag.FlagSet) {
	eachName(directoryFlag.Name, func(name string) {
		set.Var(&directoryFlag.Value, directoryFlag.Name, directoryFlag.Usage)
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
