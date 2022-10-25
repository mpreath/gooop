package gear

import (
	"gooop/poodr/ch3/wheel"
)

type Gear struct {
	chainring int
	cog       int
	wheel     *wheel.Wheel
}

func New(chainring int, cog int, wheel *wheel.Wheel) *Gear {
	return &Gear{
		chainring: chainring,
		cog:       cog,
		wheel:     wheel,
	}
}

func (g *Gear) Chainring() int {
	return g.chainring
}

func (g *Gear) SetChainring(chainring int) {
	g.chainring = chainring
}

func (g *Gear) Cog() int {
	return g.cog
}

func (g *Gear) SetCog(cog int) {
	g.cog = cog
}

func (g *Gear) Wheel() *wheel.Wheel {
	return g.wheel
}

func (g *Gear) SetWheel(wheel *wheel.Wheel) {
	g.wheel = wheel
}

func (g *Gear) Ratio() float32 {
	return float32(g.chainring) / float32(g.cog)
}

func (g *Gear) GearInches() float32 {
	return g.Ratio() * g.Diameter()
}

func (g *Gear) Diameter() float32 {
	return g.Wheel().Diameter()
}
