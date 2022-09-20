package ch2

import "math"

type Wheel struct {
	rim  float32
	tire float32
}

func NewWheel(rim float32, tire float32) *Wheel {
	return &Wheel{rim: rim, tire: tire}
}

func (w *Wheel) GetRim() float32 {
	return w.rim
}

func (w *Wheel) SetRim(rim float32) {
	w.rim = rim
}

func (w *Wheel) GetTire() float32 {
	return w.tire
}

func (w *Wheel) SetTire(tire float32) {
	w.tire = tire
}

func (w *Wheel) Diameter() float32 {
	return w.rim + (w.tire * 2)
}

func (w *Wheel) Circumference() float32 {
	return w.Diameter() * math.Pi
}
