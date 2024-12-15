/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kolukattai/kurl/boot"
	fu "github.com/kolukattai/kurl/functions"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configName, err := cmd.Flags().GetString("file")
		if err != nil {
			panic("configuration file not provided")
		}

		ctx, err := cmd.Flags().GetString("context")
		if err != nil {
			panic("context is not provided")
		}

		boot.UpdateConfig(configName, ctx)

		if len(args) == 0 {
			panic("file name is missing")
		}

		fu.AddNewCall(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("file", "f", "config.json", "configuration file name")
	addCmd.Flags().StringP("context", "c", ".", "environment file location")

}
