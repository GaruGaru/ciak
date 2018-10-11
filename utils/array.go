package utils

func StringIn(target string, array []string) bool {
	for _, item := range array {
		if target == item {
			return true
		}
	}
	return false
}
