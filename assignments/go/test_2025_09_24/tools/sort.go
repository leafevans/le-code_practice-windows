package tools

func QuickSort[T int | float32 | float64](slice []T) {
	if len(slice) <= 1 {
		return
	}
	pivot := slice[len(slice)/2]
	left, right := 0, len(slice)-1

	for left <= right {
		for slice[left] < pivot {
			left++
		}
		for slice[right] > pivot {
			right--
		}
		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}
	}

	QuickSort(slice[left:])
	QuickSort(slice[:right+1])
}
