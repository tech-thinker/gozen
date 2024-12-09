package repository

import "os/exec"

// ShellRepo is an abstructtion of shell command execution system library, no test cases required for this
type ShellRepo interface {
	Exec(name string, args ...string) ([]byte, error)
}

type shellRepo struct {
}

func (*shellRepo) Exec(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.Output()
}

func NewShellRepo() ShellRepo {
	return &shellRepo{}
}
