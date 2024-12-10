package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kolukattai/kurl/models"
	"gopkg.in/yaml.v3"
)

func GetFileData(fileName string, config *models.Config, withDocumentation bool) (frontMatter models.FrontMatter, documentationString string, err error) {

	fileLocation := filepath.Join(config.FilePath, fileName)

	fmt.Println("file location", fileLocation)

	file, err := os.Open(fileLocation)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	frontMater := ""
	documentationString = ""
	index := 0

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
