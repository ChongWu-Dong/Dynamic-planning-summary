package 青蛙跳台阶类问题_爬楼梯_

import (
	"fmt"
	"sort"
)

// 这个题和青蛙跳台阶还有机器人走路径等问题类似，类似的点在于，这个题我们每次的决策的可能数都是可知的
// 只有两种可能的决策，那就是，遍历到当前的兼职，选择做还是不做，做的话就收集结果，然后跳到下一个可以符合条件的兼职，如果不做的话，则直接进行下一个兼职。

//我们可以任意挑选一个工作开始，对于每一份工作，可以选择做，也可以选择不做，
// 如果不做，那就可以从下一份工作开始，
// 如果做，那么下一份工作必须要在当前工作结束之后
// 因此我们先将工作按照开始时间从小到大排序，
// 然后设计一个函数 dfs(i) 表示从第 i 份工作开始(当前工作不一定做)，可以获得的最大报酬。
// 最终答案即为 dfs(0)，虽然我们可以从任意一份工作开始，但是由于每一份工作都考虑了做与不做两种情况，所以从第1份工作开始考虑，就一定考虑到了全部可能性
//
// 函数 dfs(i) 的计算过程如下：
//
// 对于第 i 份工作，我们可以选择做，也可以选择不做。
//     如果不做，最大报酬就是 dfs(i+1)；
//      如果做，我们可以通过二分查找，找到在第 i 份工作结束时间之后开始的第一份工作，记为 j，那么最大报酬就是 profit[i]+dfs(j)。
// 取两者的较大值即可。

// 暴力递归
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	// 对兼职进行排序，按照开始时间从低到高
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	// 设置一个递归函数，返回兼职的收益最大化
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == len(jobs) {
			return 0
		}
		// 选择做还是不做
		// 不做
		a := dfs(i + 1)
		// 做
		// 做的话，先找到下一次的兼职开始的地方
		k := sort.Search(len(jobs), func(index int) bool {
			return jobs[index][0] >= jobs[i][1]
		})
		b := jobs[i][2] + dfs(k)
		return max1235(a, b)
	}
	return dfs(0)
}

// 改写成动态规划
func jobScheduling2(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	// 对兼职进行排序，按照开始时间从低到高
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		// 不做
		a := dp[i+1]
		// 做
		k := sort.Search(len(jobs), func(index int) bool {
			return jobs[index][0] >= jobs[i][1]
		})
		b := jobs[i][2] + dp[k]
		dp[i] = max1235(a, b)
	}
	return dp[0]
}

func max1235(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 用于测试sort.Search（）的用法，该函数会返回jobs长度的一个索引，如果没找到判断条件部位true则返回输入的长度
// 该测试表明，如果在没有找到jobs【index】>= 1000的时候，返回的索引是len（jobs）

func main() {
	// 该测试表明，如果在没有找到jobs【index】>= 1000的时候，返回的索引是len（jobs）
	jobs := []int{1, 2, 3, 4}
	k := sort.Search(len(jobs), func(index int) bool {
		return jobs[index] >= 2
		// return jobs[index] >= 1000
	})
	fmt.Println(k)
}
