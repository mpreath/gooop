package cmd

import (
	"gooop/factory"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(factoryCmd)
}

var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "Demonstrates using interfaces and factories in Go",
	Long:  `Demonstrates using interfaces and factories in Go`,
	Run: func(cmd *cobra.Command, args []string) {

		// create a road bike first
		roadBike := factory.BikeFactory("road")

		log.Printf("[%T] Front Wheel [%T]: %d, Rear Wheel [%T]: %d",
			roadBike, roadBike.GetFrontWheel(), roadBike.GetFrontWheel().GetSize(),
			roadBike.GetRearWheel(), roadBike.GetRearWheel().GetSize())

		mountainBike := factory.BikeFactory("mountain")

		log.Printf("[%T] Front Wheel [%T]: %d, Rear Wheel [%T]: %d",
			mountainBike, mountainBike.GetFrontWheel(), mountainBike.GetFrontWheel().GetSize(),
			mountainBike.GetRearWheel(), mountainBike.GetRearWheel().GetSize())

	},
}
