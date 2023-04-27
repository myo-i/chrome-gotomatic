package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpmd = &cobra.Command{
	Use:   "help",
	Short: "Print command help",
	Run: func(cmd *cobra.Command, args []string) {
		// ここに処理を書いていく
		fmt.Println("ascii             show ASCII art")
		fmt.Println("version           show CLI version")
	},
}

func init() {
	rootCmd.AddCommand(helpmd)
}
