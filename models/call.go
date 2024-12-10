package models

type FrontMatter struct {
	Method  HTTPMethod        `yaml:"method"`
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
	Body    map[string]string `yaml:"body"`
}

type APIResponse struct {
	Headers    map[string]string `yaml:"headers"`
	Body       interface{}       `yaml:"body"`
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
}
