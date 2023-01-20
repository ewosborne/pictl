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
	Short: "Disable adblocking. Optional time parameter; default is permanent disable.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disable called")
		d, _ := cmd.Flags().GetString("delay")
		fmt.Println("delay", d)
		core.Picmd(core.Toggle{Command: "disable", Delay: string(d)})

	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
	disableCmd.Flags().IntP("delay", "d", 0, "temp time delay")
}
