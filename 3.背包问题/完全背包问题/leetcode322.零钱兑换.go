package 完全背包问题

import "math"

// 完全背包问题，对于一个amout可以从coins中选择任意数量的硬币知道满足条件为止。

// 暴力递归
func coinChange(coins []int, amount int) int {
	var dfs func(amount int) int // 递归的含义是当前的数额是amount，最少需要的硬币数
	dfs = func(amount int) int {
		if amount == 0 { // 当前已经不用找零钱了
			return 0
		}
		ans := math.MaxInt64 // 这个表示最大的硬币数量
		for i := 0; i < len(coins); i++ {
			if amount-coins[i] >= 0 { // 先判断当前的的amount够不够抵扣零钱
				if dfs(amount-coins[i]) != math.MaxInt64 { // 只有返回的dfs不等于系统最大值才进行，因为系统最大值+1之后会报错
					ans = min322(ans, 1+dfs(amount-coins[i]))
				}
			}
		}
		return ans
	}
	if dfs(amount) == math.MaxInt64 {
		return -1
	}
	return dfs(amount)
}

// 改写动态规划
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		ans := math.MaxInt64
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				if dp[i-coins[j]] != math.MaxInt64 {
					ans = min322(ans, 1+dp[i-coins[j]])
				}
			}
		}
		dp[i] = ans
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min322(a, b int) int {
	if a < b {
		return a
	}
	return b
}
