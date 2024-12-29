/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kolukattai/kurl/boot"
	fu "github.com/kolukattai/kurl/functions"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log request and saved response",
	Long:  `log api call with request and saved response`,
	Run: func(cmd *cobra.Command, args []string) {
		configName, err := cmd.Root().Flags().GetString("file")
		if err != nil {
			panic(err)
		}

		ctx, err := cmd.Root().Flags().GetString("context")
		if err != nil {
			panic(err)
		}

		boot.UpdateConfig(configName, ctx)

		if len(args) == 0 {
			panic("file name is missing")
		}

		fu.Log(args[0])
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// logCmd.Flags().BoolP("save", "s", false, "save requested response")

}
