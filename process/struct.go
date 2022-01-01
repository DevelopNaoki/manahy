package process

type YamlFile struct {
	VMs []VM
}

type VM struct {
	Name       string    `yaml:name`
	Generation int       `yaml:generation`
	Memorys    Memory    `yaml:memory`
	Cpus       Cpu       `yaml:cpu`
	Path       string    `yaml:path`
	Image      string    `yaml:image`
	Disks      []Disk    `yaml:disk`
	Networks   []Network `yaml:network`
}

type Memory struct {
	Size    int  `yaml:size`
	Dynamic bool `yaml:dynamic`
}

type Cpu struct {
	Thread int  `yaml:thread`
	Nested bool `yaml:nested`
}

type Disk struct {
	Name   string `yaml:name`
	Size   int    `yaml:size`
	Path   string `yaml:path`
	Import bool   `yaml:import`
}

type Network struct {
	Name              string `yaml:name`
	Type              string `yaml:type`
	ExternalInterface string `yaml:external-interface`
	AllowManagementOS bool   `yaml:allow-management-os`
}
