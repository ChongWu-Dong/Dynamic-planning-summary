package __买卖股票问题

//
// 在第一次买入的时候，因为再第一次买入的时候没有卖出操作，所以只能是-price[i] 而不能是dp[i-1][1]-price[i]
func maxProfit4(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k)
	}
	for i := 0; i < 2*k; i = i + 2 {
		dp[0][i] = -prices[0]
		dp[0][i+1] = 0
	}
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			// 第j次买入,j==0表示第一次买入，j+1表示第一次卖出
			if j == 0 {
				dp[i][j] = Max188(dp[i-1][j], -prices[i])
				// 上次一的卖出，和现在卖出收益对比
				dp[i][j+1] = Max188(dp[i-1][j+1], dp[i-1][j]+prices[i])
			} else { // 如果不是第一次买入卖出
				dp[i][j] = Max188(dp[i-1][j], dp[i-1][j-1]-prices[i])
				dp[i][j+1] = Max188(dp[i-1][j+1], dp[i-1][j]+prices[i])
			}
		}
	}
	return dp[len(prices)-1][2*k-1]
}
func Max188(a, b int) int {
	if a > b {
		return a
	}
	return b
}
