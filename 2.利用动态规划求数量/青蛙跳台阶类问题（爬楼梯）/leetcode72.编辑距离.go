package 青蛙跳台阶类问题_爬楼梯_

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	} else if len(word1) == 0 { // 只能进行插入操作
		return len(word2)
	} else if len(word2) == 0 { // 只能进行删除操作
		return len(word1)
	}
	var travel func(i, j int) int
	travel = func(i, j int) int {
		if i == 0 && j == 0 {
			if word1[i] == word2[j] {
				return 0
			} else { // 替换
				return 1
			}
		} else if i == 0 {
			if word1[i] == word2[j] { // 两个相等只需要插入j个元素即可
				return j
			} else { // 两个元素不相等，则插入元素
				return 1 + travel(i, j-1)
			}
		} else if j == 0 {
			if word1[i] == word2[j] { // 删除i个元素
				return i
			} else { // 删除操作，删除当前不相等的元素
				return 1 + travel(i-1, j)
			}
		} else {
			if word1[i] == word2[j] {
				return travel(i-1, j-1)
			} else {
				// 删除操作
				a := 1 + travel(i-1, j)
				// 替换操作
				b := 1 + travel(i-1, j-1)
				// 插入操作
				c := 1 + travel(i, j-1)
				return Min(a, Min(b, c))
			}
		}
	}
	return travel(len(word1)-1, len(word2)-1)
}

// 动态规划
func minDistance2(word1 string, word2 string) int {
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	} else if len(word1) == 0 { // 只能进行插入操作
		return len(word2)
	} else if len(word2) == 0 { // 只能进行删除操作
		return len(word1)
	}

	N := len(word1)
	M := len(word2)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}
	if word1[0] == word2[0] {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}

	for i := 1; i < N; i++ {
		if word1[i] == word2[0] {
			dp[i][0] = i
		} else {
			dp[i][0] = 1 + dp[i-1][0]
		}
	}

	for j := 1; j < M; j++ {
		if word1[0] == word2[j] {
			dp[0][j] = j
		} else {
			dp[0][j] = 1 + dp[0][j-1]
		}
	}

	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(1+dp[i-1][j-1], Min(1+dp[i-1][j], 1+dp[i][j-1]))
			}
		}
	}
	return dp[N-1][M-1]

}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
