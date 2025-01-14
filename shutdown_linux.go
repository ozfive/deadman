package main

import (
	"fmt"
	"syscall"
)

func shutdownNow() error {

	// First, try to call the reboot system call with the appropriate arguments
	if err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF); err == nil {
		return nil
	}

	// If that fails, try to call the halt system call with the appropriate arguments
	if err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_HALT); err == nil {
		return nil
	}

	// If both system calls fail, return an error
	return fmt.Errorf("Failed to initiate system shutdown.")
}
