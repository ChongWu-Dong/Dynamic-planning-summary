package __买卖股票问题

func maxProfitII(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		// 持有股票：保持持有，当天买入
		dp[i][0] = max(dp[i-1][0], -prices[i]+dp[i-1][1])
		// 不持有股票：保持前一天的不持有，当天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	// 返回最后一天的不持有股票的收益
	return dp[len(prices)-1][1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
