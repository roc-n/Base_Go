package al

import (
	"fmt"

	"algorithm.labmem.net/al/backtrace"
	"algorithm.labmem.net/al/greedy"
	"algorithm.labmem.net/al/sort"
)

func PermutationsIDemo() {
	nums := []int{1, 2, 3}
	fmt.Println("全排列 I：", nums)
	res := backtrace.PermutationsI(nums)
	fmt.Println("全排列 I：", res)
}

func SortDemo() {
	// 这里可以调用不同的排序算法进行演示
	arr := []int{3, 1, 2, 5, 4, 10}
	// sort.SelectionSort(arr)
	// fmt.Println("选择排序结果：", arr)
	// sort.BubbleSort(arr)
	// fmt.Println("冒泡排序结果：", arr)
	// sort.InsertionSort(arr)
	// fmt.Println("插入排序结果：", arr)
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("快速排序结果：", arr)
}

func GreedyDemo() {
	// 零钱兑换问题
	coins := []int{1, 5, 10, 25}
	amt := 63
	result := greedy.CoinChangeGreedy(coins, amt)
	fmt.Println(result)
}
