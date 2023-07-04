package main

import "fmt"

func main() {
	gear := NewGear(52, 11, 26, 1.5)
	fmt.Printf("Circumference:\t%f\n", gear.Circumference())
	fmt.Printf("Gear Inches:\t%f\n", gear.GearInches())
	fmt.Printf("Gear Ratio:\t%f\n", gear.Ratio())
}
