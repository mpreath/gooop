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

		// In Go we can create instances as return values of constructors we create
		wheel1 := ch2.NewWheel(26, 1.5)
		log.Printf("Wheel 1: [Rim: %0.1f, Tire: %0.1f, Diameter: %0.1f, Circumference: %0.1f] %T\n", wheel1.Rim(), wheel1.Tire(), wheel1.Diameter(), wheel1.Circumference(), wheel1)
		// Wheel 1: [26.0, 1.5, 29.0, 91.1] *ch2.Wheel

		// Or directly from the struct and then use our accessor methods to set values
		wheel2 := &ch2.Wheel{}
		wheel2.SetRim(26)
		wheel2.SetTire(1.5)
		log.Printf("Wheel 2: [Rim: %0.1f, Tire: %0.1f, Diameter: %0.1f, Circumference: %0.1f] %T\n", wheel2.Rim(), wheel2.Tire(), wheel2.Diameter(), wheel2.Circumference(), wheel2)
		// Wheel 2: [26.0, 1.5, 29.0, 91.1] *ch2.Wheel

		// Create new gear by specifying chainring, cog, and injecting the wheel dependency
		gear1 := ch2.NewGear(52, 11, wheel1)
		log.Printf("Gear 1: [Chainring: %d, Cog: %d, GearInches: %0.1f, Ratio: %0.1f] %T\n", gear1.Chainring(), gear1.Cog(), gear1.GearInches(), gear1.Ratio(), gear1)
		// Gear 1: [52, 11, 137.1, 4.7] *ch2.Gear

		gear2 := ch2.NewGear(50, 11, wheel2)
		log.Printf("Gear 2: [Chainring: %d, Cog: %d, GearInches: %0.1f, Ratio: %0.1f] %T\n", gear2.Chainring(), gear2.Cog(), gear2.GearInches(), gear2.Ratio(), gear2)
		// Gear 2: [50, 11, 131.8, 4.5] *ch2.Gear

		gear3 := ch2.NewGear(60, 6, wheel1)
		log.Printf("Gear 3: [Chainring: %d, Cog: %d, GearInches: %0.1f, Ratio: %0.1f] %T\n", gear3.Chainring(), gear3.Cog(), gear3.GearInches(), gear3.Ratio(), gear3)

	},
}
