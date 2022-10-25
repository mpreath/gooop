package listing2_5

import "fmt"

type Gear struct {
	cog       int
	chainring int
}

func NewGear(chainring int, cog int) *Gear {
	return &Gear{
		chainring: chainring,
		cog:       cog,
	}
}

func (g *Gear) Chainring() int {
	return g.chainring
}

func (g *Gear) Cog() int {
	return g.cog
}

func (g *Gear) Ratio() float32 {
	return float32(g.Chainring()) / float32(g.Cog())
}

func Run() {

	fmt.Println("List 2.5 >>>>>>>>>>")
	fmt.Printf("Ratio: %.2f\n", NewGear(52, 11).Ratio())
	fmt.Printf("Ratio: %.2f\n", NewGear(30, 27).Ratio())
}
