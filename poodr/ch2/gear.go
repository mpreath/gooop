package main

type Gear struct {
	chainring, cog float32
	wheel          Wheel
}

func NewGear(chainring, cog float32, wheel Wheel) Gear {
	return Gear{chainring, cog, wheel}
}

func (g Gear) Ratio() float32 {
	return g.chainring / g.cog
}

func (g Gear) GearInches() float32 {
	return g.Ratio() * g.wheel.Diameter()
}
