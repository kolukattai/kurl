package handler

import (
	"net/http"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/util"
)

func HomePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title   string
			Message string
			Develop bool
		}{
			Title:   boot.Config.Title,
			Message: "This is the home page.",
			Develop: true,
		}

		util.RenderTemplate(boot.TemplateFolder, w, "home.html", data)
	})
}
