package utils

import (
	"fmt"
	"os/exec"
	"rconn/src/models"
	"syscall"
)

func ConnectRDP(params models.RDPConnectionParams) error {
	target := "TERMSRV/" + params.Host

	cmd := exec.Command("cmd.exe", "/c", fmt.Sprintf(
		"cmdkey /generic:%s /user:%s /pass:%s && start mstsc /v:%s",
		target, params.User, params.Password, params.Host,
	))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to launch RDP session with %w", err)
	}

	return nil
}
