package dp1

// 一维连续问题，讨论包含结尾的情况
func findLengthOfLCIS(nums []int) int {
	// 创建动态规划一维数组
	dp := make([]int, len(nums))
	ans := 1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(dp); i++ {
		// 如果当前的值比前一个值大，则表面递增序列可以通过前一个元素的递增序列可得
		if nums[i] > nums[i-1] {
			// 状态转移
			dp[i] = 1 + dp[i-1]
		}
		ans = max(dp[i], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
