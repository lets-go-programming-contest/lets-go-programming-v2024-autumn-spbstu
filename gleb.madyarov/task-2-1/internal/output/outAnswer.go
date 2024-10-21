package output

func OutAnswer(max int, min int) int {
	if max == 0 {
		return min
	}
	if min == 100 {
		return max
	}
	if min >= max {
		return max
	} else {
		return -1
	}
}
