package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/handler"
)

func RunDoc(port string) {

	staticFS := http.StripPrefix("/", http.FileServer(http.FS(boot.StaticFolder)))

	p := fmt.Sprintf(":%v", port)

	http.Handle("GET /data/files.json", handler.GetDrawerData())
	http.Handle("GET /data/call/{id}", handler.GetPageDetailData())
	http.Handle("GET /data/env.json", handler.GetEnv())
	http.Handle("POST /data/env.json", handler.UpdateEnv())

	http.Handle("GET /", handler.HomePage())
	http.Handle("DELETE /saved/{id}/{index}", handler.DeleteSavedResponse())

	http.Handle("GET /static/", staticFS)

	log.Fatal(http.ListenAndServe(p, nil))
}
