package listing2_3

import "fmt"

type Gear struct {
	cog       int
	chainring int
	rim       int
	tire      float32
}

func NewGear(chainring int, cog int, rim int, tire float32) *Gear {
	return &Gear{
		chainring: chainring,
		cog:       cog,
		rim:       rim,
		tire:      tire,
	}
}

func (g *Gear) Ratio() float32 {
	return float32(g.chainring) / float32(g.cog)
}

func (g *Gear) GearInches() float32 {
	return g.Ratio() * (float32(g.rim) + (g.tire * 2))
}

func Run() {

	fmt.Println("List 2.3 >>>>>>>>>>")
	fmt.Printf("Gear Inches: %.2f\n", NewGear(52, 11, 26, 1.5).GearInches())
	fmt.Printf("Gear Inches: %.2f\n", NewGear(52, 11, 24, 1.25).GearInches())
}
