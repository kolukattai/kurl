package build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kolukattai/kurl/boot"
)

func processStaticFolder() {
	skipPattern := []string{".css.map", ".scss", "template"}

	skipThis := func(ite string) bool {
		for _, v := range skipPattern {
			if strings.Contains(ite, v) {
				return true
			}
		}
		return false
	}

	err := os.MkdirAll(filepath.Join(boot.Config.Build, "static"), 0744)
	if err != nil {
		panic(err)
	}

	// Set the root folder of embedded files (can be empty for the root directory)
	basePath := "static"

	// Get the file information as an array of FileInfo
	fileInfos, err := getFileInfoFromFS(boot.StaticFolder, basePath)
	if err != nil {
		fmt.Println("Error reading embedded FS:", err)
		return
	}

	for _, v := range fileInfos {
		fmt.Println("STATIC", v.Name)
		folderPath := filepath.Join(boot.Config.Build, "static", strings.Replace(v.FullPath, v.Name, "", 1))

		if skipThis(folderPath) {
			fmt.Println("STATIC SKIPPED", folderPath)
			continue
		}

		err := os.MkdirAll(folderPath, 0744)
		if err != nil {
			panic(err)
		}

		filePath := filepath.Join(boot.Config.Build, "static", v.FullPath)

		if v.IsFolder {
			continue
		}

		if skipThis(v.Name) {
			continue
		}

		fmt.Println("PF", filePath)
		err = os.WriteFile(filePath, v.FileData, 0744)
		if err != nil {
			panic(err)
		}
	}
}
