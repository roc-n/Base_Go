package greedy

/* 零钱兑换：贪心 */
// 题目：给定不同面额的硬币 coins 和一个总金额 amt，计算最少需要多少个硬币才能凑出该金额
// 不一定能找到最优解，但在某些情况下可以快速找到一个可行解。
func CoinChangeGreedy(coins []int, amt int) int {
	// 假设 coins 列表有序
	i := len(coins) - 1
	count := 0
	// 循环进行贪心选择，直到无剩余金额
	for amt > 0 {
		// 找到小于且最接近剩余金额的硬币
		for i > 0 && coins[i] > amt {
			i--
		}
		// 选择 coins[i]
		amt -= coins[i]
		count++
	}
	// 若未找到可行方案，则返回 -1
	if amt != 0 {
		return -1
	}
	return count
}
