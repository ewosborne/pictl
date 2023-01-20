/*
Copyright © 2023 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pictl/core"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable [{-d | --delay} time]",
	Short: "Disable adblocking. Optional time parameter.",
	Run: func(cmd *cobra.Command, args []string) {
		d, _ := cmd.Flags().GetInt("delay")

		if d == 0 {
			core.Picmd(core.Toggle{Command: "disable"})
		} else {
			core.Picmd(core.Toggle{Command: "disable", Delay: fmt.Sprint(d)})
		}
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
	disableCmd.Flags().IntP("delay", "d", 0, "disable time in seconds")
}
