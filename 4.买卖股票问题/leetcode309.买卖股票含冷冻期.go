package __买卖股票问题

func maxProfit5(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	// 初始化
	// 状态1，今天保持买股票的状态
	dp[0][0] = -prices[0]
	// 状态2，卖股票,保持这个状态
	dp[0][1] = 0
	//今天卖股票
	dp[0][2] = 0
	// 冷冻期
	dp[0][3] = 0
	for i := 1; i < len(prices); i++ {
		// 今天买入股票的时候，前一天一定不能卖出状态
		dp[i][0] = Max309(dp[i-1][0], Max309(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i]))
		// 保持前一天的不持有股票的状态
		dp[i][1] = Max309(dp[i-1][3], dp[i-1][1])
		dp[i][2] = prices[i] + dp[i-1][0]
		dp[i][3] = dp[i-1][2]
	}
	return Max309(dp[len(prices)-1][1], Max309(dp[len(prices)-1][2], dp[len(prices)-1][3]))
}

func Max309(a, b int) int {
	if a > b {
		return a
	}
	return b
}
