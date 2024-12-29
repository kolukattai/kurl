package util

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/kolukattai/kurl/models"
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

func RequestTemplates(req models.FrontMatter) (res map[string]interface{}) {
	res = map[string]interface{}{}
	res["curl"] = CurlTemplate(req)
	res["javascript"] = JavaScriptTemplate(req)
	return
}

func JavaScriptTemplate(req models.FrontMatter) string {
	fmt.Println("called")
	option := map[string]string{
		"method": req.Method.Parse(),
	}

	addHeader := func() {
		headerByt, err := json.MarshalIndent(req.Headers, "  ", "  ")
		if err != nil {
			return
		}
		if req.Headers != nil {
			option["headers"] = string(headerByt)
		}
	}

	addBody := func() {
		if req.Body != nil {
			bodyByt, err := json.MarshalIndent(req.Body, "", "  ")
			if err != nil {
				return
			}
			option["body"] = string(bodyByt)
		}
	}

	addHeader()
	addBody()

	optionStr := fmt.Sprintf("{\nmethod: %v", req.Method.Parse())

	head, ok := option["headers"]
	if ok {
		optionStr += fmt.Sprintf("\theaders: %v", strings.Replace(strings.Replace(head, "\"{", "{\n", 1), "}\"", "\n}", 1))
	}

	val, ok := option["body"]
	if ok {
		optionStr += fmt.Sprintf("\tbody: %v", strings.Replace(strings.Replace(val, "\"{", "{\n", 1), "}\"", "\n}", 1))
	}

	optionStr += "}"

	// optionByt, err := json.Marshal(option)
	// if err != nil {
	// 	optionByt = []byte("")
	// }

	ca := fmt.Sprintf("fetch(\"%v\", %v)\n\t.then(res => console.log(res))\n\t.catch(err => console.error(err))", req.URL, optionStr)

	return ca
}

func CurlTemplate(req models.FrontMatter) string {

	ca := fmt.Sprintf("curl -X %v %v ", req.Method, req.URL)

	for key, val := range req.Headers {
		ca += fmt.Sprintf("-H \"%v: %v\" ", key, val)
	}

	if req.Body != nil {
		byt, err := json.Marshal(req.Body)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			ca += fmt.Sprintf("-d '%v'", string(byt))
		}
	}

	return ca
}
