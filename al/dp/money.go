package dp

import "math"

/* 零钱兑换：动态规划 */
func coinChangeDP(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	// 初始化 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amt+1)
	}
	// 状态转移：首行首列
	for a := 1; a <= amt; a++ {
		dp[0][a] = max
	}
	// 状态转移：其余行和列
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				// 若超过目标金额，则不选硬币 i
				dp[i][a] = dp[i-1][a]
			} else {
				// 不选和选硬币 i 这两种方案的较小值
				dp[i][a] = int(math.Min(float64(dp[i-1][a]), float64(dp[i][a-coins[i-1]]+1)))
			}
		}
	}
	if dp[n][amt] != max {
		return dp[n][amt]
	}
	return -1
}

/* 零钱兑换：动态规划 */
func coinChangeDPComp(coins []int, amt int) int {
	n := len(coins)
	max := amt + 1
	// 初始化 dp 表
	dp := make([]int, amt+1)
	for i := 1; i <= amt; i++ {
		dp[i] = max
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		// 正序遍历
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				// 若超过目标金额，则不选硬币 i
				// dp[a] = dp[a]
			} else {
				// 不选和选硬币 i 这两种方案的较小值
				dp[a] = int(math.Min(float64(dp[a]), float64(dp[a-coins[i-1]]+1)))
			}
		}
	}
	if dp[amt] != max {
		return dp[amt]
	}
	return -1
}

/* 零钱兑换 II：动态规划 */
func coinChangeIIDP(coins []int, amt int) int {
	n := len(coins)
	// 初始化 dp 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amt+1)
	}
	// 初始化首列
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	// 状态转移：其余行和列
	for i := 1; i <= n; i++ {
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				// 若超过目标金额，则不选硬币 i
				dp[i][a] = dp[i-1][a]
			} else {
				// 不选和选硬币 i 这两种方案之和
				dp[i][a] = dp[i-1][a] + dp[i][a-coins[i-1]]
			}
		}
	}
	return dp[n][amt]
}

/* 零钱兑换 II：空间优化后的动态规划 */
func coinChangeIIDPComp(coins []int, amt int) int {
	n := len(coins)
	// 初始化 dp 表
	dp := make([]int, amt+1)
	dp[0] = 1
	// 状态转移
	for i := 1; i <= n; i++ {
		// 正序遍历
		for a := 1; a <= amt; a++ {
			if coins[i-1] > a {
				// 若超过目标金额，则不选硬币 i
				dp[a] = dp[a]
			} else {
				// 不选和选硬币 i 这两种方案之和
				dp[a] = dp[a] + dp[a-coins[i-1]]
			}
		}
	}
	return dp[amt]
}
