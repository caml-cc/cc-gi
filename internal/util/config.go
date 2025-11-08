package util

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TemplateDir string
}

func LoadConfig() *Config {
	homeDir, _ := os.UserHomeDir()
	configDir := homeDir + "/.config/cc-gi"
	configFile := homeDir + "/.config/cc-gi/config.toml"
	templateDir := homeDir + "/.local/share/cc-gi/templates"

	_, err := os.Stat(configFile)
	if err != nil {
		os.MkdirAll(configDir, 0755)
		os.WriteFile(configFile, []byte("TemplateDir = '"+templateDir+"'\n"), 0644)

		return &Config{
			TemplateDir: templateDir,
		}
	}

	var config Config
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
