package build

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
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

	err := os.Mkdir(boot.Config.BuildDir, 0755)
	if err != nil {
		// panic(err)
		fmt.Println(err.Error())
	}

	processStaticFolder()

	data := struct {
		Title   string
		Message string
	}{
		Title:   "API Documentation",
		Message: "",
	}

	res := util.RenderTemplate(boot.TemplateFolder, nil, "home.html", data)

	os.WriteFile(filepath.Join(boot.Config.BuildDir, "index.html"), []byte(res), 0744)

	filesData();

	saveFile()
}

// // Helper function to recursively print file info with full path
// func printFileInfos(fileInfos []FileInfo, indent string) {
// 	for _, info := range fileInfos {
// 		// Print full path and type (folder or file)
// 		fmt.Printf("%s%s (Full Path: %s, IsFolder: %v)\n", indent, info.Name, info.FullPath, info.IsFolder)
// 		if info.IsFolder {
// 			printFileInfos(info.Children, indent+"  ")
// 		}
// 	}
// }
