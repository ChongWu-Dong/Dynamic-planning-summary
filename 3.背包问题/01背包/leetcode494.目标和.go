package _1背包

// 每遍历一个元素都有哦两种决策，要么变成正数，要么变成负数

// 暴力递归
func findTargetSumWays(nums []int, target int) int {
	var travel func(index, sum int) int
	travel = func(index, sum int) int {
		if index == len(nums) && sum == target {
			return 1
		}
		if index == len(nums) {
			return 0
		}
		// 加正号
		a := travel(index+1, sum+nums[index])
		// 加负号
		b := travel(index+1, sum-nums[index])
		return a + b
	}
	return travel(0, 0)
}

// 动态规划
// index:0-N
// sum变化范围-1000-1000之间
func findTargetSumWays2(nums []int, target int) int {
	N := len(nums)
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, 2001)
	}
	for i := -1000; i <= 1000; i++ {
		if i == target {
			dp[N][i+1000] = 1
		}
	}
	for i := N - 1; i >= 0; i-- {
		for j := -1000; j <= 1000; j++ {
			a := 0
			if j+nums[i]+1000 < 2001 && j+nums[i]+1000 >= 0 {
				a = dp[i+1][j+nums[i]+1000]
			}
			b := 0
			if j-nums[i]+1000 < 2001 && j-nums[i]+1000 >= 0 {
				b = dp[i+1][j-nums[i]+1000]
			}
			dp[i][j+1000] = a + b

		}
	}
	return dp[0][1000]
}
