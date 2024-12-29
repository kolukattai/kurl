package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kolukattai/kurl/models"
	"gopkg.in/yaml.v2"
)

type FileFolderInfo struct {
	FileName string           `json:"fileName"`
	IsFolder bool             `json:"isFolder"`
	FilePath string           `json:"filePath"`
	Files    []FileFolderInfo `json:"files"`
}

func (st *FileFolderInfo) GetData(conf *models.Config) (frontMatter models.FrontMatter, documentationString string, err error) {
	return GetFileData(st.FilePath, conf, true, false)
}

func FileList(basePath string) ([]FileFolderInfo, error) {
	response := []FileFolderInfo{}

	files, err := os.ReadDir(filepath.Join(basePath))
	if err != nil {
		return nil, err
	}

	for _, v := range files {
		item := FileFolderInfo{Files: []FileFolderInfo{}}

		item.FileName = v.Name()
		item.IsFolder = v.IsDir()
		item.FilePath = filepath.Join(basePath, v.Name())

		if v.IsDir() {
			children, err := FileList(filepath.Join(basePath, v.Name()))
			if err != nil {
				return nil, err
			}

			item.Files = children
		}

		response = append(response, item)
	}

	return response, nil
}

func FileExists(filename string) bool {
	// Use os.Stat to get file info
	_, err := os.Stat(filename)

	// Check if the error is because the file doesn't exist
	if err != nil {
		if os.IsNotExist(err) {
			return false // File does not exist
		}
		// Other types of errors can be handled here if necessary
		fmt.Println("Error checking file:", err)
		return false
	}

	// File exists
	return true
}

func GetFileData(fileName string, config *models.Config, withDocumentation bool, skipFrontMatter bool) (frontMatter models.FrontMatter, documentationString string, err error) {

	fileLocation := filepath.Join(config.Path, fileName)

	fileLocation = strings.Replace(fileLocation, config.Path + "/", "", 1)
	
	fileLocation = filepath.Join(config.Path, fileLocation)

	fileLocation = fmt.Sprintf("%v.md", fileLocation)

	fileLocation = strings.ReplaceAll(fileLocation, ".md.md", ".md")

	file, err := os.Open(fileLocation)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	frontMater := ""
	documentationString = ""
	index := 0

	if skipFrontMatter {
		index = 2
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()

		if str == "---" && index == 0 {
			index = 1
			continue
		}

		if index == 1 {
			if str == "---" {
				index = 2
				continue
			}
			frontMater += fmt.Sprintf("%v\n", str)
			continue
		}

		if withDocumentation {
			if index > 1 {
				documentationString += fmt.Sprintf("%v\n", str)
				continue
			}
		} else {
			break
		}

	}

	err = yaml.Unmarshal([]byte(frontMater), &frontMatter)

	return
}

func GetFileName(fileName string) (name string, err error) {

	fileLocation := filepath.Join("", fileName)

	fileLocation = fmt.Sprintf("%v.md", fileLocation)

	fileLocation = strings.ReplaceAll(fileLocation, ".md.md", ".md")

	file, err := os.Open(fileLocation)
	if err != nil {
		return "", err
	}

	defer file.Close()

	nameArr := strings.Split(fileName, "/")
	name = strings.ReplaceAll(strings.ReplaceAll(nameArr[len(nameArr)-1], "-", " "), ".md", "")

	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()

		if str == "---" && index < 2 {
			index++
			continue
		}

		if strings.Contains(str, "# ") && index > 1 {
			name = strings.Replace(str, "# ", "", 1)
			break
		}

	}

	return
}
