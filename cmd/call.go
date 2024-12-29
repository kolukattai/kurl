/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kolukattai/kurl/boot"
	fu "github.com/kolukattai/kurl/functions"
	"github.com/spf13/cobra"
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configName, err := cmd.Root().Flags().GetString("file")
		if err != nil {
			panic(err)
		}

		ctx, err := cmd.Root().Flags().GetString("context")
		if err != nil {
			panic(err)
		}

		saveResponse, err := cmd.Flags().GetString("save")
		if err != nil {
			panic(err)
		}

		boot.UpdateConfig(configName, ctx)

		if len(args) == 0 {
			panic("file name is missing")
		}

		fu.Call(args[0], saveResponse)
	},
}

func init() {
	rootCmd.AddCommand(callCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// callCmd.Flags().StringP("file", "f", "config.yaml", "configuration file name")
	// callCmd.Flags().StringP("context", "c", ".", "environment file location")
	callCmd.Flags().StringP("save", "s", "", "save requested response")

}
