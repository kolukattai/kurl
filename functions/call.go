package functions

import (
	"fmt"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
)

func Call(fileName string) {
	fm, _, err := util.GetFileData(fileName, boot.Config, false)

	if err != nil {
		panic(err)
	}

	fmt.Println(fm)
}
