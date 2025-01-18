package util

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/models"
	"gopkg.in/yaml.v2"
)

func PrintEnv(stru interface{}) {
	yamlData, err := yaml.Marshal(&stru)

	if err != nil {
		log.Fatalf("Error while Marshaling. %v", err)
	}

	err = os.WriteFile(".env.example.yml", yamlData, fs.ModeAppend)
	if err != nil {
		log.Fatalf("Error creating file. %v", err)
	}
	fmt.Println("environment file name .env.example.yml is create in base path")
	fmt.Println(string(yamlData))
}

func UpdateFrontMatterWithEnvVariable(fm *models.FrontMatter) error {

	url := fm.URL

	// updating path params
	if fm.Params != nil {
		for key, val := range fm.Params {
			url = strings.ReplaceAll(url, fmt.Sprintf("{{%v}}", key), fmt.Sprint(val))
		}
	}

	// updating query params
	if fm.QueryParams != nil {
		for key, val := range fm.QueryParams {
			url = strings.ReplaceAll(url, fmt.Sprintf("{{%v}}", key), fmt.Sprint(val))
		}
	}

	// update url
	fm.URL = url

	fmByt, err := json.Marshal(fm)
	if err != nil {
		return err
	}

	fmStr := UpdateEnvVariable(string(fmByt))

	err = json.Unmarshal([]byte(fmStr), &fm)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEnvVariable(input string) string {
	if boot.Config.EnvVariables != nil {
		for key, value := range boot.Config.EnvVariables {
			if len(value) > 0 {
				input = strings.ReplaceAll(input, fmt.Sprintf("{{%v}}", key), value)
			}
		}
	}
	return input
}
