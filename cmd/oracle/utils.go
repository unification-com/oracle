package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

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
	fmt.Println()
	os.Exit(1)
}

func MkDataDir(dirPath string) {
	if err := os.MkdirAll(dirPath, 0700); err != nil {
		Fatalf("Could not create datadir", "datadir", dirPath, "err", err)
	}
}
