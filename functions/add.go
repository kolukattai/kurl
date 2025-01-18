package functions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
)

func AddNewCall(name string, folderName ...string) {
	fileName := ""
	if len(name) > 0 {
		fileName = filepath.Join(boot.Config.Path, name)
	}

	for _, v := range folderName {
		fileName = filepath.Join(v, fileName)
	}

	fmt.Printf("CONTENT: [%v] [%v] [%v]", boot.Config.Path, name, fileName)

	fileName = fmt.Sprintf("%v.md", fileName)

	fileName = strings.ReplaceAll(fileName, ".md.md", ".md")

	newUUID := uuid.New()

	headers := "headers: {\n\t\"Content-Type\": \"application/json\"\n}\n"

	fm := fmt.Sprintf("\nmethod: \"GET\"\nurl: \"{{BASE}}/api\"\n%v", headers)

	refComment := "# do not change refID, this key is used to connect this api with it's saved response"

	tmp := fmt.Sprintf("---\n%v\nrefID: %v\n%v\n---\n\n# %v\napi document goes here", refComment, newUUID, fm, strings.ReplaceAll(name, "-", " "))

	nameParts := strings.Split(fileName, "/")

	pathName := strings.Join(nameParts[:len(nameParts)-1], "/")

	err := os.MkdirAll(pathName, 0744)
	if err != nil {
		fmt.Println("nameParts", nameParts)
		panic(err)
	}

	if util.FileExists(fileName) {
		log.Fatal("File ", name, "already exists in the location", boot.Config.Path)
	}

	err = os.WriteFile(fileName, []byte(tmp), 0744)

	if err != nil {
		panic(err)
	}
}
