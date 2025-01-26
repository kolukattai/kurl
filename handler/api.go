package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/functions"
	"github.com/kolukattai/kurl/models"
	"github.com/kolukattai/kurl/util"
	"gopkg.in/yaml.v2"
)

func GetDrawerData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths, err := util.FileList(boot.Config.Path)
		if err != nil {
			fmt.Println("path", boot.Config.Path)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": paths})
	})
}

func GetEnv() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(boot.Config)
	})
}

func MakeCall() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		reqObj := models.FrontMatter{}

		err = json.Unmarshal(val, &reqObj)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		id := r.PathValue("name")

		fileName, err := base64.StdEncoding.DecodeString(id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		resp := functions.Call(string(fileName), "")

		respByt, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write(respByt)
	})
}

func UpdateRequest() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		reqObj := models.FrontMatter{}

		err = json.Unmarshal(val, &reqObj)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		id := r.PathValue("name")

		fmByt, err := yaml.Marshal(reqObj)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		fileName, err := base64.StdEncoding.DecodeString(id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		fName := strings.Replace(string(fileName), boot.Config.Path, "", 1)

		_, docs, err := util.GetFileData(fName, boot.Config, true, false)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		comments := `# do not change refID, this key is used to connect this api with it's saved response`

		fileVal := fmt.Sprintf("---\n%v\n%v\n---\n\n%v", comments, string(fmByt), docs)

		fmt.Println("DA", fName)
		err = os.WriteFile(filepath.Join(boot.Config.Path, fName), []byte(fileVal), 0755)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		// _, docs, err := util.GetFileData(id, boot.Config, true, true)
		// if err != nil {
		// 	w.WriteHeader(400)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		w.WriteHeader(200)
		w.Write([]byte(fileVal))

	})
}

func UpdateEnv() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		val, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		envObj := map[string]string{}

		err = json.Unmarshal(val, &envObj)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		boot.Config.EnvVariables = envObj

		byt, err := yaml.Marshal(boot.Config)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		err = os.WriteFile("config.yaml", byt, 0744)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(boot.Config)
	})
}

func DeleteSavedResponse() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		indexStr := r.PathValue("index")

		index, err := strconv.Atoi(indexStr)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}

		saved := util.DeleteSaved(id, index)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(saved)
	})
}

func GetPageDetailData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := strings.Replace(r.PathValue("id"), ".json", "", 1)

		data, err := PageDetail(id, true)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)
	})
}

func PageDetail(id string, base64Encoded bool) (data map[string]interface{}, err error) {
	fileName := []byte(id)

	if base64Encoded {
		fileName, err = base64.StdEncoding.DecodeString(id)
		if err != nil {
			return nil, err
		}
	}

	fName := strings.Replace(string(fileName), boot.Config.Path, "", 1)

	// skipFm := fName == "README.md" || fName == "index.md" || fName == "/README.md"

	fmt.Println(fName)

	frontMatter, documentation, err := util.GetFileData(fName, boot.Config, true, false)
	if err != nil {
		fmt.Println(1, err.Error())
		return nil, err
	}

	fmt.Println(string(fileName))

	responseList := util.GetSavedResponse(frontMatter.RefID)

	data = map[string]interface{}{
		"request":  frontMatter,
		"docs":     string(util.MdToHTML([]byte(documentation))),
		"response": responseList,
	}

	return data, nil
}
