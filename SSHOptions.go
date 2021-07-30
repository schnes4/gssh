package gssh

type SSHOptions struct {
	StrictHostKeyChecking bool
	UserKnownHostsFile string
	LogLevel string
	ForwardAgent bool
	RemotePortForwardList []string
	LocalPortForwardList []string
	User string
	Port string
	Verbose bool
}
