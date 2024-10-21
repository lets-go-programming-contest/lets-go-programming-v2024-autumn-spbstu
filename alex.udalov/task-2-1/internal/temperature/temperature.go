package temperature

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

func (t *Tracker) ProcTemperature(op string, value int) int {

	switch op {
	case ">=":
		if value > t.maxTemp {
			return -1
		}
		if value > t.minTemp {
			t.minTemp = value
		}
	case "<=":
		if value < t.minTemp {
			return -1
		}
		if value < t.maxTemp {
			t.maxTemp = value
		}
	}

	if t.minTemp > t.maxTemp {
		return -1
	}

	return t.minTemp
}
