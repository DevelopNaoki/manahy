package process

type YamlFile struct {
	Vms      map[string]Vm      `yaml:"vms"`
	Disks    map[string]Disk    `yaml:"disks"`
	Networks map[string]Network `yaml:"networks"`
}

type Vm struct {
	Name       string   `yaml:"name"`
	Generation int      `yaml:"generation"`
	CPU        CPU      `yaml:"cpu"`
	Memory     Memory   `yaml:"memory"`
	Path       string   `yaml:"path"`
	Image      string   `yaml:"image"`
	Disk       []string `yaml:"disk"`
	Network    []string `yaml:"network"`
}

type CPU struct {
	Thread int  `yaml:"thread"`
	Nested bool `yaml:"nested"`
}

type Memory struct {
	Size    int  `yaml:"size"`
	Dynamic bool `yaml:"dynamic"`
}

type Disk struct {
	Path       string `yaml:"path"`
	Size       string    `yaml:"size,omitempty"`
	Type       string `yaml:"type"`
	ParentPath string `yaml:"parent-path,omitempty"`
	SourceDisk int `yaml:"source-disk,omitempty"`
	Import     bool   `yaml:"import,omitempty"`
}

type Network struct {
	Name               string `yaml:"name"`
	Type               string `yaml:"type"`
	ExternameInterface string `yaml:"extername-interface,omitempty"`
	AllowManagementOs  bool   `yaml:"allow-management-os,omitempty"`
}
