package wrappers

import "os/exec"

// ShellWrapper is an abstructtion of shell command execution system library, no test cases required for this
type ShellWrapper interface {
	Exec(name string, args ...string) ([]byte, error)
}

type shellWrapper struct {
}

func (*shellWrapper) Exec(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.Output()
}

func NewShellWrapper() ShellWrapper {
	return &shellWrapper{}
}
