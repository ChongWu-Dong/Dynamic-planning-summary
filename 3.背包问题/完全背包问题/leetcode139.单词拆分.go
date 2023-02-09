package 完全背包问题

import "fmt"

// 可以视为一个完全背包问题，字典dict中的单词数量可以无限取，每次取一个都要和目标单词匹配
// 如果匹配则对目标单词进行操作

// 暴力递归
func wordBreak(s string, wordDict []string) bool {
	var dfs func(restLen int) bool
	dfs = func(restLen int) bool { // 递归表示的是，长度为restLen的字符串s能否被wordDict匹配
		if restLen == 0 {
			return true
		}
		for i := 0; i < len(wordDict); i++ {
			if isSubWord(s, wordDict[i], restLen) { // 表示当前的wordDict[i]可以被匹配，则进入下一步的递归
				if dfs(restLen - len(wordDict[i])) {
					return true
				}
			}
		}
		return false
	}
	return dfs(len(s))
}

func isSubWord(target string, dict string, restLen int) bool { // 这个函数的意思是，如果匹配则返回处理后的target，并返回true。
	index := len(target) - restLen
	if len(dict) > restLen {
		return false
	}
	for i := 0; i < len(dict); i++ {
		if dict[i] != target[index] {
			return false
		} else {
			index++
		}
	}
	return true
}

// 动态规划
func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if isSubWord(s, wordDict[j], i) {
				if dp[i-len(wordDict[j])] {
					dp[i] = true
				}
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "aaa"
	fmt.Println(s[3:])
}
