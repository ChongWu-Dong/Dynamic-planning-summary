package 青蛙跳台阶类问题_爬楼梯_

// 这个和62区别在于这个增加了一些判断条件，即增加了障碍物，遇到障碍物，表明这一次的结果不会收集
// 遇到障碍物则返回0
// 递归的意思表示是，从i和j出发一共有多少中可能到达目的地
// 暴力递归
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m-1 && j == n-1 && obstacleGrid[m-1][n-1] != 1 {
			return 1
		}
		if i >= m || j >= n || obstacleGrid[i][j] == 1 { // 越界或者遇到障碍物
			return 0
		}
		a := dfs(i+1, j)
		b := dfs(i, j+1)
		return a + b
	}
	return dfs(0, 0)
}

// 动态规划
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 && obstacleGrid[m-1][n-1] != 1 {
				dp[m-1][n-1] = 1
			} else if obstacleGrid[i][j] == 1 { // 遇到障碍物
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
