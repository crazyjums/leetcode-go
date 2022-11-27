package simple

// BubbleSort
// @arr     待排序的数组
// @reverse 是否倒叙，true-倒序，false-正序
func BubbleSort(arr []int, reverse bool) []int {
	size := len(arr)
	if size <= 0 {
		return arr
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if reverse && arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else if !reverse && arr[j] >= arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func BubbleSortPtr(arrPtr *[]int, size int, reverse bool) {
	if size <= 0 {
		return
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if reverse && (*arrPtr)[j] < (*arrPtr)[j+1] {
				(*arrPtr)[j], (*arrPtr)[j+1] = (*arrPtr)[j+1], (*arrPtr)[j]
			} else if !reverse && (*arrPtr)[j] >= (*arrPtr)[j+1] {
				(*arrPtr)[j], (*arrPtr)[j+1] = (*arrPtr)[j+1], (*arrPtr)[j]
			}
		}
	}
}

func SelectionSort(arr []int, reverse bool) []int {
	size := len(arr)
	if size <= 0 {
		return arr
	}

	var minOrMaxIndex int
	for i := 0; i < size-1; i++ {
		minOrMaxIndex = i
		for j := i + 1; j < size; j++ {
			if reverse && arr[j] >= arr[minOrMaxIndex] {
				minOrMaxIndex = j
			} else if !reverse && arr[j] < arr[minOrMaxIndex] {
				minOrMaxIndex = j
			}
		}
		arr[i], arr[minOrMaxIndex] = arr[minOrMaxIndex], arr[i]
	}

	return arr
}

func SelectionSortPtr(arrPtr *[]int, size int, reverse bool) {
	if size <= 0 {
		return
	}

	var minOrMaxIndex int
	for i := 0; i < size-1; i++ {
		minOrMaxIndex = i
		for j := i + 1; j < size; j++ {
			if reverse && (*arrPtr)[j] >= (*arrPtr)[minOrMaxIndex] {
				minOrMaxIndex = j
			} else if !reverse && (*arrPtr)[j] < (*arrPtr)[minOrMaxIndex] {
				minOrMaxIndex = j
			}
		}
		(*arrPtr)[i], (*arrPtr)[minOrMaxIndex] = (*arrPtr)[minOrMaxIndex], (*arrPtr)[i]
	}
}
