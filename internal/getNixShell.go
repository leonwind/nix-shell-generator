package internal

import "fmt"

func GetNixShellFile(source string, destination string, overwrite bool) error {
	if !FileExists(source) {
		return FileExistsError{message: fmt.Sprintf("Can't find file %s.", source)}
	}

	return CopyFile(source, destination, overwrite)
}
