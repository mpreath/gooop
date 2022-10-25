package main

import (
	"fmt"
	gear2 "gooop/poodr/ch4/examples/gear"
	wheel2 "gooop/poodr/ch4/examples/wheel"
	"gooop/poodr/ch4/gear"
	"gooop/poodr/ch4/gearwrapper"
	"gooop/poodr/ch4/wheel"
)

func main() {

	// #################################################
	// 4. DEPENDENCY MANAGEMENT
	// #################################################

	// wheel and gear example using dependency injection
	// injecting the dependency wheel into gear (w -> g)
	w := wheel.New(26, 1.5)
	fmt.Printf("Circumference: %.2f\n", w.Circumference())

	fmt.Printf("Gear Inches: %.2f\n", gear.New(52, 11, w).GearInches())

	fmt.Printf("Ratio: %.2f\n", gear.New(52, 11, nil).Ratio())

	// with the dependency reversed (g -> w)
	// injecting the dependency gear into wheel
	g2 := gear2.New(52, 11)
	w2 := wheel2.New(26, 1.5, g2)

	fmt.Printf("(Reverse Depedency Wheel) Gear Inches: %.2f Ratio: %.2f\n", w2.GearInches(), g2.Ratio())

	// example using a wrapper to control which gear we return if we have
	// multiple options (e.g. factory pattern)
	g := gearwrapper.Gear(52, 11, w)

	fmt.Printf("(Wrapped Gear) Gear Inches: %.2f Ratio: %.2f\n", g.GearInches(), g.Ratio())

}
