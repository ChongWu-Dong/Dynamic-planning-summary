package 青蛙跳台阶类问题_爬楼梯_

// 机器人走路径
// 机器人只有两种走的可能性，一种是向右一种是向下
// 但是要判断机器人越界，如果机器人越界则表明这次的结果返回0，不是一次合法的路径
// 如果达到目的地则收集这个结果

// 暴力递归
// 递归的意思表示是，从i和j出发一共有多少中可能到达目的地
func uniquePaths(m int, n int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m-1 && j == n-1 {
			return 1
		}
		// 越界的情况都属于没有找到最下方的情况
		if i >= m || j >= n {
			return 0
		}
		// 向下移动
		a := dfs(i+1, j)
		// 向右边移动
		b := dfs(i, j+1)
		return a + b
	}
	return dfs(0, 0)
}

// 动态规划
func uniquePaths2(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
