// Package weather exposes a Forecast method that provides the weather details of the provided city with condition. 
package weather
var (
// CurrentCondition comment.
	CurrentCondition string
// CurrentLocation comment.
	CurrentLocation  string
)
// Forecast Comment.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
