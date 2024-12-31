package cmd

import (
	"github.com/spf13/cobra"
)

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Copies a directory and updates path references within text-based files.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
}
