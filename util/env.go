package util

import (
	"fmt"
	"io/fs"
	"log"
	"os"

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
