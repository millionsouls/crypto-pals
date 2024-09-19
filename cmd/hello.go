package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Define the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints hello message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

// Init function to add the hello command to the root command
func init() {
	rootCmd.AddCommand(helloCmd)
}
