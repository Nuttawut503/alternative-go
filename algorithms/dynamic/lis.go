package dynamic

func getOnes(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = 1
	}
	return arr
}

func LongestIncreasingSubsequence(arr []int) int {
	lis := getOnes(len(arr))
	longestSize := 1
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] >= arr[j] || lis[i]+1 <= lis[j] {
				continue
			}
			lis[j] = lis[i] + 1
			if longestSize < lis[j] {
				longestSize = lis[j]
			}
		}
	}
	return longestSize
}
