package slices

func Exists(slice []string, valueToFind string) bool {
	for _, value := range slice {
		if value == valueToFind {
			return true
		}
	}

	return false
}
