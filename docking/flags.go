package main


import (
	"sort"
	"github.com/docking/cli"
	flag "github.com/docker/docker/pkg/mflag"
)

var (
	flHelp    = flag.Bool([]string{"h", "-help"}, false, "Print usage")
	flVersion = flag.Bool([]string{"v", "-version"}, false, "Print version information and quit")
)

type byName []cli.Command

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

var dockingCommands []cli.Command


func init() {
	for _, cmd := range cli.DockingCommands {
		dockingCommands = append(dockingCommands, cmd)
	}
	sort.Sort(byName(dockingCommands))
}

