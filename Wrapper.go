package gssh

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
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

func (s SSH) Run() error {
	pipe, _ := s.cmd.StdoutPipe()
	err := s.cmd.Start()

	fmt.Println("argument list:")
	fmt.Println(strings.Join(s.cmd.Args, " "))

	if err != nil {
		fmt.Printf("error")
	} else {
		readerInline := bufio.NewReader(pipe)
		inline, err := readerInline.ReadString('\r')
		for err == nil {
			fmt.Printf("\r%s", inline)
			inline, err = readerInline.ReadString('\r')
		}
	}

	return err
}
