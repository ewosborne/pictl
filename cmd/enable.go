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
		host, _ := cmd.Flags().GetString("host")
		x := core.NewToggle(host)
		x.Command = "enable"
		core.Picmd(x)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
