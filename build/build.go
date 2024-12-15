package build

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kolukattai/kurl/boot"
)

// FileInfo struct to represent file details, including whether it's a folder
type FileInfo struct {
	Name     string     // File or folder name
	IsFolder bool       // Flag if it's a folder
	Children []FileInfo // If it's a folder, contain child files/folders
	FullPath string     // Full path to the file or folder
	FileData []byte     // If it's a file, contains the file content as a byte slice
}

// Function to retrieve the list of files from the embedded filesystem, including full paths and file content
func getFileInfoFromFS(fSys embed.FS, basePath string) ([]FileInfo, error) {
	var fileInfos []FileInfo

	// Walk through the directory structure using embedded FS
	err := fs.WalkDir(fSys, basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the base path itself
		if path == basePath {
			return nil
		}

		// Get the full relative path in the embedded FS
		relPath, _ := filepath.Rel(basePath, path)

		// Check if it's a directory
		isDir := d.IsDir()
		if isDir {
			// It's a directory, add to the fileInfos array, and recursively handle children
			children, err := getFileInfoFromFS(fSys, path)
			if err != nil {
				return err
			}
			fileInfos = append(fileInfos, FileInfo{
				FullPath: relPath,
				Name:     filepath.Base(path),
				IsFolder: true,
				Children: children,
			})
		} else {
			// It's a file, read the file content into FileData
			data, err := fSys.ReadFile(path)
			if err != nil {
				return err
			}
			fileInfos = append(fileInfos, FileInfo{
				FullPath: relPath,
				Name:     filepath.Base(path),
				IsFolder: false,
				Children: nil,
				FileData: data,
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}

func Run() {

	skipPattern := []string{".css.map", ".scss"}

	err := os.Mkdir(boot.Config.BuildDir, 0744)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(boot.Config.BuildDir, "static"), 0744)
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
		folderPath := filepath.Join(boot.Config.BuildDir, "static", strings.Replace(v.FullPath, v.Name, "", 1))

		err := os.MkdirAll(folderPath, 0744)
		if err != nil {
			panic(err)
		}

		filePath := filepath.Join(boot.Config.BuildDir, "static", v.FullPath)

		if v.IsFolder {
			continue
		}

		skipThis := func(ite string) bool {
			for _, v := range skipPattern {
				if strings.Contains(ite, v) {
					return true
				}
			}
			return false
		}

		if skipThis(v.Name) {
			continue
		}

		err = os.WriteFile(filePath, v.FileData, 0744)
		if err != nil {
			panic(err)
		}
	}

	// Print the file info
	printFileInfos(fileInfos, "")

}

// Helper function to recursively print file info with full path
func printFileInfos(fileInfos []FileInfo, indent string) {
	for _, info := range fileInfos {
		// Print full path and type (folder or file)
		fmt.Printf("%s%s (Full Path: %s, IsFolder: %v)\n", indent, info.Name, info.FullPath, info.IsFolder)
		if info.IsFolder {
			printFileInfos(info.Children, indent+"  ")
		}
	}
}
