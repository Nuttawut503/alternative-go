package sorting

func merge(arr []int, mid int) {
	temp, l, r, i := make([]int, len(arr)), 0, mid, 0
	for l < mid && r < len(arr) {
		if arr[l] < arr[r] {
			temp[i] = arr[l]
			l++
		} else {
			temp[i] = arr[r]
			r++
		}
		i++
	}
	for l < mid {
		temp[i] = arr[l]
		i++
		l++
	}
	for r < len(arr) {
		temp[i] = arr[r]
		i++
		r++
	}
	copy(arr, temp)
}

func mergesort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		mergesort(arr[:mid])
		mergesort(arr[mid:])
		merge(arr, mid)
	}
}

func Mergesort(arr []int) []int {
	modified := make([]int, len(arr))
	copy(modified, arr)
	mergesort(modified)
	return modified
}
