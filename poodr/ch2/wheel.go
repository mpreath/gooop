package main

type Wheel struct {
	rim, tire float32
}

func NewWheel(rim, tire float32) Wheel {
	return Wheel{rim, tire}
}

func (w Wheel) Diameter() float32 {
	return w.rim + (w.tire * 2)
}

func (w Wheel) Circumference() float32 {
	return w.Diameter() * 3.14159
}
