package main

import (
    "github.com/docking/cli"
    flag "github.com/docker/docker/pkg/mflag"
)

var (
	commonFlags = &cli.CommonFlags{FlagSet: new(flag.FlagSet)}

)


func init() {
    
    // Set here common flag
    // and add element to cli.CommonFlags struct
    
 //   cmd := commonFlags.FlagSet
    
    
}

func postParseCommon() {
    // Set here control ifSet, ifExist, ...
}