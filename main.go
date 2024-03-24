package main

import (
	"flag"
	"fmt"
	"gotem/command"
	"gotem/version"
	"os"
)

func main() {
	const USAGE = `Usage: gotem command [options] A simple tool to generate and manage custom templates 
	Options: 
	- Commands: 	Adds a template to the collection from a local file edit Uses the default 
			text editor to modify a stored template list Lists all stored templates create 
			Generates an instance of a template in the current directory delete Removes a 
			stored template version Prints version info to console `

	var cmd *command.Command

	switch os.Args[1] {
	case "version":
		cmd = version.NewVersionCommand()
	default:
		usageAndExit(fmt.Sprintf("gotem: '%s' is not a gotem command.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()

}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}
