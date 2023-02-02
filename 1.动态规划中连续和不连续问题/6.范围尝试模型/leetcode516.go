package __范围尝试模型

// 将字符串反转形成s2
// 样本对应模型
// 然后本题就转换为求s和s2的最长公共子序列问题，最长公共子序列就是最长的回文子序列。
func longestPalindromeSubseq(s string) int {
	s2 := reverse(s)
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == 0 && j == 0 {
				if s[i] == s2[j] {
					dp[i][j] = 1
				}
			} else if i == 0 {
				if s[i] == s2[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1]
				}
			} else if j == 0 {
				if s[i] == s2[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				a := dp[i-1][j]
				b := dp[i][j-1]
				c := 0
				if s[i] == s2[j] {
					c = 1 + dp[i-1][j-1]
				}
				dp[i][j] = max(a, max(b, c))
			}
		}
	}
	return dp[len(s)-1][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func reverse(s string) string {
	ss := []byte(s)
	l := 0
	r := len(s) - 1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}
	return string(ss)
}

// 范围尝试模型
// str[L...R]的最长回文子序列是多少，L和R为边界
// 可能性：
// 1.一定不包含i但是可能包含j
// 2.一定不包含j但是可能包含i
// 3.一定包含i和j
func longestPalindromeSubseq2(s string) int {
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i == j {
			return 1
		}
		if i+1 == j {
			if s[i] == s[j] {
				return 2
			}
			return 1
		}
		a := dfs(i+1, j)
		b := dfs(i, j-1)
		c := 0
		if s[i] == s[j] {
			c = 2 + dfs(i+1, j-1)
		}
		return max(a, max(b, c))
	}
	return dfs(0, len(s)-1)
}

// 将范围尝试模型改成动态规划
func longestPalindromeSubseq3(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
			} else if i+1 == j {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else {
				a := dp[i+1][j]
				b := dp[i][j-1]
				c := 0
				if s[i] == s[j] {
					c = 2 + dp[i+1][j-1]
				}
				dp[i][j] = max(a, max(b, c))
			}
		}
	}
	return dp[0][len(s)-1]
}
