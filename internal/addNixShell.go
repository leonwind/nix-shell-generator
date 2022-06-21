package internal

func AddNixShellFile(source string, destination string, overwrite bool) error {
	return CopyFile(source, destination, overwrite)
}
