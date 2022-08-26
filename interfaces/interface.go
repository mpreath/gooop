package interfaces

// we learned about composition previously
// in order to have maximum flexibility when
// composing objects we should use interfaces

// interfaces define the methods expected on an
// object
type Bike interface {
	SetFrontWheel(Wheel)
	GetFrontWheel() Wheel
	SetRearWheel(Wheel)
	GetRearWheel() Wheel
}

type Wheel interface {
	SetSize(int)
	GetSize() int
}

// in our base bike struct we define our wheels
// using the wheel interface
type BaseBike struct {
	frontWheel Wheel
	rearWheel  Wheel
}

type RoadBike struct {
	BaseBike
}

type MountainBike struct {
	BaseBike
}

// the base bike methods use the Wheel interface
// for method arguments and return values
// this allows us to set the wheels to either type of wheel
// road or offroad
func (b *BaseBike) SetFrontWheel(wheel Wheel) {
	b.frontWheel = wheel
}

func (b *BaseBike) GetFrontWheel() Wheel {
	return b.frontWheel
}

func (b *BaseBike) SetRearWheel(wheel Wheel) {
	b.rearWheel = wheel
}

func (b *BaseBike) GetRearWheel() Wheel {
	return b.rearWheel
}

// we create a basewheel class that will serve
// as the base class for other wheel types

type BaseWheel struct {
	size int
}

func (w *BaseWheel) SetSize(size int) {
	w.size = size
}

func (w *BaseWheel) GetSize() int {
	return w.size
}

// we create both road and Offroad wheels
// that "inherit" from the base class both variables and methods
// since our Wheel interface requires SetSize() and GetSize()
// by inheriting from the base class the road and offroad wheels
// comply with the interface
type RoadWheel struct {
	BaseWheel
}

type OffroadWheel struct {
	BaseWheel
}
