package models

type FrontMatter struct {
	RefID   string            `yaml:"refID" json:"refID"`
	Name    string            `yaml:"name" json:"name"`
	Method  HTTPMethod        `yaml:"method" json:"method"`
	URL     string            `yaml:"url" json:"url"`
	Headers map[string]string `yaml:"headers" json:"headers"`
	Body    map[string]string `yaml:"body" json:"body"`
}

type APIResponse struct {
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
	BodyStr    string            `json:"body_str"`
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
	Cookies    []string          `json:"cookies"`
}
