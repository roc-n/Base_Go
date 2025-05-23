package dp

/* 回溯 */
func backtrack(choices []int, state, n int, res []int) {
	// 当爬到第 n 阶时，方案数量加 1
	if state == n {
		res[0] = res[0] + 1
	}
	// 遍历所有选择
	for _, choice := range choices {
		// 剪枝：不允许越过第 n 阶
		if state+choice > n {
			continue
		}
		// 尝试：做出选择，更新状态
		backtrack(choices, state+choice, n, res)
		// 回退
	}
}

/* 爬楼梯：回溯 */
func ClimbingStairsBacktrack(n int) int {
	// 可选择向上爬 1 阶或 2 阶
	choices := []int{1, 2}
	// 从第 0 阶开始爬
	state := 0
	res := make([]int, 1)
	// 使用 res[0] 记录方案数量
	res[0] = 0
	backtrack(choices, state, n, res)
	return res[0]
}

/* 搜索 */
func dfs(i int) int {
	// 已知 dp[1] 和 dp[2] ，返回之
	if i == 1 || i == 2 {
		return i
	}
	// dp[i] = dp[i-1] + dp[i-2]
	count := dfs(i-1) + dfs(i-2)
	return count
}

/* 爬楼梯：搜索 */
func ClimbingStairsDFS(n int) int {
	return dfs(n)
}

/* 爬楼梯：动态规划 */
func climbingStairsDP(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	// 初始化 dp 表，用于存储子问题的解
	dp := make([]int, n+1)
	// 初始状态：预设最小子问题的解
	dp[1] = 1
	dp[2] = 2
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
