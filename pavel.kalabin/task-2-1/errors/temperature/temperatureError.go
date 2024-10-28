package temperatureError

import "fmt"

type TemperatureError struct {}

func (e TemperatureError) Error() string {
    return fmt.Sprintf("Wrong temperature: temperature must be between 15 and 30")
}
