# github.com/DevelopNaoki/manahy/process

## Overview
### function name list
| function name | file name |
| - | - |
| Error | error.go |
| ReadFile | file.go |
| GetDiskList | disk.go |
| GetDIskListInfo | disk.go |
| CreateDisk | disk.go |
| GetSwitchList | switch.go |
| GetSwitchType | switch.go |
| ChangeSwitchType | switch.go |
| CreateSwitch | switch.go |
| RemoveSwitch | switch.go |
| GetVmList | vm.go |
| GetVmState | vm.go |
| SetVmProcessor | vm.go |
| CreateVm | vm.go |
| CheckVmOption | vm.go |
| isFileExist | internal.go |
| isFolderExist | internal.go |
| computCapacity | interna.go |
| listingOfExecuteResults | internal.go |

---
### struct name list
| struct name | file name |
| - | - |
| YamlFile | struct.go |
| Vm | struct.go |
| Cpu | struct.go |
| Memory | struct.go |
| Disk | struct.go |
| Network | struct.go |
| DiskList | struct.go |

---
## More
### function list
- Error (code int)
    - github.com/DevelopNaoki/manahy/process内のerrorを処理する
    - 参照先
        - fmt
- ReadFile(name string) (data YamlFile)
    - manahy.yamlからYamlFileに型変換する
    - 参照先
        - ioutil.ReadFile
        - fmt
        - yaml.Unmarshal
        - os.Exit
- GetDiskList () (diskList DiskList)
    - Storageのリストを取得する
    - 参照先
        - GetDiskListInfo
        - computCapacity
        - strconv.Atoi
- GetDiskListInfo (parameter string) (list []string)
    - Storageのリストの一パラメータを取得する
    - 参照先
        - listingOfExecuteResults
        - exec.Command
- CreateDisk (newDisk Disk)
    - 仮想ディスクを作成する
    - 参照先
        - isFileExist
        - fmt
- GetSwitchList (state string) (list []string)
    - Switchのリストを取得する
    - 参照先
        - listingOfExecuteResults
        - exec.Command
- GetSwitchType (name string) (state string)
    - SwitchのTypeを取得する
    - 参照先
        - listingOfExecuteResults
        - exec.Command
- ChangeSwitchType (name string, switchType string)
    - SwitchのTypeを変更する
    - 参照先
        - exec.Command
- CreateSwitch (newSwitch Network)
    - Switchを作成する
    - 参照先
        - GetSwitchType
        - exec.Command
        - fmt
- RemoveSwitch (name string)
    - Switchを削除する
    - 参照先
        - GetSwitchType
        - exec.Command
        - fmt

### struct list
- YamlFile
    - manahy.yamlに適用する
    - Including
        - Vm
        - Disk
        - Network
- Vm
    - VMの構成情報
    - Including
        - Cpu
        - Memory
- Cpu
    - CPUの構成情報
- Memory
    - Memoryの構成情報
- Disk
    - Diskの構成情報
- Network
    - Networkの構成情報
- DiskList
    - 実行結果を一時的に保存する

