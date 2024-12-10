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
		return nil, err
	}

	for key, value := range param.Headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	return &models.APIResponse{}, nil
}
