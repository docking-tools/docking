package client

import (
        "fmt"
)

// CmdBuild builds a new app from the source code at a given path.
//
// If '-' is provided instead of a path or URL, Docking will build an app from xxxx a xxx read from STDIN.
// Usage: docking build [OPTIONS] PATH | URL | -

func (cli *DockingCli) CmdBuild(args ...string) error {
   fmt.Print("Build  cmd\n")  
   
   	return nil
}