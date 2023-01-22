/*
Copyright © 2023 Eric Osborne
No header.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Version = "0.0.12"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pictl",
	Long:    `pictl is a CLI tool to interact with your pi.hole API.`,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP("host", "", "pi.hole", "host to reach")
}
