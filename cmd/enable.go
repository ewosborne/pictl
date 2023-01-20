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

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable pihole ad blocking",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enable called")
		core.Picmd(core.Toggle{Command: "enable"})
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
