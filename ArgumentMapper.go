package gssh

func mapArguments(options SSHOptions) []string {
	arguments := []string{}


	if !options.StrictHostKeyChecking {
		arguments = append(arguments, "-o", "StrictHostKeyChecking=no")
	}

	if options.UserKnownHostsFile != "" {
		arg := "UserKnownHostsFile=" + options.UserKnownHostsFile
		arguments = append(arguments, "-o", arg)
	}

	if options.LogLevel != "" {
		arg := "LogLevel=" + options.LogLevel
		arguments = append(arguments, "-o", arg)
	}

	if options.ForwardAgent {
		arguments = append(arguments, "-A")
	}

	if len(options.RemotePortForwardList) > 0 {
		for _, portForward := range options.RemotePortForwardList {
			arguments = append(arguments, "-R", portForward)
		}
	}

	if len(options.LocalPortForwardList) > 0 {
		for _, portForward := range options.LocalPortForwardList {
			arguments = append(arguments, "-L", portForward)
		}
	}

	if options.User != "" {
		arguments = append(arguments, "-l", options.User)
	}

	if options.Port != "" {
		arguments = append(arguments, "-p", options.Port)
	}

	if options.Verbose {
		arguments = append(arguments, "-v")
	}

	return arguments
}
