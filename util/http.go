package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kolukattai/kurl/models"
)

func HTTPClient(param models.FrontMatter, config *models.Config) (*models.APIResponse, error) {

	url := param.URL
	method := param.Method

	var payloadBody io.Reader

	if param.Body != nil {
		byt, err := json.Marshal(param.Body)
		if err != nil {
			panic(err)
		}
		payloadBody = strings.NewReader(string(byt))
	}

	client := &http.Client{}
	req, err := http.NewRequest(string(method), url, payloadBody)

	if err != nil {
		fmt.Println("ERR", err)
		return nil, err
	}

	for key, value := range param.Headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERR", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERR", err)
		return nil, err
	}
	defer res.Body.Close()

	var rBody interface{}

	err = json.Unmarshal(body, &rBody)
	if err != nil {
		rBody = string(body)
	}

	response := &models.APIResponse{}
	response.Status = res.Status
	response.StatusCode = res.StatusCode
	response.Body = rBody
	response.BodyStr = string(body)
	response.Cookies = []string{}

	for _, v := range res.Cookies() {
		response.Cookies = append(response.Cookies, v.Raw)
	}

	resHeader := map[string]string{}

	for k, v := range res.Header {
		resHeader[k] = strings.Join(v, ",")
	}

	response.Headers = resHeader

	return response, nil
}
