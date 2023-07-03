package main

import "fmt"

func main() {
	wheel := NewWheel(26, 1.5)
	fmt.Printf("Circumference:\t%f\n", wheel.Circumference())

	gear := NewGear(52, 11, wheel)
	fmt.Printf("Gear Inches:\t%f\n", gear.GearInches())
	fmt.Printf("Gear Ratio:\t%f\n", gear.Ratio())
}
