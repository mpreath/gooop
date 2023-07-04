package main

type Gear struct {
	chainring, cog float32
	wheel          *Wheel
}

func NewGear(chainring, cog, rim, tire float32) *Gear {
	return &Gear{chainring, cog, NewWheel(rim, tire)}
}

func (g *Gear) Ratio() float32 {
	return g.chainring / g.cog
}

func (g *Gear) GearInches() float32 {
	return g.Ratio() * g.wheel.Diameter()
}

func (g *Gear) Circumference() float32 {
	return g.wheel.Circumference()
}
