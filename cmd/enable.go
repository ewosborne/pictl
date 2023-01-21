/*
Copyright © 2023 Eric Osborne
No header.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"pictl/core"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable pihole ad blocking",
	Run: func(cmd *cobra.Command, args []string) {
		core.Picmd(core.NewToggle(cmd, "enable"))
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
