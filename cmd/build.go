/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kolukattai/kurl/boot"
	"github.com/kolukattai/kurl/build"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build api documentation",
	Long:  `build api documentation`,
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

		build.Run()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// buildCmd.Flags().StringP("file", "f", "config.yaml", "configuration file name")
	// buildCmd.Flags().StringP("context", "c", ".", "environment file location")
}
