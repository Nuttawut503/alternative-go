package sorting

func heapify(arr []int, start int) {
	i := start
	for i*2+1 < len(arr) {
		if i*2+2 < len(arr) {
			if arr[i*2+1] < arr[i*2+2] {
				if arr[i] < arr[i*2+2] {
					arr[i], arr[i*2+2] = arr[i*2+2], arr[i]
					i = i*2 + 2
				} else {
					break
				}
			} else if arr[i] < arr[i*2+1] {
				arr[i], arr[i*2+1] = arr[i*2+1], arr[i]
				i = i*2 + 1
			} else {
				break
			}
		} else if arr[i] < arr[i*2+1] {
			arr[i], arr[i*2+1] = arr[i*2+1], arr[i]
			i = i*2 + 1
		} else {
			break
		}
	}
}

func heapsort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr[:i], 0)
	}
}

func Heapsort(arr []int) []int {
	modified := make([]int, len(arr))
	copy(modified, arr)
	heapsort(modified)
	return modified
}
