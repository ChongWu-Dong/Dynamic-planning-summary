package 青蛙跳台阶类问题_爬楼梯_

import "strconv"

// 有点类似与青蛙跳台的问题，青蛙可以跳一级或者两级
//这个题和青蛙跳台阶的问题类似。青蛙可以选择只跳一步，也可以选择跳两步。
//该题也是类似的，翻译单词的时候可以选择只翻译一个，也可以选择翻译两个。只有两种请情况
//但是要注意的是：翻译两个单词的情况的时候，要保证现在存在的数字数量大于等于2，并且要保证这两个数字的组合要在10和25之间
//因为两个数字的组合06是不符合条件的，所以要保证他们在10和25之间。
//
//从暴力递归改写成动态规划
func translateNum(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	var dfs func(n int) int // 函数的意思是，现在含有的位数是n位，有多少种翻译的方法、
	dfs = func(n int) int {
		if n == 0 { // 已经没有字母可以翻译了,表明翻译到最后了，这是一个方案
			return 1
		}
		a := 0
		// 选择翻译当前的字母,选择一个字母
		a += dfs(n - 1)
		// 选择两个字母
		if n >= 2 {
			if str[n-2:n] < "26" && str[n-2:n] >= "0" {
				a += dfs(n - 2)
			}
		}
		return a
	}
	return dfs(n)
}

// 改写成动态规划
func translateNum2(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1]
		if i >= 2 {
			if str[i-2:i] < "26" && str[i-2:i] >= "10" { // 要注意，这里的范围为10-25
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n]
}
