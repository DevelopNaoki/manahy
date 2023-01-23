// Hyper-v operation and management package
package hyperv

type Vm struct {
	VmId      string
	VmName    string
	State     string
	Processor string
	Memory    string
}

type Vmswitch struct {
	VmswitchId   string
	VmswitchName string
	VmswitchType string
}
