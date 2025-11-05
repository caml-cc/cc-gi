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
		fmt.Println("Usage: cc-gi [options] <command>")
		fmt.Println("Options:")
		fmt.Println("  -o, --offline    Use only offline templates")
		fmt.Println("  -v, --verbose    Enable verbose output")
		fmt.Println("Commands:")
		fmt.Println("  generate")
		return
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
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		os.Exit(1)
	}
}
