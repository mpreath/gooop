package cmd

import (
	"gooop/interfaces"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(interfaceCmd)
}

var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "Demonstrates using interfaces in Go",
	Long:  `Demonstrates using interfaces in Go`,
	Run: func(cmd *cobra.Command, args []string) {

		// create a road bike first
		roadBike := interfaces.RoadBike{
			BaseBike: interfaces.BaseBike{},
		}

		// setting the road bike's front wheel to a roadwheel
		// interface arguments assume a pointer rather than
		// referenced value
		roadBike.SetFrontWheel(&interfaces.RoadWheel{})
		roadBike.GetFrontWheel().SetSize(26)
		roadBike.SetRearWheel(&interfaces.RoadWheel{})
		roadBike.GetRearWheel().SetSize(26)

		log.Printf("[%T] Front Wheel [%T]: %d, Rear Wheel [%T]: %d",
			roadBike, roadBike.GetFrontWheel(), roadBike.GetFrontWheel().GetSize(),
			roadBike.GetRearWheel(), roadBike.GetRearWheel().GetSize())

		mountainBike := interfaces.MountainBike{
			BaseBike: interfaces.BaseBike{},
		}

		mountainBike.SetFrontWheel(&interfaces.OffroadWheel{})
		mountainBike.GetFrontWheel().SetSize(29)
		mountainBike.SetRearWheel(&interfaces.OffroadWheel{})
		mountainBike.GetRearWheel().SetSize(29)

		log.Printf("[%T] Front Wheel [%T]: %d, Rear Wheel [%T]: %d",
			mountainBike, mountainBike.GetFrontWheel(), mountainBike.GetFrontWheel().GetSize(),
			mountainBike.GetRearWheel(), mountainBike.GetRearWheel().GetSize())

	},
}
