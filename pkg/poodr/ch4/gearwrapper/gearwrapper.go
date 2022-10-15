package gearwrapper

import (
	"gooop/pkg/poodr/ch4/gear"
	"gooop/pkg/poodr/ch4/wheel"
)

func Gear(chainring int, cog int, wheel *wheel.Wheel) *gear.Gear {
	return gear.New(chainring, cog, wheel)
}
