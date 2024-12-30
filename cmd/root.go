package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "copy-pp",
	Short: "copy-pp is a utility to copy directories and update old path strings in text-based files with new path.",
	Long:  `Copy-Path-Patcher (copy-pp) is a powerful and efficient tool designed to simplify the process of copying entire directories while seamlessly updating path references within text-based files.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
