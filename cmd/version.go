package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of geekhaku-cli",
	Long:  "All software has versions. This is geekhaku-cli's",
	Run: func(cmd *cobra.Command, args []string) {
		// ここに処理を書いていく
		fmt.Println("version 0.1 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
