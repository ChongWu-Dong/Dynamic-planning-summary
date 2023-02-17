package main

// 暴力递归
func rob(nums []int) int {
	var travel func(index int) int
	travel = func(index int) int {
		if index >= len(nums) { // 遍历到数组长度，没有可以偷的房子了
			return 0
		}
		// 偷这一家
		a := nums[index]
		// 筛除掉无效的情况
		if index+2 < len(nums) {
			a += travel(index + 2)
		}
		// 不偷这一家
		b := travel(index + 1)
		return Max(a, b)
	}
	return travel(0)
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划
func rob2(nums []int) int {
	N := len(nums)
	// dp
	dp := make([]int, N+1)
	for i := N - 1; i >= 0; i-- {
		// 偷这家
		a := nums[i]
		if i+2 < N {
			a += dp[i+2]
		}
		b := dp[i+1]
		dp[i] = Max(a, b)
	}
	return dp[0]

}
func main() {

}
