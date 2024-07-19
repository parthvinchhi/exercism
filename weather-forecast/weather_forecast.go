// Package weather provides tools to know the weather condition at the given location.
package weather

// CurrentCondition represents certain weather condition.
var CurrentCondition string

// CurrentLocation represents certain location.
var CurrentLocation string

// Forecast returns an interger value describing the weather condition at the given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
