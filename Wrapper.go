package gssh

import (
	"bufio"
	"fmt"
	"os/exec"
)

func NewSSH(options SSHOptions, destination string, command string, args []string) *SSH {
	allArgs := append(mapArguments(options), destination, command)
	allArgs = append(allArgs, args...)

	return &SSH{
		SSHOptions: options,
		Destination: destination,
		Command: command,
		Args: args,
		cmd: exec.Command("ssh", allArgs...),
	}
}

func (s SSH) Run() {
	pipe, _ := s.cmd.StdoutPipe()
	err := s.cmd.Start()

	if err != nil {
		fmt.Printf("error")
	} else {
		readerInline := bufio.NewReader(pipe)
		inline, _ := readerInline.ReadString('\r')
		for inline != "" {
			fmt.Printf("\r%s", inline)
			inline, err = readerInline.ReadString('\r')
		}
	}
}

func (s SSH) GetFullArguments() []string {
	allArgs := append(mapArguments(s.SSHOptions), s.Destination, s.Command)
	allArgs = append(allArgs, s.Args...)
	return allArgs
}
