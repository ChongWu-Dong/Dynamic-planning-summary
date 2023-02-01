package dp2

// 对于二维连续的问题一定要讨论包含i和j的情况
// 暴力递归
func findLength(nums1 []int, nums2 []int) int {
	// 讨论包含i和j的情况
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		// 分类讨论
		if i == 0 && j == 0 {
			if nums1[i] == nums2[j] {
				return 1
			}
			return 0
		} else if i == 0 {
			if nums1[i] == nums2[j] {
				return 1
			}
		} else if j == 0 {
			if nums1[i] == nums2[j] {
				return 1
			}
		} else { // i和j都不为0
			a := 0
			if nums1[i] == nums2[j] {
				a = 1 + dfs(i-1, j-1)
			}
			return a
		}
		return 0
	}
	ans := 0
	// 枚举i和j的所有情况
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			a := dfs(i, j)
			if a > ans {
				ans = a
			}
		}
	}
	return ans
}

// 动态规划
// 根据暴力递归改写动态规划
func findLength2(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
	}
	ans := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if i == 0 && j == 0 {
				if nums1[i] == nums2[j] {
					dp[i][j] = 1
				}
			} else if i == 0 {
				if nums1[i] == nums2[j] {
					dp[i][j] = 1
				}
			} else if j == 0 {
				if nums1[i] == nums2[j] {
					dp[i][j] = 1
				}
			} else {
				if nums1[i] == nums2[j] {
					dp[i][j] = 1 + dp[i-1][j-1]
				}
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
