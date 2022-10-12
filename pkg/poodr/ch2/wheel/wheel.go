package wheel

import "math"

type Wheel struct {
	rim  int
	tire float32
}

func New(rim int, tire float32) *Wheel {
	return &Wheel{
		rim:  rim,
		tire: tire,
	}
}

func (w *Wheel) Rim() int {
	return w.rim
}

func (w *Wheel) SetRim(rim int) {
	w.rim = rim
}

func (w *Wheel) Tire() float32 {
	return w.tire
}

func (w *Wheel) SetTire(tire float32) {
	w.tire = tire
}

func (w *Wheel) Circumference() float32 {
	return w.Diameter() * math.Pi
}

func (w *Wheel) Diameter() float32 {
	return float32(w.rim) + (w.tire * 2)
}
