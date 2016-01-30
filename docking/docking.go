package main

import  (
    "fmt"
    "os"
    "github.com/docking/dockingversion"
    "github.com/docking/api/client"
    "github.com/docking/cli"
    flag "github.com/docker/docker/pkg/mflag"
)


/*
* Interface for command line and args runner.
*/
type Command interface {
    
}

func main() {
  fmt.Print("Hello, world\n")
  
  // Define global flag here
  
  
  // end global flag definition
  
  
  flag.Usage = func() {
		fmt.Fprint(os.Stdout, "Usage: docking [OPTIONS] COMMAND [arg...]\n       docking [ --help | -v | --version ]\n\n")
		fmt.Fprint(os.Stdout, "A self-sufficient runtime for containers.\n\nOptions:\n")

		flag.CommandLine.SetOutput(os.Stdout)
		flag.PrintDefaults()

		help := "\nCommands:\n"

		for _, cmd := range dockingCommands {
			help += fmt.Sprintf("    %-10.10s%s\n", cmd.Name, cmd.Description)
		}

		help += "\nRun 'docking COMMAND --help' for more information on a command."
		fmt.Fprintf(os.Stdout, "%s\n", help)
	}
  flag.Parse()
  
  if *flVersion {
		showVersion()
		return
	}
  

  if *flHelp {
		// if global flag --help is present, regardless of what other options and commands there are,
		// just print the usage.
		flag.Usage()
		return
   }
   
  
	clientCli := client.NewDockingCli(os.Stdin, os.Stdout, os.Stderr, clientFlags)

	c := cli.New(clientCli)
	if err := c.Run(flag.Args()...); err != nil {
		if sterr, ok := err.(cli.StatusError); ok {
			if sterr.Status != "" {
				fmt.Fprintln(os.Stderr, sterr.Status)
				os.Exit(1)
			}
			os.Exit(sterr.StatusCode)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}   

  
}

// Construct version for --version Flag
func showVersion() {
		fmt.Printf("Docking version %s, build %s\n", dockingversion.Version, dockingversion.GitCommit)
}