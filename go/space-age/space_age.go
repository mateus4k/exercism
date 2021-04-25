// Package space provides mehods around space math
package space

// Planet represents the name of a planet in the solar system
type Planet string

// EarthYearInSeconds defines the number of seconds contained in a Earth year
const EarthYearInSeconds float64 = 31557600

var planetMap = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the age in Earth years of someone given in another planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / (planetMap[planet] * EarthYearInSeconds)
}
