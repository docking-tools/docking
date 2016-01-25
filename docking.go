package main

import  (
    "flag"
    "fmt"
    "os"
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

		help += "\nRun 'docker COMMAND --help' for more information on a command."
		fmt.Fprintf(os.Stdout, "%s\n", help)
	}
  flag.Parse()
  
  if *flVersion {
		showVersion()
		return
	}
  
  fmt.Println("tail:", flag.Args())

  
}

// Construct version for --version Flag
func showVersion() {
		fmt.Printf("Docking version %s, build %s\n", dockingversion.Version, dockingversion.GitCommit)
}