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
		core.Picmd(core.NewToggle(cmd, "status"))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
