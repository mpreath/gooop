package cmd

import (
	"gooop/poodr/ch2"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commandCmd)
}

var commandCmd = &cobra.Command{
	Use:   "ch2",
	Short: "Demonstrates examples from chapter 2 in POODR",
	Long:  `Demonstrates examples from chapter 2 in POODR`,
	Run: func(cmd *cobra.Command, args []string) {

		gear := ch2.Gear{}
		log.Printf("%T\n", gear)
	},
}
