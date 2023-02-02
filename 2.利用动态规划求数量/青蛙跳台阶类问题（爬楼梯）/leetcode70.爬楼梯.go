package 青蛙跳台阶类问题_爬楼梯_

// 跳台阶的问题，每次跳台阶只有两种选择，选择一步或者两部
// 当台阶数位0的时候说明这是一种跳跃的结果
// 递归
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	a := climbStairs(n - 1)
	b := 0
	if n-2 >= 0 {
		b = climbStairs(n - 2)
	}

	return a + b
}

// 改写动态规划
// 还剩下n个台阶又多少种方法，剩下0个台阶的时候，说明这时一种结果，可以收集这一种结果
// 所以在n==0的时候要返回1
func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		if n >= 2 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
