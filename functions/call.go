package functions

import (
	"encoding/json"
	"fmt"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
)

func Call(fileName string, saveResponse string) {

	fm, _, err := util.GetFileData(fileName, boot.Config, false, false)

	if err != nil {
		panic(err)
	}

	err = util.UpdateFrontMatterWithEnvVariable(&fm)

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

	resp, err := util.HTTPClient(&fm, boot.Config)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\nResponse Headers\n---------------------")

	for k, v := range resp.Headers {
		fmt.Println(k, v)
	}

	fmt.Println("Response Status: ", resp.Status)

	if len(resp.Cookies) > 0 {
		fmt.Println("\nResponse Cookies\n---------------------")
		for k, v := range resp.Cookies {
			fmt.Printf("key: %v, Value: %v\n", k, v)
		}
	}

	if resp.Body != "" {
		fmt.Println("\nResponse Body\n---------------------")
		fmt.Println(resp.BodyStr)
	}

	if len(saveResponse) > 0 {
		resp.Request.Name = saveResponse
		err := util.SaveResponse(fm.RefID, resp)
		if err != nil {
			panic(err)
		}

		fmt.Println("Response saved...")
	}

}
