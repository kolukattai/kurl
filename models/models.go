package models

import "strings"

type HTTPMethod string

func (st *HTTPMethod) Parse() string {
	return strings.ToUpper(string(*st))
}

const (
	HTTPMethodGET     HTTPMethod = "GET"
	HTTPMethodPUT     HTTPMethod = "PUT"
	HTTPMethodPOST    HTTPMethod = "POST"
	HTTPMethodPATCH   HTTPMethod = "PATCH"
	HTTPMethodDELETE  HTTPMethod = "DELETE"
	HTTPMethodOPTIONS HTTPMethod = "OPTIONS"
)

type HTPClientConf struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
	Method  HTTPMethod        `json:"method"`
}

type EnvVariables struct {
	Name      string            `json:"name"`
	Variables map[string]string `json:"variables"`
}

type Config struct {
	FilePath     string         `json:"file_folder"`
	EnvVariables []EnvVariables `json:"env_variables"`
	DefaultEnv   int            `json:"selected_env"`
}
