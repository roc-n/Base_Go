package backtrace

import "slices"

/* 回溯算法：全排列 I */
func backtrackI(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(*state) == len(*choices) {
		newState := slices.Clone(*state)
		*res = append(*res, newState)
	}
	// 遍历所有选择
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		// 剪枝：不允许重复选择元素
		if !(*selected)[i] {
			// 尝试：做出选择，更新状态
			(*selected)[i] = true
			*state = append(*state, choice)
			// 进行下一轮选择
			backtrackI(state, choices, selected, res)
			// 回退：撤销选择，恢复到之前的状态
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

/* 全排列 I */
func PermutationsI(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backtrackI(&state, &nums, &selected, &res)
	return res
}

/* 回溯算法：全排列 II */
func backtrackII(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(*state) == len(*choices) {
		newState := slices.Clone(*state)
		*res = append(*res, newState)
	}
	// 遍历所有选择
	duplicated := make(map[int]struct{}, 0)
	for i := range *choices {
		choice := (*choices)[i]
		// 剪枝：不允许重复选择元素 且 不允许重复选择相等元素
		if _, ok := duplicated[choice]; !ok && !(*selected)[i] {
			// 尝试：做出选择，更新状态
			// 记录选择过的元素值
			duplicated[choice] = struct{}{}
			(*selected)[i] = true
			*state = append(*state, choice)
			// 进行下一轮选择
			backtrackII(state, choices, selected, res)
			// 回退：撤销选择，恢复到之前的状态
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

/* 全排列 II */
func PermutationsII(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backtrackII(&state, &nums, &selected, &res)
	return res
}
