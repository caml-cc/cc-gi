package cli

import (
	"cc-gi/internal/util"
	"io"
	"net/http"
	"os"
)

func Generate(config util.Config, f *Flags, args []string) {
	if len(args) == 0 {
		util.WarnLog("no sub-command arguments given")
		return
	}

	if f.Verbose {
		util.InfoLog("Generating .gitignore... (offline=%t)", f.Offline)
	}

	templateDir := config.TemplateDir

	templates := make([][]byte, 0)

	for _, arg := range args {
		templatePath := templateDir + "/" + arg + ".gitignore"

		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			if f.Offline {
				if f.Verbose {
					util.DebugLog("skipping fetch for %s", arg)
				}
				continue
			}

			if f.Verbose {
				util.InfoLog("fetching %s from API", arg)
			}

			URL := "https://gi.caml.cc/" + arg
			resp, err := http.Get(URL)
			if err != nil || resp.StatusCode != http.StatusOK {
				if f.Verbose {
					if f.Offline {
						util.DebugLog("failed to fetch template for %s", arg)
					}
					util.WarnLog("failed to fetch template for %s", arg)
				}
				continue
			}

			defer resp.Body.Close()

			content, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}

			os.MkdirAll(templateDir, 0755)
			os.WriteFile(templatePath, content, 0644)

			templates = append(templates, content)
		} else {
			util.InfoLog("reading %s", arg+".gitignore")

			content, err := os.ReadFile(templatePath)
			if err != nil {
				if f.Verbose {
					util.WarnLog("failed to read local template for", arg)
				}
				continue
			}
			util.InfoLog("appending %s to .gitignore", arg+".gitignore")
			templates = append(templates, content)
		}
	}

	var combinedContent []byte
	for _, template := range templates {
		if len(template) == 0 || template[len(template)-1] != '\n' {
			template = append(template, '\n')
		}
		combinedContent = append(combinedContent, template...)
	}

	outputFile := ".gitignore"
	if err := os.WriteFile(outputFile, combinedContent, 0644); err != nil {
		if f.Verbose {
			util.ErrorLog("Failed to write .gitignore file: %v", err)
		}
		return
	}

	if f.Verbose {
		util.InfoLog(".gitignore file generated successfully")
	}
}
