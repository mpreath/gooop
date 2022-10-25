package gear

type Gear struct {
	chainring int
	cog       int
	wheel     WheelInterface
}

type WheelInterface interface {
	Diameter() float32
}

func New(chainring int, cog int, wheel WheelInterface) *Gear {
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

func (g *Gear) Wheel() WheelInterface {
	return g.wheel
}

func (g *Gear) SetWheel(wheel WheelInterface) {
	g.wheel = wheel
}

func (g *Gear) Ratio() float32 {
	return float32(g.chainring) / float32(g.cog)
}

func (g *Gear) GearInches() float32 {
	return g.Ratio() * g.diameter()
}

func (g *Gear) diameter() float32 {
	return g.Wheel().Diameter()
}
