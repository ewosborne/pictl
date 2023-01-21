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
	Short: "Disable adblocking.  Optional time parameter.",
	Run: func(cmd *cobra.Command, args []string) {
		d, _ := cmd.Flags().GetInt("delay")

		host, _ := cmd.Flags().GetString("host")
		x := core.NewToggle(host)

		x.Command = "disable"

		if d == 0 {
			core.Picmd(x)
		} else {
			x.Delay = fmt.Sprint(d)
			core.Picmd(core.Toggle(x))
		}
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
	disableCmd.Flags().IntP("delay", "d", 0, "disable time in seconds")
}
