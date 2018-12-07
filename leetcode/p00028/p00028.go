package p00028

// StrStr ...
func StrStr(haystack string, needle string) int {
	result := -1
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			result = i
			break
		}
	}
	return result
}
