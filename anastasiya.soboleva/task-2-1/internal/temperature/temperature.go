package temperature

import (
	"errors"
)

type Tracker struct {
	minTemp int
	maxTemp int
}

func NewTracker() *Tracker {
	return &Tracker{
		minTemp: 15,
		maxTemp: 30,
	}
}

func (t *Tracker) ProcessTemperature(op string, value int) (int, error) {
	if value > 30 || value < 15 {
		return -1, errors.New("некорректная температура")
	}

	switch op {
	case ">=":
		if value > t.maxTemp {
			return -1, nil
		}
		if value > t.minTemp {
			t.minTemp = value
		}
	case "<=":
		if value < t.minTemp {
			return -1, nil
		}
		if value < t.maxTemp {
			t.maxTemp = value
		}
	default:
		return -1, errors.New("некорректная операция")
	}

	if t.minTemp > t.maxTemp {
		return -1, nil
	}

	return t.minTemp, nil
}
