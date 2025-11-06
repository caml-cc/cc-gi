package cli

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Flags struct {
	Offline bool
	Verbose bool
}

func ParseFlags() *Flags {
	pflag.Usage = func() {
		fmt.Println("Usage: cc-gi [options] <command> [templates]")
		fmt.Println("Options:")
		fmt.Println("  -o, --offline    Use only offline templates")
		fmt.Println("  -v, --verbose    Enable verbose output")
		fmt.Println("  -h, --help       Display help")
		fmt.Println("Commands:")
		fmt.Println("  generate, clean")
		os.Exit(0)
	}

	offline := pflag.BoolP("offline", "o", false, "local files only")
	verbose := pflag.BoolP("verbose", "v", false, "verbose output")
	help := pflag.BoolP("help", "h", false, "display help")

	if *help {
		pflag.Usage()
	}

	pflag.Parse()

	return &Flags{
		Offline: *offline,
		Verbose: *verbose,
	}
}
