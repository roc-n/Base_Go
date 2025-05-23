package dp

import "math"

/* 最小路径和：暴力搜索

   仍存在重复子问题
*/

func minPathSumDFS(grid [][]int, i, j int) int {
	// 若为左上角单元格，则终止搜索
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	// 若行列索引越界，则返回 +∞ 代价
	if i < 0 || j < 0 {
		return math.MaxInt
	}
	// 计算从左上角到 (i-1, j) 和 (i, j-1) 的最小路径代价
	up := minPathSumDFS(grid, i-1, j)
	left := minPathSumDFS(grid, i, j-1)
	// 返回从左上角到 (i, j) 的最小路径代价
	return int(math.Min(float64(left), float64(up))) + grid[i][j]
}

/* 最小路径和：记忆化搜索

   通过二维数组 mem 来存储已经计算过的子问题的解，避免重复子问题
*/
func minPathSumDFSMem(grid, mem [][]int, i, j int) int {
	// 若为左上角单元格，则终止搜索
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	// 若行列索引越界，则返回 +∞ 代价
	if i < 0 || j < 0 {
		return math.MaxInt
	}
	// 若已有记录，则直接返回
	if mem[i][j] != -1 {
		return mem[i][j]
	}
	// 左边和上边单元格的最小路径代价
	up := minPathSumDFSMem(grid, mem, i-1, j)
	left := minPathSumDFSMem(grid, mem, i, j-1)
	// 记录并返回左上角到 (i, j) 的最小路径代价
	mem[i][j] = int(math.Min(float64(left), float64(up))) + grid[i][j]
	return mem[i][j]
}

/* 最小路径和：动态规划

   通过 dp 表来存储已经计算过的子问题的解，避免重复子问题
*/
func minPathSumDP(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	// 初始化 dp 表
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	// 状态转移：首行
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 状态转移：首列
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 状态转移：其余行和列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = int(math.Min(float64(dp[i][j-1]), float64(dp[i-1][j]))) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}
