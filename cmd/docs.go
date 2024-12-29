/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/server"
	"github.com/spf13/cobra"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "run documentation server",
	Long:  `run a documentation server which will provide api documentation on given port (default port 8080)`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}

		configName, err := cmd.Root().Flags().GetString("file")
		if err != nil {
			panic(err)
		}

		ctx, err := cmd.Root().Flags().GetString("context")
		if err != nil {
			panic(err)
		}

		boot.UpdateConfig(configName, ctx)

		server.RunDoc(port)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// docsCmd.Flags().StringP("file", "f", "config.yaml", "configuration file name")
	// docsCmd.Flags().StringP("context", "c", ".", "environment file location")
	docsCmd.Flags().StringP("port", "p", "8080", "alter document running port")
}
