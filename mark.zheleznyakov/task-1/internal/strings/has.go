package strings

func Has(arr []string, str string) bool {
	strMap := make(map[string]struct{}, len(arr))
	for _, s := range arr {
		strMap[s] = struct{}{}
	}
	_, found := strMap[str]
	return found
}

