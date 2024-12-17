package util

import (
	"bytes"
	"embed"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(templates embed.FS, w http.ResponseWriter, tmpl string, data interface{}) (res string) {
	// Parse and execute the template using embedded files
	t, err := template.ParseFS(templates, "templates/layout/base.html", "templates/"+tmpl)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	var result bytes.Buffer

	err = t.Execute(&result, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
	}

	res = result.String()

	if w != nil {
		// Render the template with the provided data
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Println("Error rendering template:", err)
		}
	}

	return res
}
