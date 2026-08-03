// Package cmd provides cmd  INFO: root command
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/swampPr/ansify/internal/ansify"
)

var rootCmd = &cobra.Command{
	Use:   "ansify",
	Short: "Turn an image into an ANSI art style",
	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must provide a file to ansify")
			os.Exit(1)
		}

		ansify.ANSIfy(args[0])
	},
}

// Execute function  INFO:  Executes command and child commands
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
