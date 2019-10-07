package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res:=true
	helper(root,&res)
	return res
}

func helper(root *TreeNode, res *bool) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, res) + 1
	right := helper(root.Right, res) + 1
	n := left - right
	if n < -1 || n > 1 {
		*res = false
	}
	return max(left,right)
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {

}
