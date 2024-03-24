package version

import (
	"flag"
	"fmt"
	"gotem/command"
	"os"
)

var versionUsage = `Print the app version and build info for the current context. Usage: gotem version [options] Options: --short If true, print just the version number. Default false. `
var (
	build   = "v1.0"
	version = "Build 1.0.0"
	short   = false
)

func NewVersionCommand() *command.Command {
	cmd := &command.Command{
		Flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	cmd.Flags.BoolVar(&short, "short", false, "")
	cmd.Flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}

	return cmd
}

var versionFunc = func(cmd *command.Command, args []string) {
	if short {
		fmt.Printf("brief version: v%s", version)
	} else {
		fmt.Printf("brief version: v%s, build: %s", version, build)
	}
	os.Exit(0)
}
