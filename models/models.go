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

type Config struct {
	Path         string            `json:"path" yaml:"path"`
	Title        string            `json:"title" yaml:"title"`
	EnvVariables map[string]string `json:"env" yaml:"env"`
	Build        string            `json:"build" yaml:"build"`
}
