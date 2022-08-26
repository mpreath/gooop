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
// embedded values which act a bit like inheritence
type MountainBike struct {
	Bike // we just specify the struct type and no variable name
	// the variables and methods of Bike will be available as variables
	// and methods on MountainBike
}

// we can override the inherited GearRatio() method
// and calculate the value differently
func (m *MountainBike) GearRatio() float64 {
	return float64(m.Chainring.Teeth) / float64(m.Cog.Teeth)
}
