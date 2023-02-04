package 青蛙跳台阶类问题_爬楼梯_

//  每一次的只有两个选择，从爬一个台阶还是爬两个台阶
//  开始的台阶也有两种选择，一种是从第0个台阶出发，另一种是从第一个台阶出发
// 	定义一个递归函数，表示到达第i个台阶的最小费用
func minCostClimbingStairs(cost []int) int {
	var dfs func(n int) int // 表示从第n个台阶出发的最小费用
	dfs = func(n int) int {
		if n == len(cost) { // 当n == len（cost）的时候，已经没有台阶了，这时候费用为0
			return 0
		}
		// 跳一步
		a := cost[n] + dfs(n+1)
		// 跳两步
		b := cost[n]
		if n+2 <= len(cost) {

			b += dfs(n + 2)
		}
		return min(a, b)
	}
	return min(dfs(0), dfs(1))
}

// 改成动态规划
func minCostClimbingStairs2(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := len(cost) - 1; i >= 0; i-- {
		b := cost[i]
		if i+2 <= len(cost) {
			b += dp[i+2]
		}
		a := cost[i] + dp[i+1]
		dp[i] = min(a, b)
	}
	return min(dp[0], dp[1])
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
