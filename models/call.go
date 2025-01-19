package models

type FrontMatterContent struct {
	Name    string `yaml:"name" json:"name"`
	Content string `yaml:"content" json:"content"`
}

type FormDataType string

const (
	FormDataTypeText FormDataType = "text"
	FormDataTypeFile FormDataType = "file"
)

type FormData struct {
	Key   string       `yaml:"key" json:"key"`
	Type  FormDataType `yaml:"type" json:"type"`
	File  string       `yaml:"file" json:"file"`
	Value string       `yaml:"value" json:"value"`
}

type FrontMatter struct {
	RefID       string                 `yaml:"refID" json:"refID"`
	Name        string                 `yaml:"name" json:"name"`
	Method      HTTPMethod             `yaml:"method" json:"method"`
	URL         string                 `yaml:"url" json:"url"`
	Headers     map[string]string      `yaml:"headers" json:"headers"`
	Body        map[string]string      `yaml:"body" json:"body"`
	Title       string                 `yaml:"title" json:"title"`
	Content     []FrontMatterContent   `yaml:"content" json:"content"`
	Params      map[string]interface{} `yaml:"params" json:"params"`
	QueryParams map[string]interface{} `yaml:"queryParams" json:"queryParams"`
	FormData    []FormData             `yaml:"formData" json:"formData"`
}

type APIResponse struct {
	Headers    map[string]string      `json:"headers"`
	Body       interface{}            `json:"body"`
	BodyStr    string                 `json:"body_str"`
	StatusCode int                    `json:"status_code"`
	Status     string                 `json:"status"`
	Cookies    []string               `json:"cookies"`
	Request    FrontMatter            `json:"request"`
	Templates  map[string]interface{} `json:"templates"`
}
