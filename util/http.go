package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/kolukattai/kurl/models"
)

func HTTPClient(param *models.FrontMatter, config *models.Config) (*models.APIResponse, error) {

	url := param.URL
	method := param.Method

	var payloadBody bytes.Buffer

	if param.Body != nil {
		byt, err := json.Marshal(param.Body)
		if err != nil {
			panic(err)
		}
		payloadBody.Write(byt)
		// payloadBody = strings.NewReader(string(byt))
	}
	if param.FormData != nil {
		pBody, err := UpdateFormData(param, config)
		if err != nil {
			panic(err)
		}
		payloadBody = pBody
	}

	client := &http.Client{}
	req, err := http.NewRequest(string(method), url, &payloadBody)

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
	response.Request = *param

	return response, nil
}

func UpdateFormData(param *models.FrontMatter, config *models.Config) (requestBody bytes.Buffer, err error) {

	// Create a buffer to store the multipart form data
	// var requestBody bytes.Buffer
	// Create a multipart writer for the buffer
	writer := multipart.NewWriter(&requestBody)

	for _, v := range param.FormData {
		switch v.Type {
		case models.FormDataTypeText:
			er := writer.WriteField(v.Key, v.Value)
			if er != nil {
				return requestBody, er
			}
		case models.FormDataTypeFile:
			// Open the file
			file, er := os.Open(v.File)
			if er != nil {
				return requestBody, er
			}
			defer file.Close()

			// Create the form file field, add the file
			formFile, er := writer.CreateFormFile(v.Key, v.File)
			if er != nil {
				return requestBody, er
			}

			// Copy the content of the file into the form file
			_, er = io.Copy(formFile, file)
			if er != nil {
				return requestBody, er
			}

			// Close the writer to finalize the form data
			er = writer.Close()
			if er != nil {
				return requestBody, er
			}
		}

	}

	param.Headers["Content-Type"] = writer.FormDataContentType()

	return
}
