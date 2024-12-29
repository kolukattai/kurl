package boot

import (
	"os"
	"path/filepath"

	"github.com/kolukattai/kurl/models"
	"gopkg.in/yaml.v2"
)

func UpdateConfig(configFileName string, context string) {
	defaultVal := func() {
		Config = &models.Config{
			Path:         "./",
			EnvVariables: map[string]string{},
			Build:        "dist",
			Title:        "Kurl Docs",
		}
	}

	filePath := filepath.Join(context, configFileName)

	file, err := os.ReadFile(filePath)
	if err != nil {
		defaultVal()
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		defaultVal()
	}
}
