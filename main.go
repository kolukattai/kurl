/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"

	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/cmd"
)

//go:embed static/*
var staticFS embed.FS

//go:embed templates/*
var templatesFS embed.FS

func init() {
	boot.StaticFolder = staticFS
	boot.TemplateFolder = templatesFS
}
func main() {
	cmd.Execute()
}
