package pkg

import "os/exec"

func RunCmd(cmd string, args ...string) error {
	return exec.Command(cmd, args...).Start()
}
