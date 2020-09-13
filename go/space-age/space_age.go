package space

import (
	"fmt"
	"strings"
)

type Planet string

func (planet Planet) normalize() Planet {
	return Planet(strings.ToLower(string(planet)))
}

var orbitalPeriods = map[Planet]float64{
	"mercury": 0.24084670,
	"venus":   0.61519726,
	"earth":   1.00000000,
	"mars":    1.88081580,
	"jupiter": 11.86261500,
	"saturn":  29.44749800,
	"uranus":  84.01684600,
	"neptune": 164.79132000,
}

const secondsInEarthYear = 31557600.0

func Age(ageInSeconds float64, onPlanet Planet) float64 {
	onPlanet = onPlanet.normalize()
	if orbitalPeriod, present := orbitalPeriods[onPlanet]; present {
		ageInEarthYears := ageInSeconds / secondsInEarthYear
		ageInPlanetYears := ageInEarthYears / orbitalPeriod
		return ageInPlanetYears
	} else {
		// We are not expected to return an error, so we panic instead.
		panic(fmt.Sprintf("%s is not a valid planet name", onPlanet))
	}
}
