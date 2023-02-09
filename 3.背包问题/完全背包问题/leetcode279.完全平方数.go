package 完全背包问题

// 对于完全平方数的问题，可以多次取不大于n以下的数，只要是满足一定的条件即可

// 暴力递归
func numSquares(n int) int {

	var dfs func(n int) int // 递归的含义是，整数n的完全平方最少数量
	dfs = func(n int) int {
		if n == 0 {
			return 0
		}
		ans := n // 最大就是n，全部都是1
		for i := 1; i*i <= n; i++ {
			ans = min(ans, 1+dfs(n-i*i))
		}
		return ans
	}
	return dfs(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 改写动态规划
func numSquares2(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}
	return dp[n]
}
