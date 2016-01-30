package cli

import flag "github.com/docker/docker/pkg/mflag"

// ClientFlags represents flags for the docking client.
type ClientFlags struct {
	FlagSet   *flag.FlagSet
	Common    *CommonFlags
	PostParse func()

	ConfigDir string
}