package cmdcli

import (
	"flag"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type Cli struct {
	callName string
	cmds     map[string]func(args []string)
	flags    map[string]*string
	logger   log.Logger
}

var TestString *string

func NewCLI(callName string, logger log.Logger) (c *Cli) {
	return &Cli{
		callName: callName,
		cmds:     make(map[string]func(args []string)),
		flags:    make(map[string]*string),
		logger:   logger,
	}
}

// RegisterCmd - registers a new command options for the command CLI
// By default, commands refer to the first argment called
func (c *Cli) RegisterCmd(name string, f func(args []string)) (err error) {
	if _, ok := c.cmds[name]; ok {
		return errors.New(fmt.Sprintf("command %s already registered", name))
	}

	c.cmds[name] = f

	return nil
}

func (c *Cli) run() (err error) {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	cmdName := args[0]
	if cmd, ok := c.cmds[cmdName]; !ok {
		return errors.New(fmt.Sprintf("command `%s` unknown", cmdName))
	} else {
		cmd(args[1:])
	}

	return nil
}

// Run - executes the command CLI
// first args references the corresponding registered command
// untagged arguments are passed directly to receiving command
// TODO: add flag support
func (c *Cli) Run() {
	err := c.run()
	if err != nil {
		c.logger.Fatal(err)
	}
}
