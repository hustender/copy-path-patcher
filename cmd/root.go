package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "copy-pp",
	Short: "copy-pp is a utility to copy directories or a single file and update old path strings in text-based files with new path.",
	Long:  `Copy-Path-Patcher (copy-pp) is a powerful and efficient tool designed to simplify the process of copying entire directories or a single file while seamlessly updating path references within text-based files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Usage: \n" +
			"  copy-pp <mode> <target>\n")
		fmt.Printf("Modes: \n" +
			"  dir (copy a directory)\n" +
			"  file (copy a file)\n")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
