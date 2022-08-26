package cmd

import (
	"encoding/json"
	"gooop/composition"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(composeCmd)
}

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Demonstrates composing complex objects in Go",
	Long:  `Demonstrates composing complex objects in Go`,
	Run: func(cmd *cobra.Command, args []string) {

		// create a wheel object
		wheel := composition.Wheel{Size: 26}

		// create the set of gears
		chainring := composition.Gear{Teeth: 32}
		cog := composition.Gear{Teeth: 11}

		// create the Bike object
		// inject the dependencies
		bike := composition.Bike{
			FrontWheel: wheel,
			RearWheel:  wheel,
			Chainring:  chainring,
			Cog:        cog,
		}

		js, _ := json.MarshalIndent(bike, "", "  ")
		log.Print(js)

	},
}
