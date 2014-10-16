package strutil

func SliceIndex(slice []string, str string) int {
	for i, s := range slice {
		if str == s {
			return i
		}
	}
	return -1
}
