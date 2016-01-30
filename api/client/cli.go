package client

import (
    "io"
    "github.com/docking/cli"
)

// DockingCli represents the docking command line client.
// Instances of the client can be returned from NewDockingCli.
type DockingCli struct {
	// initializing closure
	init func() error

	// configFile has the client configuration file
//	configFile *cliconfig.ConfigFile
	// in holds the input stream and closer (io.ReadCloser) for the client.
	in io.ReadCloser
	// out holds the output stream (io.Writer) for the client.
	out io.Writer
	// err holds the error stream (io.Writer) for the client.
	err io.Writer
	// keyFile holds the key file as a string.
	keyFile string
	// inFd holds the file descriptor of the client's STDIN (if valid).
	inFd uintptr
	// outFd holds file descriptor of the client's STDOUT (if valid).
	outFd uintptr
	// isTerminalIn indicates whether the client's STDIN is a TTY
	isTerminalIn bool
	// isTerminalOut indicates whether the client's STDOUT is a TTY
	isTerminalOut bool
	// client is the http client that performs all API operations
//	client client.APIClient
	// state holds the terminal state
//	state *term.State
}

// NewDockingCli returns a DockingCli instance with IO output and error streams set by in, out and err.
// The key file, protocol (i.e. unix) and address are passed in as strings, along with the tls.Config. If the tls.Config
// is set the client scheme will be set to https.
func NewDockingCli(in io.ReadCloser, out, err io.Writer, clientFlags *cli.ClientFlags) *DockingCli {
    
    cli := &DockingCli{
		in:      in,
		out:     out,
		err:     err,
//		keyFile: clientFlags.Common.TrustKey,
	}
    
    return cli
}
