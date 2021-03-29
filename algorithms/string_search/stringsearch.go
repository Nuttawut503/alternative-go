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

func buildKMPIndices(pattern string) []int {
	kmpIndices := make([]int, len(pattern))
	kmpIndices[0] = 0
	i, j := 1, 0
	for i < len(pattern) {
		if pattern[j] == pattern[i] {
			kmpIndices[i] = j + 1
			j++
			i++
		} else if j != 0 {
			j = kmpIndices[j-1]
		} else {
			kmpIndices[i] = 0
			i++
		}
	}
	return kmpIndices
}

func SearchKMP(text, pattern string) int {
	kmpIndices := buildKMPIndices(pattern)
	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		} else if j != 0 {
			j = kmpIndices[j-1]
		} else {
			i++
		}
	}
	if j == len(pattern) {
		return i - j
	}
	return -1
}
