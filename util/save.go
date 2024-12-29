package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kolukattai/kurl/models"
)

func GetSavedResponse(refID string) []*models.APIResponse {
	folderLocation := "./.saved"

	fileLocation := filepath.Join(folderLocation, refID)

	file, err := os.ReadFile(fileLocation)

	if err != nil {
		fmt.Println(err.Error())
		return []*models.APIResponse{}
	}

	file, err = GZip().UnPack(file)

	if err != nil {
		fmt.Println(err.Error())
		return []*models.APIResponse{}
	}

	response := []*models.APIResponse{}

	err = json.Unmarshal(file, &response)

	if err != nil {
		fmt.Println(err.Error())
		return []*models.APIResponse{}
	}

	return response
}

func DeleteSaved(refID string, index int) []*models.APIResponse {
	folderLocation := "./.saved"

	fileLocation := filepath.Join(folderLocation, refID)

	file, err := os.ReadFile(fileLocation)

	if err != nil {
		fmt.Println(err.Error())
		return []*models.APIResponse{}
	}

	file, err = GZip().UnPack(file)

	if err != nil {
		fmt.Println(err.Error())
		return []*models.APIResponse{}
	}

	response := []*models.APIResponse{}

	err = json.Unmarshal(file, &response)

	if err != nil {
		fmt.Println(err.Error())
		return []*models.APIResponse{}
	}

	newResponse := []*models.APIResponse{}
	newResponse = append(newResponse, response[:index]...)
	newResponse = append(newResponse, response[index+1:]...)

	file, err = json.Marshal(newResponse)

	if err != nil {
		panic(err)
	}

	file, err = GZip().Pack(file)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileLocation, file, 0644)

	if err != nil {
		panic(err)
	}

	return newResponse
}

func SaveResponse(refID string, resp *models.APIResponse) error {
	folderLocation := "./.saved"

	err := os.MkdirAll(folderLocation, 0744)

	if err != nil {
		return err
	}

	fileLocation := filepath.Join(folderLocation, refID)

	file, err := os.ReadFile(fileLocation)

	if err != nil {
		file = []byte("")
	}

	file, err = GZip().UnPack(file)
	if err != nil {
		file = []byte("[]")
	}

	fileResp := []*models.APIResponse{}

	err = json.Unmarshal(file, &fileResp)
	if err != nil {
		panic(err)
	}

	fileResp = append(fileResp, resp)

	file, err = json.Marshal(fileResp)

	if err != nil {
		panic(err)
	}

	file, err = GZip().Pack(file)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileLocation, file, 0744)

	if err != nil {
		panic(err)
	}

	return nil
}
