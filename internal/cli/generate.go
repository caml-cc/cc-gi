package cli

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Generate(f *Flags, args []string) {
	homeDir, _ := os.UserHomeDir()
	templateDir := homeDir + "/.local/share/cc-gi/templates"

	templates := make([][]byte, 0)

	for _, arg := range args {
		templatePath := templateDir + "/" + arg + ".gitignore"

		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			if f.Offline {
				if f.Verbose {
					fmt.Println("skipping fetch for", arg)
				}
				continue
			}
			
			if f.Verbose {
				fmt.Println("fetching " + arg + " from API")
			}

			URL := "https://gi.caml.cc/" + arg
			resp, err := http.Get(URL)
			if err != nil || resp.StatusCode != http.StatusOK {
				if f.Verbose {
					fmt.Println("failed to fetch template for", arg)
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
			content, err := os.ReadFile(templatePath)
			if err != nil {
				if f.Verbose {
					fmt.Println("failed to read local template: ", arg)
				}
				continue
			}
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
			fmt.Println("Failed to write .gitignore file:", err)
		}
		return
	}

	if f.Verbose {
		fmt.Println(".gitignore file generated successfully")
	}
}
