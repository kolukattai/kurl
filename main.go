/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/kolukattai/kurl/cmd"

	"github.com/kolukattai/kurl/boot"
	fu "github.com/kolukattai/kurl/functions"
)

func Call() {
	boot.UpdateConfig("config.json", "./example")
	fmt.Println(boot.Config)
	fu.Call("example-get-api.md")
}

func main() {
	cmd.Execute()
}
