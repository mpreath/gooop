package main

func main() {
	wheel := NewWheel(26, 1.5)
	println(wheel.Circumference())

	gear := NewGear(52, 11, wheel)
	println(gear.GearInches())
	println(gear.Ratio())
}
