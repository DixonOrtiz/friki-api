package slices

func Exists(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}

	return false
}
