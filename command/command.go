package command

import (
	"flag"
)

type Command struct {
	Flags   *flag.FlagSet
	Execute func(cmd *Command, args []string)
}

func (c *Command) Init(args []string) error {
	return c.Flags.Parse(args)
}

func (c *Command) Called() bool {
	return c.Flags.Parsed()
}

func (c *Command) Run() {

	c.Execute(c, c.Flags.Args())
}
