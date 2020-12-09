package cmd

import "os/exec"

// RunPsl ... Run command with powershell.
func RunPsl(command string) ([]byte, error) {
	return exec.Command("powershell", command).Output()
}

// RunLx ... Run command with linux shell.
func RunLx(command string) ([]byte, error) {
	return exec.Command(command).Output()
}
