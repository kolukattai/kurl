package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
)

func GetDrawerData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths, err := util.FileList(boot.Config.FilePath)
		if err != nil {
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

func GetPageDetailData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		fileName, err := base64.StdEncoding.DecodeString(id)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}

		fName := strings.Replace(string(fileName), boot.Config.FilePath, "", 1)

		frontMatter, documentation, err := util.GetFileData(fName, boot.Config, true)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}

		title, err := util.GetFileName(string(fileName))
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}

		content := util.GetSavedResponse(frontMatter.RefID)

		data := map[string]interface{}{
			"request":  frontMatter,
			"docs":     string(util.MdToHTML([]byte(documentation))),
			"response": content,
			"name":     title,
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)
	})
}
