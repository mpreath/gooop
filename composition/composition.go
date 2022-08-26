package composition

// in Go we can compose more complicated objects using
// structs

// Each Bike object will have two wheels and two gear sprockets
type Bike struct {
	FrontWheel Wheel
	RearWheel  Wheel
	Chainring  Gear
	Cog        Gear
}

func (b *Bike) GearRatio() float64 {
	return float64(b.Chainring.Teeth) / float64(b.Cog.Teeth)
}

// Each wheel has a size
type Wheel struct {
	Size int
}
type Gear struct {
	Teeth int
}

// in Go we can also compose objects using
// embedded values
