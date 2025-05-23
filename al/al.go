package al

import (
	"fmt"

	"algorithm.labmem.net/al/backtrace"
)

func PermutationsIDemo() {
	nums := []int{1, 2, 3}
	fmt.Println("全排列 I：", nums)
	res := backtrace.PermutationsI(nums)
	fmt.Println("全排列 I：", res)
}
