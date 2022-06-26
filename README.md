# Nix-Shell-Generator

## Usages

### Store an existing `shell.nix`
Go into your project directory with an existing `shell.nix` file and run
```shell
nix-shell-generator --add nix-shell-descriptive-name
```

E.g. for a `shell.nix` file for a `C++` project using `GCC, cmake, boost, valgrind, gdb` you can run
```shell
nix-shell-generator --add gcc-boost-debugger
```

You can also add an `--path path/to/shell.nix` argument if your `shell.nix` is not in your current working directory.

### Add existing `shell.nix` into your project
Run `nix-shell-generator --get` to open a fuzzy finder to search and select your `shell.nix` template.
The template gets renamed and copied to `./shell.nix`.


## Installation 
Install by running

```shell
go install github.com/leonwind/nix-shell-generator/cmd@latest
```

Your `shell.nix` templates will be stored in `~/.config/nix-shell-generator/`.

For more practical usage, set up a shell alias: `alias nsg = "nix-shell-generator"`