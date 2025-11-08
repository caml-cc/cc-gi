package cli

import (
	"cc-gi/internal/util"
	"os"

	"github.com/spf13/pflag"
)

func Execute(f *Flags, config util.Config) {
	args := pflag.Args()

	if len(args) == 0 {
		util.WarnLog("No arguments provided")
		pflag.Usage()
		os.Exit(0)
	}

	util.InfoLog("Command: %s, Arguments: %v", args[0], args[1:])

	if f.Verbose {
		util.InfoLog("Verbose mode enabled")
	}

	switch args[0] {
	case "generate":
		Generate(config, f, args[1:])
	case "clean":
		if f.Verbose {
			util.InfoLog("Cleaning local templates...")
		}
		Clean(config)
	default:
		util.ErrorLog("Unknown command: %s", args[0])
		os.Exit(1)
	}
}
