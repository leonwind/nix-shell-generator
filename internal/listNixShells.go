package internal

import (
	"os"
	"sort"
)

func ListNixShellFiles(path string) ([]string, error) {
	root, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	names, err := root.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	sort.Strings(names)
	return names, err
}
