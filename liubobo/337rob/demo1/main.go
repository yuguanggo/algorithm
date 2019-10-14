package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res:=dfs(root)
	return max(res[0],res[1])
}

func dfs(root *TreeNode) [2]int  {
	if root==nil{
		return [2]int{0,0}
	}
	l:=dfs(root.Left)
	r:=dfs(root.Right)
	return [2]int{max(l[0],l[1])+max(r[0],r[1]),root.Val+l[0]+r[0]}
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {

}
