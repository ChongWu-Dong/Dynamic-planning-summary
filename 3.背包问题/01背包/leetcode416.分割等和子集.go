package _1背包

// 这是一个01背包问题，判断是否可以将数组等分
// 所以这里背包的容量就是数组和的一半
// 如果这里数组的和的一半不是整数说明不以等分直接返回false
// 数组和的一半为背包的容量，判断数组中的元素是否可以将背包装满，也就是背包最大能装下多少东西
// 如果背包可以装满说明可以等分返回true，否则返回false

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 { // sum不能等分，说明这里一定要返回false
		return false
	}
	weight := sum / 2 // 背包容量
	var dfs func(index int, weight int) int
	dfs = func(index int, weight int) int {
		if index == len(nums) {
			return 0
		}
		if weight == 0 {
			return 0
		}
		a := dfs(index+1, weight)
		b := 0
		if weight-nums[index] >= 0 {
			b = dfs(index+1, weight-nums[index]) + nums[index]
		}
		return max(a, b)
	}
	n := dfs(0, weight)
	if n == weight {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划
func canPartition2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	weight := sum / 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, weight+1)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 1; j <= weight; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] = max(dp[i+1][j], nums[i]+dp[i+1][j-nums[i]])
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][weight] == weight
}
