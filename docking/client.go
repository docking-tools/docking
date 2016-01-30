package main

import (
    "github.com/docking/cli"
    flag "github.com/docker/docker/pkg/mflag"
)

// Define client set of Flag and Command.

var clientFlags = &cli.ClientFlags{FlagSet: new(flag.FlagSet), Common: commonFlags}

func init() {
//	client := clientFlags.FlagSet
//	client.StringVar(&clientFlags.ConfigDir, []string{"-config"}, cliconfig.ConfigDir(), "Location of client config files")

	clientFlags.PostParse = func() {
		clientFlags.Common.PostParse()

//		if clientFlags.ConfigDir != "" {
//			cliconfig.SetConfigDir(clientFlags.ConfigDir)
//		}

	}
}