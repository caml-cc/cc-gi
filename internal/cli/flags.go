package cli

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	Offline bool
	Verbose bool
}

func ParseFlags() *Flags {
	offline := pflag.BoolP("offline", "o", false, "local files only")
	verbose := pflag.BoolP("verbose", "v", false, "verbose output")

	pflag.Parse()

	return &Flags{
		Offline: *offline,
		Verbose: *verbose,
	}
}
