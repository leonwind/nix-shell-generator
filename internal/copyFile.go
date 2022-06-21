package internal

import (
	"fmt"
	"io"
	"os"
)

type FileExistsError struct {
	message string
}

func (f FileExistsError) Error() string {
	return f.message
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CopyFile(source string, destination string, overwrite bool) error {
	in, err := os.Open(source)
	if err != nil {
		return err
	}

	if FileExists(destination) && !overwrite {
		return FileExistsError{fmt.Sprintf("File '%s' already exists, use '-f' if you want to overwrite it.", destination)}
	}

	out, err := os.Create(destination)
	if err != nil {
		fmt.Println("Open out")
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	if err = in.Close(); err != nil {
		return err
	}

	if err = out.Close(); err != nil {
		return err
	}

	return nil
}
