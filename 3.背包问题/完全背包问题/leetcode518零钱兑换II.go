package 完全背包问题

// 这是一个完全背包问题，所有的物品都可以重复的选取

// 可以采用回溯算法求解和动态规划算法求解

// 回溯
func change(amount int, coins []int) int {
	// 结果
	cout := 0

	var backtracking func(sum, startIndex int)
	backtracking = func(sum, startIndex int) {
		if sum == amount { // 收集答案
			cout++
			return
		}
		if sum > amount { // 剪枝
			return
		}
		for i := startIndex; i < len(coins); i++ {
			sum += coins[i]
			backtracking(sum, i)
			sum -= coins[i]
		}
	}
	backtracking(0, 0)
	return cout
}

// 动态规划的暴力递归
func change2(amount int, coins []int) int {

	var travel func(sum, startIndex int) int
	travel = func(sum, startIndex int) int {
		cout := 0
		if startIndex == len(coins) {
			return 0
		}
		// 表示组合有效
		if sum == amount {
			return 1
		}
		// 表示组合无效
		if sum > amount {
			return 0
		}
		for i := startIndex; i < len(coins); i++ {
			cout += travel(sum+coins[i], i)
		}
		return cout
	}
	return travel(0, 0)
}

// 动态规划
// 记忆化搜索
func change3(amount int, coins []int) int {
	N := len(coins)
	// dp
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < amount+1; j++ {
			dp[i][j] = -1
		}
	}

	var travel func(sum, startIndex int) int
	travel = func(sum, startIndex int) int {
		if dp[startIndex][sum] != -1 {
			return dp[startIndex][sum]
		}
		cout := 0
		if startIndex == len(coins) {
			cout = 0
		} else if sum == amount {
			cout = 1
		} else if sum > amount {
			cout = 0
		} else {
			for i := startIndex; i < len(coins); i++ {
				if sum+coins[i] <= amount {
					cout += travel(sum+coins[i], i)
				}
			}
		}
		dp[startIndex][sum] = cout
		return cout

	}
	return travel(0, 0)
}

// 动态规划
func change4(amount int, coins []int) int {
	N := len(coins)
	// dp
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := N; i >= 0; i-- {
		for j := amount; j >= 0; j-- {
			if i == N {
				dp[i][j] = 0
			} else if j == amount {
				dp[i][j] = 1
			} else {
				cout := 0
				for k := i; k < N; k++ {
					if j+coins[k] <= amount {
						cout += dp[k][j+coins[k]]
					}
				}
				dp[i][j] = cout
			}
		}
	}
	return dp[0][0]

}
