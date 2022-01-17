package main

import (
	"flag"
	"fmt"
	"os"
)

const usageString string = `Go is a tool for managing Go source code.

Usage:

	microblog-client [COMMAND] [arguments]

The commands are:

`

type command struct {
	usage string
	f     func(*app)
}

func newCommand(usage string, f func(*app)) command {
	return command{
		usage: usage,
		f:     f,
	}
}

type app struct {
	commands map[string]command
	username string
}

func (app *app) run(subcommand string) {
	if subcommand == "" {
		subcommand = "help"
	}
	cmd, ok := app.commands[subcommand]
	if !ok {
		fmt.Printf("Unknown subcommand %s\n", subcommand)
		os.Exit(1)
	}

	cmd.f(app)
}

func help(app *app) {
	fmt.Printf("%s", usageString)
	for name, command := range app.commands {
		fmt.Printf("\t%s\t%s\n", name, command.usage)
	}
	fmt.Printf("\nArguments:\n")
	flag.PrintDefaults()
}
