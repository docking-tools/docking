package cli

import (
	flag "github.com/docker/docker/pkg/mflag"	
)

// CommonFlags represents flags that are common to both the client and the daemon.
type CommonFlags struct {
	FlagSet   *flag.FlagSet
	PostParse func()

	
}

// Command is the struct contains command name and description
type Command struct {
	Name        string
	Description string
}

var dockingCommands = []Command{
	{"build", "Build an app"},
	{"rm", "Remove "},
	{"start", "Start one or more application environment"},
	{"stop", "Stop a running application environment"},
}

// DockingCommands stores all the docking command
var DockingCommands = make(map[string]Command)

func init() {
	for _, cmd := range dockingCommands {
		DockingCommands[cmd.Name] = cmd
	}
}