package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Runs various challenges for CryptoPals challenges",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Base command doesn't do anything
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
