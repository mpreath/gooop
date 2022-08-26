package factory

import (
	"gooop/interfaces"
)

// the power of interfaces can be seen
// when using factories
func BikeFactory(bikeType string) interfaces.Bike {
	switch bikeType {
	case "mountain":
		bike := &interfaces.MountainBike{}
		bike.SetFrontWheel(WheelFactory("offroad", 29))
		bike.SetRearWheel(WheelFactory("offroad", 29))
		return bike
	default:
		bike := &interfaces.RoadBike{}
		bike.SetFrontWheel(WheelFactory("road", 26))
		bike.SetRearWheel(WheelFactory("road", 26))
		return bike
	}
}

func WheelFactory(wheelType string, size int) interfaces.Wheel {
	switch wheelType {
	case "offroad":
		wheel := &interfaces.OffroadWheel{}
		wheel.SetSize(size)
		return wheel
	default:
		wheel := &interfaces.RoadWheel{}
		wheel.SetSize(size)
		return wheel
	}
}
