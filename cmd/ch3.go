package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ch3Cmd)
}

var ch3Cmd = &cobra.Command{
	Use:   "ch3",
	Short: "Demonstrates examples from chapter 3 in POODR",
	Long:  `Demonstrates examples from chapter 3 in POODR`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
