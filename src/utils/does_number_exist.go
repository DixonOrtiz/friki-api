package utils

func DoesNumberExist(number int, numbers []int) bool {
	for _, v := range numbers {
		if number == v {
			return true
		}
	}
	return false
}
