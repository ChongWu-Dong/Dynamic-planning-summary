package 一维子序列问题但是要讨论包含边界

// 子序列的问题
func lengthOfLIS(nums []int) int {
	// dp
	dp := make([]int, len(nums))
	// 初始化dp
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		// 对于每一个i都要遍历其前面的元素，找到该包含该元素的最长递增子序列
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], 1+dp[j])
			}
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
