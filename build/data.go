package build

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/handler"
	"github.com/kolukattai/kurl/util"
)

func createFileData(li []util.FileFolderInfo) {
	for _, v := range li {
		if v.IsFolder {
			// os.MkdirAll(path.Join(boot.Config.BuildDir, "data", v.FilePath), 0755)
			createFileData(v.Files)
			continue
		}

		fn := strings.Replace(v.FilePath, boot.Config.FilePath + "/", "", 1)

		fmt.Println(fn)

		id := base64.StdEncoding.EncodeToString([]byte(fn))
		data, err := handler.PageDetail(filepath.Join(boot.Config.FilePath, v.FilePath), false)
		if err != nil {
			panic(err)
		}

		byt, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		fmt.Println("NAME: ",fn)
		outFile := path.Join(boot.Config.BuildDir, "data", "call", fmt.Sprintf("%v.json", id))
		err = os.WriteFile(outFile, byt, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func updateFileList(li []util.FileFolderInfo) []util.FileFolderInfo {
	for i, v := range li {
		li[i].FilePath = strings.Replace(v.FilePath, boot.Config.FilePath+"/", "", 1)
		li[i].Files = updateFileList(li[i].Files)
	}
	return li
}

func filesData() {

	os.MkdirAll(path.Join(boot.Config.BuildDir, "data"), 0755)
	os.MkdirAll(path.Join(boot.Config.BuildDir, "data", "call"), 0755)

	location := path.Join(boot.Config.BuildDir, "data", "files.json")

	list, err := util.FileList(boot.Config.FilePath)
	if err != nil {
		panic(err)
	}

	list = updateFileList(list)

	createFileData(list)

	byt, err := json.Marshal(map[string]interface{}{"data": list})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(location, byt, 0755)
	if err != nil {
		panic(err)
	}

}
