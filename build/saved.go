package build

import (
	"encoding/json"
	"os"

	"github.com/kolukattai/kurl/models"
	"github.com/kolukattai/kurl/util"
)

func saveFile() {

	saved, err := util.FileList(".saved")
	if err != nil {
		return
	}

	for _, v := range saved {
		byt, err := os.ReadFile(v.FilePath)
		if err != nil {
			panic(err)
		}

		file, err := util.GZip().UnPack(byt)
		if err != nil {
			panic(err)
		}

		var data []*models.APIResponse

		err = json.Unmarshal(file, &data)
		if err != nil {
			panic(err)
		}

	}

}
