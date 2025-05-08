package ds

import (
	"fmt"
	"math/rand"
)

func Array() {
	var arr [5]int
	nums := [5]int{1, 2, 3, 4, 5}
	nums_ := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(nums)

	fmt.Println(nums_)
	fmt.Println(len(nums_))
	fmt.Println(cap(nums_))

	nums_ = append(nums_, 6)
	fmt.Println(len(nums_))
	fmt.Println(cap(nums_))

	// insert
	array := []int{1, 2, 3, 4, 5}
	insert(array, 6, 2)
	fmt.Println(array)
	remove(array, 2)
	fmt.Println(array)
	array = extend(array, 5)
	fmt.Println(array)
}

func randomAccess(nums []int) (random int) {
	randomIndex := rand.Intn(len(nums))
	randomNumber := nums[randomIndex]
	return randomNumber
}

/* 在数组的索引 index 处插入元素 num */
func insert(nums []int, num int, index int) {
	// 把索引 index 以及之后的所有元素向后移动一位
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	// 将 num 赋给 index 处的元素
	nums[index] = num
}

/* 删除索引 index 处的元素 */
func remove(nums []int, index int) {
	// 把索引 index 之后的所有元素向前移动一位，末尾元素无意义
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

/* 扩展数组长度 */
func extend(nums []int, enlarge int) []int {
	// 初始化一个扩展长度后的数组
	res := make([]int, len(nums)+enlarge)
	// 将原数组中的所有元素复制到新数组
	for i, num := range nums {
		res[i] = num
	}
	// 返回扩展后的新数组
	return res
}
