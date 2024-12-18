package functions

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
)

func Log(fileName string) {

	fm, _, err := util.GetFileData(fileName, boot.Config, true, false)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Request %v %v\n", fm.Method, fm.URL)
	if fm.Body != nil {
		byt, err := json.Marshal(fm.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Request Body", string(byt))
	}

	responseList := util.GetSavedResponse(fm.RefID)

	for i, resp := range responseList {
		log.Info("Saved Response " + fmt.Sprint(i+1))
		fmt.Println("\nHeaders\n---------------------")

		for k, v := range resp.Headers {
			fmt.Println(k, v)
		}

		fmt.Println("Status: ", resp.Status)

		if len(resp.Cookies) > 0 {
			fmt.Println("\n Cookies\n---------------------")
			for k, v := range resp.Cookies {
				fmt.Printf("key: %v, Value: %v\n", k, v)
			}
		}

		if resp.Body != "" {
			fmt.Println("\nBody\n---------------------")
			fmt.Println(resp.BodyStr)
		}
	}

}
