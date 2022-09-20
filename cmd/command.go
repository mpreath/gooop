package cmd

import (
	"gooop/poodr/ch1"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commandCmd)
}

var commandCmd = &cobra.Command{
	Use:   "ch1",
	Short: "Demonstrates examples from chapter 1 in POODR",
	Long:  `Demonstrates examples from chapter 1 in POODR`,
	Run: func(cmd *cobra.Command, args []string) {

		gear := ch1.Gear{}
		log.Printf("%T\n", gear)
	},
}
