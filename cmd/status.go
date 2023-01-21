/*
Copyright © 2023 Eric Osborne
No header.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"pictl/core"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status from pihole",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		x := core.NewToggle(host)
		x.Command = "status"
		x.Host = host
		core.Picmd(x)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
