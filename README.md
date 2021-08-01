# gssh - golang ssh wrapper

[![Build Status](https://travis-ci.com/schnes4/gssh.svg)](https://travis-ci.com/schnes4/gssh)
[![codecov](https://codecov.io/gh/schnes4/gssh/branch/master/graph/badge.svg)](https://codecov.io/gh/schnes4/gssh)

This lib is a lightweight ssh wrapper for golang.
The lib provides a convenient wrapper for configure and execute ssh commands on remote systems.

## Configuration

All available configurations are available in the `SSHOptions` struct.
The struct covers all configurations which are necessary for my current use cases.
Feel free to contribute to the wrapper!

## Usage

To execute the ssh process call the `Run()` method.

If you only interested in the list of arguments and run the command yourself you could run the `GetFullArguments()` method.

## Example

```golang
package main

import (
	"fmt"
	"github.com/schnes4/gssh"
)

func main() {
	sshConfig := gssh.SSHOptions{
		StrictHostKeyChecking: false,
		UserKnownHostsFile:    "/dev/null",
		LogLevel:              "ERROR",
		ForwardAgent:          true,
		User:                  "root",
		Port:                  "22000",
		RemotePortForwardList: []string{"22000:10.0.0.1:22"},
	}

	ssh := gssh.NewSSH(sshConfig, "10.0.0.2", "ls", []string{"-lah", "/root"})

	argList := ssh.GetFullArguments()
	fmt.Println(argList)

	ssh.Run()
}
```
