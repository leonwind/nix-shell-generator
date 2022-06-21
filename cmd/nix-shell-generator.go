package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/leonwind/nix-shell-generator/internal"
	"os"
	"path/filepath"
	"strings"
)

type stringFlag struct {
	set   bool
	value string
}

func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *stringFlag) String() string {
	return sf.value
}

// Storage path should be /home/user/.config/nix-shell-generator under nix systems.
func getStoragePath() string {
	config, err := os.UserConfigDir()
	if err != nil {
		throwAndExit(err)
	}
	return filepath.Join(config, "nix-shell-generator")
}

// Create ~/.config/nix-shell-generator if it doesn't exist yet
func initStorage() {
	if !internal.FileExists(getStoragePath()) {
		if err := os.MkdirAll(getStoragePath(), os.ModePerm); err != nil {
			throwAndExit(err)
		}
	}
}

func throwAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	initStorage()

	var addNixShellName stringFlag
	flag.Var(&addNixShellName, "add", "Add ./shell.nix to storage with a custom name.")

	var nixShellPath = flag.String("path", "shell.nix", "[Optional] Path of the nix.shell file to add to storage.")

	var getNixShellName stringFlag
	flag.Var(&getNixShellName, "get", "Copy the shell.nix from storage by name in the current working dir.")

	overwrite := flag.Bool("force", false, "Force to overwrite existing files.")
	listNixShells := flag.Bool("list", false, "List all available nix.shell files in storage.")

	flag.Parse()

	if addNixShellName.set {
		destination := filepath.Join(getStoragePath(), addNixShellName.value)
		if !strings.HasSuffix(destination, ".nix") {
			destination += ".nix"
		}

		if err := internal.AddNixShellFile(*nixShellPath, filepath.Join(getStoragePath(), addNixShellName.value), *overwrite); err != nil {
			throwAndExit(err)
		}
		return
	}

	if getNixShellName.set {
		wd, err := os.Getwd()
		if err != nil {
			throwAndExit(err)
		}

		source := filepath.Join(getStoragePath(), getNixShellName.value+".nix")
		destination := filepath.Join(wd, "shell.nix")
		err = internal.GetNixShellFile(source, destination, *overwrite)

		if err != nil {
			throwAndExit(err)
		}
		return
	}

	if *listNixShells {
		names, err := internal.ListNixShellFiles(getStoragePath())
		if err != nil {
			throwAndExit(err)
		}

		for _, name := range names {
			fmt.Println(name)
		}
		return
	}

	throwAndExit(errors.New("use --help for usage information"))
}
