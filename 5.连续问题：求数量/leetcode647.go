package __连续问题_求数量

// 暴力递归
// 回文串的问题，区间法，包含两个变量i和j
// 讨论所有i和j的情况，一定包含i和一定包含j的所有回文串之和
func countSubstrings2(s string) int {
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i == j {
			return 1
		}
		if i+1 == j {
			if s[i] == s[j] {
				return 1
			} else {
				return 0
			}
		}
		if s[i] == s[j] {
			return dfs(i+1, j-1)
		} else {
			return 0
		}
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			ans += dfs(i, j)
		}
	}
	return ans
}

// 动态规划
func countSubstrings3(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
			} else if i+1 == j {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			ans += dp[i][j]
		}
	}
	return ans
}
