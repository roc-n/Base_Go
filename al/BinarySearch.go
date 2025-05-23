package al

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* 二分查找（双闭区间） */
func binarySearch(nums []int, target int) int {
	// 初始化双闭区间 [0, n-1] ，即 i, j 分别指向数组首元素、尾元素
	i, j := 0, len(nums)-1
	// 循环，当搜索区间为空时跳出（当 i > j 时为空）
	for i <= j {
		m := i + (j-i)/2      // 计算中点索引 m，避免溢出
		if nums[m] < target { // 此情况说明 target 在区间 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target { // 此情况说明 target 在区间 [i, m-1] 中
			j = m - 1
		} else { // 找到目标元素，返回其索引
			return m
		}
	}
	// 未找到目标元素，返回 -1
	return -1
}

/* 二分查找插入点（无重复元素） */
func BinarySearchInsertionSimple(nums []int, target int) int {
	// 初始化双闭区间 [0, n-1]
	i, j := 0, len(nums)-1
	for i <= j {
		// 计算中点索引 m
		m := i + (j-i)/2
		if nums[m] < target {
			// target 在区间 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target {
			// target 在区间 [i, m-1] 中
			j = m - 1
		} else {
			// 找到 target ，返回插入点 m
			return m
		}
	}
	// 未找到 target ，返回插入点 i
	return i
}

/* 二分查找插入点（存在重复元素） */
func BinarySearchInsertion(nums []int, target int) int {
	// 初始化双闭区间 [0, n-1]
	i, j := 0, len(nums)-1
	for i <= j {
		// 计算中点索引 m
		m := i + (j-i)/2
		if nums[m] < target {
			// target 在区间 [m+1, j] 中
			i = m + 1
		} else if nums[m] > target {
			// target 在区间 [i, m-1] 中
			j = m - 1
		} else {
			// 首个小于 target 的元素在区间 [i, m-1] 中
			j = m - 1
		}
	}
	// 返回插入点 i
	return i
}

func BS() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入整数数组（用空格分隔）：")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	strNums := strings.Split(line, " ")
	nums := make([]int, 0, len(strNums))
	for _, s := range strNums {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("输入有误:", s)
			return
		}
		nums = append(nums, n)
	}

	fmt.Print("请输入要查找的目标值：")
	targetLine, _ := reader.ReadString('\n')
	targetLine = strings.TrimSpace(targetLine)
	target, err := strconv.Atoi(targetLine)
	if err != nil {
		fmt.Println("目标值输入有误")
		return
	}

	result := binarySearch(nums, target)
	if result != -1 {
		fmt.Println("Found target at index:", result)
	} else {
		fmt.Println("Target not found")
	}
}
