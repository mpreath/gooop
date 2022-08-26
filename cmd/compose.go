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

		// create the Bike object
		// inject the dependencies
		bike := composition.Bike{
			FrontWheel: composition.Wheel{Size: 26},
			RearWheel:  composition.Wheel{Size: 26},
			Chainring:  composition.Gear{Teeth: 32},
			Cog:        composition.Gear{Teeth: 11},
		}

		log.Printf("[%T] Gear ratio: %f", bike, bike.GearRatio())

		js, _ := json.MarshalIndent(bike, "", "  ")
		log.Print(string(js))

		// create a mountainbike object
		mountainBike := composition.MountainBike{
			Bike: composition.Bike{
				FrontWheel: composition.Wheel{Size: 29},
				RearWheel:  composition.Wheel{Size: 29},
				Chainring:  composition.Gear{Teeth: 38},
				Cog:        composition.Gear{Teeth: 24},
			},
		}

		log.Printf("[%T] Gear ratio: %f", mountainBike, mountainBike.GearRatio())

		js, _ = json.MarshalIndent(mountainBike, "", "  ")
		log.Print(string(js))

	},
}
