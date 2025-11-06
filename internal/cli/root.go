package cli

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func Execute() {
	f := ParseFlags()
	args := pflag.Args()

	if len(args) == 0 {
		pflag.Usage()
		os.Exit(0)
	}

	if f.Verbose {
		fmt.Println("Verbose mode enabled")
	}

	switch args[0] {
	case "generate":
		if f.Verbose {
			fmt.Printf("Generating .gitignore... (offline=%t)\n", f.Offline)
		}
		Generate(f, args[1:])
	case "clean":
		if f.Verbose {
			fmt.Println("Cleaning local templates...")
		}
		Clean(f, args[1:])
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		os.Exit(1)
	}
}
