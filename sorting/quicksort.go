package sorting

func partition(arr []int) {
	if len(arr) < 3 {
		if len(arr) == 2 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return
	}
	pivotIndex := len(arr) / 2 // or random between [0, len(arr))
	arr[0], arr[pivotIndex] = arr[pivotIndex], arr[0]
	pivotIndex = 0

	lessItemSize := 0 // compare to the pivot
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[pivotIndex] {
			lessItemSize++
			arr[i], arr[lessItemSize] = arr[lessItemSize], arr[i]
		}
	}

	arr[pivotIndex], arr[lessItemSize] = arr[lessItemSize], arr[pivotIndex]
	partition(arr[:lessItemSize])
	partition(arr[lessItemSize+1:])
}

func Quicksort(arr []int) []int {
	modified := make([]int, len(arr))
	copy(modified, arr)
	partition(modified)
	return modified
}
