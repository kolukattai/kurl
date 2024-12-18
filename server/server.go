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

	http.Handle("GET /", handler.HomePage())

	http.Handle("GET /static/", staticFS)

	log.Fatal(http.ListenAndServe(p, nil))
}
