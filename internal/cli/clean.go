package cli

import (
	"cc-gi/internal/util"
	"os"
)

func Clean(config util.Config) {
	util.DebugLog("getting template directory")
	templateDir := config.TemplateDir
	util.DebugLog("deleting template directory")
	os.RemoveAll(templateDir)
}
