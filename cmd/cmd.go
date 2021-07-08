package cmd

import (
	"go.zoe.im/x/cli"
	"go.zoe.im/x/version"
)

var (
	// global cmd contains all sub command
	cmd = cli.New(
		cli.Name("gou"),
		cli.Short("A user service as a service implemented with Go."),
		version.NewOption(true),
		cli.Run(func(c *cli.Command, args ...string) {
			c.Help()
		}),
	)
)

// Register create a sub command
func Register(cmds ...*cli.Command) error {
	return cmd.Register(cmds...)
}

// Run execute command
func Run(opts ...cli.Option) error {
	return cmd.Run(opts...)
}
