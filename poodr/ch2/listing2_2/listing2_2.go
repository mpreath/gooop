package listing2_2

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

func (g *Gear) Ratio() float32 {
	return float32(g.chainring) / float32(g.cog)
}

func Run() {

	fmt.Println("List 2.2 >>>>>>>>>>")
	fmt.Printf("Ratio: %.2f\n", NewGear(52, 11).Ratio())
	fmt.Printf("Ratio: %.2f\n", NewGear(30, 27).Ratio())
}
