package slices

func RemoveElement(slice []string, element string) []string {
	for i, value := range slice {
		if value == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
