package cli

import (
	"os"
)

func Clean(f *Flags, args []string) {
	homeDir, _ := os.UserHomeDir()
	templateDir := homeDir + "/.local/share/cc-gi/templates"
	os.RemoveAll(templateDir)
}
