package gear

type Gear struct {
	chainring int
	cog       int
}

func New(chainring int, cog int) *Gear {
	return &Gear{
		chainring: chainring,
		cog:       cog,
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

func (g *Gear) Ratio() float32 {
	return float32(g.chainring) / float32(g.cog)
}
