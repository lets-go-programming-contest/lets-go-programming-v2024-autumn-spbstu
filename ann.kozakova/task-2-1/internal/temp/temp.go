package temp

import (
	"errors"
)

func FindTemp(topT, bottomT, temperature *int, sign string) (int, error) {
	if sign == "<=" && *temperature < *topT {
		*topT = *temperature
	}
	if sign == ">=" && *temperature > *bottomT {
		*bottomT = *temperature
	}
	if *topT >= *bottomT {
		return *bottomT, nil
	} else {
		return 0, errors.New("impossible")
	}
}
