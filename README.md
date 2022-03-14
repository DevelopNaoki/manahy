# manahy
manahy is management tool on hyper-v
- Works only on Windows

See implementedFeature.md for implemented features.

## Setup
```
> go mod download
> go build
```
if you usse not ASCII code, execute setup.bat
```
>.\setup.bat
```

## Usage
```
> manahy.exe -h
manahy is management tool on Hyper-V

Usage:
  manahy [flags]
  manahy [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  disk        disk is management virtual disk
  help        Help about any command
  switch      management switch on Hyper-V
  vm          vm is management vm on Hyper-V

Flags:
  -h, --help   help for manahy

Use "manahy [command] --help" for more information about a command.
```
