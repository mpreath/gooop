package listing2_1

import "fmt"

func Run() {
	// Listing 2.1
	fmt.Println("Listing 2.1 >>>>>>>>>")

	cog := 52
	chainring := 11
	ratio := float32(cog) / float32(chainring)

	fmt.Printf("Ratio: %.2f\n", ratio)

	cog = 30
	chainring = 27
	ratio = float32(cog) / float32(chainring)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
