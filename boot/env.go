package boot

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/kolukattai/kurl/models"
)

func UpdateConfig(configFileName string, context string) {
	defaultVal := func() {
		Config = &models.Config{
			FilePath:     "api",
			EnvVariables: []models.EnvVariables{},
			DefaultEnv:   0,
			BuildDir:     "dist",
			Title:        "Kurl Docs",
		}
	}

	filePath := filepath.Join(context, configFileName)

	file, err := os.ReadFile(filePath)
	if err != nil {
		defaultVal()
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		defaultVal()
	}
}
