package __买卖股票问题

// 这里是直接限定了只能进行两次买卖
// 在第一次买入的时候，因为再第一次买入的时候没有卖出操作，所以只能是-price[i] 而不能是dp[i-1][1]-price[i]

// 动态规划，写不出暴力递归
func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 5)
	}
	// 初始化
	// 不操作
	dp[0][0] = 0
	// 第一次买
	dp[0][1] = -prices[0]
	// 第一次卖
	dp[0][2] = 0
	// 第二次买
	dp[0][3] = -prices[0]
	// 第二次卖
	dp[0][4] = 0
	for i := 1; i < len(prices); i++ {
		// 无操作
		dp[i][0] = dp[i-1][0]
		// 第一次买
		dp[i][1] = Max123(dp[i-1][1], -prices[i])
		// 第一次卖
		dp[i][2] = Max123(prices[i]+dp[i-1][1], dp[i-1][2])
		// 第二次买,收益为，选取最大的收益
		dp[i][3] = Max123(dp[i-1][3], dp[i-1][2]-prices[i])
		// 第二次卖
		dp[i][4] = Max123(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}
func Max123(a, b int) int {
	if a > b {
		return a
	}
	return b
}
