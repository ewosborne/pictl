/*
Copyright © 2023 Eric Osborne
No header.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"pictl/core"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable [{-t | --time} time]",
	Short: "Disable adblocking. Optional time parameter.",
	Run: func(cmd *cobra.Command, args []string) {
		core.Picmd(core.NewCliArgs(cmd, "disable"))
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
	disableCmd.Flags().IntP("time", "t", 0, "disable time in seconds")
}
