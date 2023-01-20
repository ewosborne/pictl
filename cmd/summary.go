/*
Copyright © 2023 Eric Osborne
No header.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"pictl/core"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Get summary from the pihole",
	Run: func(cmd *cobra.Command, args []string) {
		core.Picmd(core.Toggle{Command: "summaryRaw"})
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

}
