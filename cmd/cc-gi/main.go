package main

import (
	"cc-gi/internal/cli"
	"cc-gi/internal/util"
)

func main() {
	f := cli.ParseFlags()
	util.SetVerbose(f.Verbose)

	config := util.LoadConfig()
	cli.Execute(f, *config)
}
