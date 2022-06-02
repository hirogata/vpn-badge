package main

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

var version string
var commit string
var date string

type VpnBadge struct {
	VpnName string
}

func NewVpnBadge(name string) *VpnBadge {
	badge := &VpnBadge{
		VpnName: name,
	}

	return badge
}

func (badge *VpnBadge) ScanNetwork() (bool, error) {
	cmd := exec.Command("CMD.EXE", "/C", "chcp 437 && netsh interface show interface")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	output := strings.Split(string(out), "\n")

	found := false
	for _, s := range output {
		if strings.Contains(s, badge.VpnName) && strings.Contains(s, "Connected") {
			found = true
			break
		}
	}

	return found, err
}
