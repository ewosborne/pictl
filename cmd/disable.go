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
	Use:   "disable",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disable called")
		d, _ := cmd.Flags().GetString("delay")
		fmt.Println("delay", d)
		core.Picmd(core.Toggle{Command: "disable", Delay: d})

	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
	// TODO: pass in an optional time delay flag, string, default to 0
	disableCmd.Flags().StringP("delay", "d", "0", "temp time delay")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
