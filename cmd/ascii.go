package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "ascii",
	Short: "Print the ascii art of cat",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("ascii.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
