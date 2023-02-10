package __买卖股票问题

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	// 初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		// 状态一，持有股票：保持持有，当天买入
		dp[i][0] = Max(dp[i-1][0], -prices[i])
		// 状态二：不持有股票，保持不持有，当天卖出
		dp[i][1] = Max(dp[i-1][1], prices[i]+dp[i-1][0])
	}
	// 返回最后一天不持有股票的结果
	return dp[len(prices)-1][1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
