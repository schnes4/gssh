package gssh

import (
	"os/exec"
)

type SSH struct {
	SSHOptions  SSHOptions
	Destination string
	Command     string
	Args        []string

	cmd         *exec.Cmd
}
