package main

type Gear struct {
	Chainring, Cog float32
	Wheel          Wheel
}

func NewGear(chainring, cog float32, wheel Wheel) Gear {
	return Gear{chainring, cog, wheel}
}

func (g Gear) Ratio() float32 {
	return g.Chainring / g.Cog
}

func (g Gear) GearInches() float32 {
	return g.Ratio() * g.Wheel.Diameter()
}
