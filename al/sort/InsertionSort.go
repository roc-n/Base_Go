package sort

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		base := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > base {
				arr[j+1] = arr[j] // move larger element one position ahead
			} else {
				break
			}
		}
		arr[j+1] = base // place the base element in its correct position
	}
}
