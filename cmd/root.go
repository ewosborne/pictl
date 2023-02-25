/*
Copyright © 2023 Eric Osborne
No header.
*/
package cmd

import (
	"golang.org/x/exp/slog"
	"os"

	"github.com/spf13/cobra"
)

var Version = "0.1.1"

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

	// check rootCmd to see if verbose is set?
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	textHandler := opts.NewTextHandler(os.Stderr)
	slog.SetDefault(slog.New(textHandler))
	verbose, _ := rootCmd.Flags().GetBool("verbose")
	if verbose == true {
		print("verbose is on")
	}
}

func init() {

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP("host", "", "pi.hole", "host to reach")
}
