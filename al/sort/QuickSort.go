package sort

func partition(arr []int, left, right int) int {
	// 默认指定左边界为基准点
	pivot := left
	i, j := left, right

	for i < j {
		// 这里一定要先减小j，否则会有边界问题
		for i < j && arr[j] >= arr[pivot] {
			j--
		}
		for i < j && arr[i] <= arr[pivot] {
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	// 基数交换至数组分界线
	arr[pivot], arr[i] = arr[i], arr[pivot]
	return i
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)

		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
}
