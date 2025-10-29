package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fast I/O Helper
var reader = bufio.NewReader(os.Stdin)

func readString() string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

func readInt() int {
	str := readString()
	val, _ := strconv.Atoi(str)
	return val
}

func readInts() []int {
	str := readString()
	parts := strings.Split(str, " ")
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p)
	}
	return nums
}

func main() {
	// --- 读取第一行: n, k, a, b ---
	// 注意：这里假设 n, k, a, b 在同一行
	// 如果它们在不同行，需要多次调用 readInt()
	params := readInts()
	n, k, a, b := params[0], params[1], params[2], params[3]

	// --- 读取第二行: 能力值 ---
	ab := readInts()

	// --- 读取第三行: 合作值 ---
	co := readInts()

	// 打印出来验证一下
	fmt.Println("n, k, a, b:", n, k, a, b)
	fmt.Println("能力值:", ab)
	fmt.Println("合作值:", co)

	// 在这里继续你的算法逻辑...
	// 比如，将你的 Python 算法逻辑翻译过来
	cnt := 0
	for i := range k - 1 {
		if ab[i] >= a && co[i] >= b {
			cnt++
		}
	}

	result := 0
	for i := k - 1; i < n; i++ {
		if ab[i] >= a && co[i] >= b {
			cnt++
		}
		if i >= k {
			if ab[i-k] >= a && co[i-k] >= b {
				cnt--
			}
		}
		if cnt == k {
			result++
		}
	}
	fmt.Println("最终结果:", result)
}
