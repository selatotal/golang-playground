// Package weather Show forecast weather contitions.
package weather


var (
    // CurrentCondition weather.
	CurrentCondition string
    // CurrentLocation information.
	CurrentLocation  string
)

// Forecast Show current weather condition at current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
