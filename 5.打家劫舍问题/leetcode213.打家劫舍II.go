package main

// 正确的做法是，小偷为了使偷的钱最多，所以开始偷的房屋的做法只有两种可能
// 从第0个房间开始偷，否则就从第1个房间开始偷

func rob3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	var travel func(start, end int) int
	travel = func(start, end int) int {
		if start == end {
			return 0
		}
		a := nums[start]
		if start+2 < end {
			a += travel(start+2, end)
		}
		b := travel(start+1, end)
		return Max(a, b)
	}
	n1 := travel(0, len(nums)-1)
	n2 := travel(1, len(nums))
	return Max(n1, n2)
}

// 正确的动态规划
// 将打家劫舍II转换成打家劫舍问题
// 小偷为了使得偷的最多，则0和1必然会进行偷盗
// 0和1偷盗的时候遍历的范围不一样，找出该偷盗范围，则又变成了打家劫舍问题了
func rob4(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	var dpf func(start, end int) int
	dpf = func(start, end int) int {
		dp := make([]int, end+1)
		for i := end - 1; i >= start; i-- {
			// 偷
			a := nums[i]
			if i+2 < end {
				a += dp[i+2]
			}
			// 不偷
			b := dp[i+1]
			dp[i] = Max(a, b)
		}
		return dp[start]
	}
	res1 := dpf(0, len(nums)-1)
	res2 := dpf(1, len(nums))
	return Max(res1, res2)
}
