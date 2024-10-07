package calculation

import (
	"errors"
	"math"
)

func CalculateNewTemp(minTemp, maxTemp, currentTemp, newTemp *float64, sign string) error {

	switch sign {
	case "<=":
		if *currentTemp > *newTemp || *currentTemp == -1 {
			*currentTemp = *newTemp
		}
		*maxTemp = math.Min(float64(*maxTemp), float64(*newTemp))

	case ">=":
		if *currentTemp < *newTemp || *currentTemp == -1 {
			*currentTemp = *newTemp
		}
		*minTemp = math.Max(float64(*minTemp), float64(*newTemp))

	default:
		return errors.New("invalid character input, only <= or >= allowed")
	}

	return nil
}
