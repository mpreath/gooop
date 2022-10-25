package listing2_4

import (
	"fmt"
	"gooop/poodr/ch2/listing2_2"
)

func Run() {
	fmt.Println("Listing 2.4 >>>>>>>>>>")
	// ok
	_ = listing2_2.NewGear(52, 11)
	// not ok
	// _ = listing2_3.NewGear(52, 11)
}
