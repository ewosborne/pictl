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

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Get summary from the pihole",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		fmt.Println("host is", host)
		x := core.NewToggle(host)
		x.Command = "summaryRaw"
		core.Picmd(core.Toggle(x))
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

}
