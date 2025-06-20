package sort

func SelectionSort(arr []int) {
	n := len(arr)
	for i := range n - 1 {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			// Swap the found minimum element with the first element
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
