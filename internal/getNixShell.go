package internal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func GetNixShellFile(storagePath string, destination string, overwrite bool) error {
	source, err := fuzzySearch(storagePath)
	if err != nil {
		return err
	}

	return CopyFile(source, destination, overwrite)
}

func fuzzySearch(path string) (string, error) {
	names, err := ListNixShellFiles(path)
	if err != nil {
		return "", err
	}

	idx, err := fuzzyfinder.Find(
		names,
		func(i int) string {
			return names[i]
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			content, err := ioutil.ReadFile(filepath.Join(path, names[i]))
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf(string(content))
		}))

	if err != nil {
		return "", err
	}

	return filepath.Join(path, names[idx]), nil
}
