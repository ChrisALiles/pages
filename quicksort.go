package pages

func quicksort(arr []int64, low int, high int) {
	if low < high {
		piv := partition(arr, low, high)
		quicksort(arr, low, piv-1)
		quicksort(arr, piv+1, high)
	}
}

func partition(arr []int64, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
