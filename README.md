# manahy
manahy is a tool to manage Hyper-V. Currently it has few features, but we plan to add more in the future.

This tool only works on Windows.

## Purpose
There are many commands to manipulate Hyper-V, and they are not unified into one command.
manahy brings these commands together to make it easier to manipulate Hyper-V.

## Install
Download the executable file from the <a href="https://github.com/DevelopNaoki/manahy/releases/tag/v0.0.0-beta">following page</a> and save it in a PATHed location or pass PATH
Execute "manahy version" in powershell or cmd, and if the following results are displayed, installation is complete.
```
> manahy.exe version
manahy version 0.0.0 (beta)
```

## Usage
```
> manahy.exe -h
manahy is management tool on Hyper-V

Usage:
  manahy [flags]
  manahy [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  hyperv      management Hyper-V
  version     print manahy version
  vm          management vm on Hyper-V

Flags:
  -h, --help   help for manahy

Use "manahy [command] --help" for more information about a command.
```
### completion
Generates a script that can complete subcommands
```
> manahy.exe completion powershell -h
Generate the autocompletion script for powershell.

To load completions in your current shell session:

        manahy completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

Usage:
  manahy completion powershell [flags]

Flags:
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### hyperv
Command to configure Hyper-V settings.
```
> manahy.exe hyperv -h
management Hyper-V

Usage:
  manahy hyperv [flags]
  manahy hyperv [command]

Available Commands:
  check       check Hyper-V Enabled

Flags:
  -h, --help   help for hyperv

Use "manahy hyperv [command] --help" for more information about a command.
```

### vm
This command is used to operate and manage VMs.
```
> manahy.exe vm -h
management vm on Hyper-V

Usage:
  manahy vm [flags]
  manahy vm [command]

Available Commands:
  destroy     destroy VM
  list        Print VM list
  reboot      reboot VM
  resume      resume VM
  save        save VM
  shutdown    shutdown VM
  start       start VM
  suspend     suspend VM

Flags:
  -h, --help   help for vm

Use "manahy vm [command] --help" for more information about a command.
```
* There are two ways to operate vm power: by specifying the vm name or by specifying the vmid.
If vms with the same name exist, specifying the vm name will perform the operation on all of them. However, if vmid is used, the operation will be performed on only one vm with that vmid.

## Help
### If the command execution results in garbled characters
Execute the following command This command changes the terminal character encoding to UTF-8
```
> chcp 65001
```
or
```
> ./configCharCode.bat
```

### When exiting with an error
May not have enough authority. Try running it with administrator privileges.
