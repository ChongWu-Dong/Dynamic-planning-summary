package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 打家劫舍
// 偷当前的房屋，就要找到下一个可以遍历的房屋
func robIII(root *TreeNode) int {

	var travel func(node *TreeNode) int

	travel = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 偷当前节点
		a := node.Val
		if node.Left != nil && node.Left.Left != nil {
			a += travel(node.Left.Left)
		}
		if node.Left != nil && node.Left.Right != nil {
			a += travel(node.Left.Right)
		}
		if node.Right != nil && node.Right.Left != nil {
			a += travel(node.Right.Left)
		}
		if node.Right != nil && node.Right.Right != nil {
			a += travel(node.Right.Right)
		}

		//不偷当前节点
		b := 0
		if node.Left != nil {
			b += travel(node.Left)
		}
		if node.Right != nil {
			b += travel(node.Right)
		}
		return Max(a, b)
	}
	return travel(root)
}

func MaxIII(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划，记忆化搜索
// 傻缓存
func robIII2(root *TreeNode) int {

	dp := map[*TreeNode]int{}
	var travel func(node *TreeNode) int
	travel = func(node *TreeNode) int {
		_, ok := dp[node]
		if ok {
			return dp[node]
		}
		if node == nil {
			dp[node] = 0
			return 0
		}
		// 偷当前节点
		a := node.Val
		if node.Left != nil && node.Left.Left != nil {
			a += travel(node.Left.Left)
		}
		if node.Left != nil && node.Left.Right != nil {
			a += travel(node.Left.Right)
		}
		if node.Right != nil && node.Right.Left != nil {
			a += travel(node.Right.Left)
		}
		if node.Right != nil && node.Right.Right != nil {
			a += travel(node.Right.Right)
		}
		//不偷当前节点
		b := 0
		if node.Left != nil {
			b += travel(node.Left)
		}
		if node.Right != nil {
			b += travel(node.Right)
		}
		dp[node] = Max(a, b)
		return Max(a, b)
	}
	return travel(root)
}
