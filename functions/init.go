package functions

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/models"
	"gopkg.in/yaml.v2"
)

func Init(name string) {
	if name == "." {
		name = ""
	} else {
		err := os.Mkdir(name, 0744)
		if err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
	}

	file, err := boot.StaticFolder.ReadFile("static/template/README.md")
	if err != nil {
		panic(err)
	}


	conf := models.Config{
		Path: "api",
		Title: "API Documentation",
		EnvVariables: map[string]string{
			"BASE": "http://localhost:9000",
		},
		Build: "dist",
	}

	boot.Config = &conf

	err = os.MkdirAll(filepath.Join(name, conf.Path), 0744)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(name, conf.Path, "README.md"), file, 0744)
	if err != nil {
		panic(err)
	}

	byt, err := yaml.Marshal(conf)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(name, "config.yaml"), byt, 0744)
	if err != nil {
		panic(err)
	}

	AddNewCall("my-first-api-call", name)
}
