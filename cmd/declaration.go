package cmd

import (
	"github.com/DevelopNaoki/manahy/modules"
)

var vmListOption struct {
	active   bool
	saved    bool
	inactive bool
	paused   bool
	all      bool
}
var newVmName string

var diskCreateOption modules.Disk

var switchListOption struct {
	external bool
	internal bool
	private  bool
	all      bool
}
var switchCreateOption modules.Network
var newSwitchName string
var switchType string
var netAdapter string
