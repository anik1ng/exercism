// Package weather provides tools to report the current weather condition.
package weather

var (
    // CurrentCondition stores the current weather condition.
	CurrentCondition string
    // CurrentLocation stores the name of the city for the current weather report.
	CurrentLocation  string
)

// Forecast sets the given city and weather condition, and returns a formatted weather report.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
