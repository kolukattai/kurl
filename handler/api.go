package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
)

func GetDrawerData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths, err := util.FileList(boot.Config.Path)
		if err != nil {
			fmt.Println("path",boot.Config.Path)
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
		fmt.Println(1,err.Error())
		return nil, err
	}

	fmt.Println(string(fileName))

	title, err := util.GetFileName(string(fileName))
	if err != nil {
		fmt.Println(2,err.Error())
		return nil, err
	}

	responseList := util.GetSavedResponse(frontMatter.RefID)

	data = map[string]interface{}{
		"request":  frontMatter,
		"docs":     string(util.MdToHTML([]byte(documentation))),
		"response": responseList,
		"name":     title,
	}

	return data, nil
}
