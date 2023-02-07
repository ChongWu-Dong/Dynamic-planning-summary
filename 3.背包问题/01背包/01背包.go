package _1背包

// 在 0 1 背包中，每次便利的时候也只有两选择，就是当前的物品取或者不取
// 只有两种情况，取当前的物品不取当前的物品，然后返回最大值
// 通过暴力递归改写动态规划

func bag(weight []int, value []int, bagMax int) int {

	var travel func(num int, value []int, weight []int, bagMax int, val int) int
	travel = func(num int, value []int, weight []int, bagMax int, val int) int {
		if bagMax < 0 {
			return 0
		}
		//if bagMax == 0 { // 包的重量等于0不一定要返回，因为可能存在货物的重量为0
		//	return val
		//}
		if num == len(weight) {
			return val
		}
		a := travel(num+1, value, weight, bagMax-weight[num], val+value[num]) // 要当前的货
		b := travel(num+1, value, weight, bagMax, val)                        // 不要当前的货
		return Max(a, b)
	}
	res := travel(0, value, weight, bagMax, 0)
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
