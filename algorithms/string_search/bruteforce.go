package stringsearch

func SearchBruteForce(text, pattern string) int {
	var found bool
	for i := 0; i+len(pattern)-1 < len(text); i++ {
		if text[i] == pattern[0] {
			found = true
			for j := 1; j < len(pattern); j++ {
				if text[i+j] != pattern[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}
