package modules

import (
  
)

func BuildByStruct (summarize Summarize){
  for _, disk := range summarize.Disks {
    CreateDisk(disk)
  }
  
  for _, network := range summarize.Networks {
    CreateSwitch(network)
  }
  
  for _, vm := range summarize.Vms {
    CreateVm(vm)
  }
}

