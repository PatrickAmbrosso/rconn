package utils

import (
	"fmt"
	"os/exec"
	"syscall"
)

type RDPConnection struct {
	HostAddress string
	Username    string
	Password    string
}

func ConnectRDP(params RDPConnection) error {
	target := "TERMSRV/" + params.HostAddress

	cmd := exec.Command("cmd.exe", "/c", fmt.Sprintf(
		"cmdkey /generic:%s /user:%s /pass:%s && start mstsc /v:%s",
		target, params.Username, params.Password, params.HostAddress,
	))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = nil // suppress output
	cmd.Stderr = nil // suppress error output too

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to launch RDP session with %w", err)
	}

	return nil
}
