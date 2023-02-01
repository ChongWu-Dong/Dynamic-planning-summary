package __动态规划二维子序列问题

// 暴力递归
// 讨论三种情况
// 一定不包含i但是可能包含j
// 一定不包含j但是可能包含i
// 一定包含i和j
func longestCommonSubsequence(text1 string, text2 string) int {
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i == 0 && j == 0 {
			if text1[i] == text2[j] {
				return 1
			} else {
				return 0
			}
		} else if i == 0 {
			if text1[i] == text2[j] {
				return 1
			} else {
				return dfs(i, j-1)
			}
		} else if j == 0 {
			if text1[i] == text2[j] {
				return 1
			} else {
				return dfs(i-1, j)
			}
		} else {
			// 一定不含i，但是可能含有j
			a := dfs(i-1, j)
			// 一定不含有j，但是可能含有i
			b := dfs(i, j-1)
			// 一定含有i和j
			c := 0
			if text1[i] == text2[j] {
				c = 1 + dfs(i-1, j-1)
			}
			return max(a, max(b, c))
		}
	}
	return dfs(len(text1)-1, len(text2)-1)
}

// 暴力递归改成动态规划

func longestCommonSubsequence2(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2))
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if i == 0 && j == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
				}
			} else if i == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1]
				}
			} else if j == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				a := dp[i][j-1]
				b := dp[i-1][j]
				c := 0
				if text1[i] == text2[j] {
					c = 1 + dp[i-1][j-1]
				}
				dp[i][j] = max(a, max(b, c))
			}
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
