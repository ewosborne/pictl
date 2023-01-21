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
	Use:   "disable [{-d | --delay} time]",
	Short: "Disable adblocking. Optional time parameter.",
	Run: func(cmd *cobra.Command, args []string) {
		core.Picmd(core.NewToggle(cmd, "disable"))
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
	disableCmd.Flags().IntP("delay", "d", 0, "disable time in seconds")
}
