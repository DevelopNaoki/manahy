package process

type Summarize struct {
	Vms      map[string]Vm      `yaml:"vms" json:"vm"`
	Disks    map[string]Disk    `yaml:"disks" json:"disk"`
	Networks map[string]Network `yaml:"networks" json:"network"`
}

type Host struct {
	ConputerName string `yaml:"conputer-name"`
}

type Vm struct {
	Name       string   `yaml:"name" json:"name"`
	Count      int      `yaml:"count",omitempty`
	Generation int      `yaml:"generation" json:"generation"`
	Cpu        Cpu      `yaml:"cpu" json:"cpu"`
	Memory     Memory   `yaml:"memory" json:"network"`
	Path       string   `yaml:"path" json:"path"`
	Image      string   `yaml:"image,omitempty" json:"image"`
	Disk       []string `yaml:"disk"`
	Network    []string `yaml:"network"`
}

type Cpu struct {
	Thread         int  `yaml:"thread"`
	Reserve        int  `yaml:"reserve"`
	Maximum        int  `yaml:"maximum"`
	RelativeWeight int  `yaml:"relative-weight"`
	Nested         bool `yaml:"nested"`
}

type Memory struct {
	Size    string `yaml:"size"`
	Dynamic bool   `yaml:"dynamic"`
}

type Disk struct {
	Path       string `yaml:"path"`
	Size       string `yaml:"size,omitempty"`
	Type       string `yaml:"type,omitempty"`
	ParentPath string `yaml:"parent-path,omitempty"`
	SourceDisk int    `yaml:"source-disk,omitempty"`
	Import     bool   `yaml:"import,omitempty"`
}

type Network struct {
	Name               string `yaml:"name"`
	Type               string `yaml:"type"`
	ExternameInterface string `yaml:"extername-interface,omitempty"`
	AllowManagementOs  bool   `yaml:"allow-management-os,omitempty"`
}

type DiskList struct {
	Number       []string
	FriendlyName []string
	Size         []float64
	SizeUnit     []string
}

type SwitchList struct {
	External []string
	Internal []string
	Private  []string
}

type VmList struct {
	Running []string
	Saved   []string
	Paused  []string
	Off     []string
}
