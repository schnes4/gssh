# gssh - golang ssh wrapper

[![Build Status](https://travis-ci.com/schnes4/gssh.svg)](https://travis-ci.com/schnes4/gssh)
[![codecov](https://codecov.io/gh/schnes4/gssh/branch/master/graph/badge.svg)](https://codecov.io/gh/schnes4/gssh)

This lib is a lightweight ssh wrapper for golang.
The lib provides a convenient wrapper for configure and execute ssh commands on remote systems.

## Usage

All available configurations are available in the `SSHOptions` struct.
The struct covers all configurations which are necessary for my current use cases.
Feel free to contribute to the wrapper!

To execute the ssh process call the `Run()` method.

## Example

```golang
package main

import (
	"github.com/schnes4/gssh"
)

func main () {
	sshConfig := gssh.SSHOptions{
		StrictHostKeyChecking: false,
		UserKnownHostsFile: "/dev/null",
		LogLevel: "ERROR",
		ForwardAgent: true,
		User: "root",
		Port: "22000",
		RemotePortForwardList: []string{"22000:10.0.0.1:22"},
	}

	ssh := gssh.NewSSH(sshConfig, "10.0.0.2", "ls", []string{"-lah", "/root"})
	ssh.Run()
}
```
