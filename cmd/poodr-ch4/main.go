package main

import (
	"fmt"
	"gooop/pkg/poodr/ch4/gear"
	"gooop/pkg/poodr/ch4/wheel"
)

func main() {

	w := wheel.New(26, 1.5)
	fmt.Printf("Circumference: %.2f\n", w.Circumference())

	fmt.Printf("Gear Inches: %.2f\n", gear.New(52, 11, w).GearInches())

	fmt.Printf("Ratio: %.2f\n", gear.New(52, 11, nil).Ratio())
}
