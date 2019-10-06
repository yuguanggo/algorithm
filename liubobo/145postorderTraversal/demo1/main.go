package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func postorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	if root==nil{
		return res
	}
	stack:=make([]*TreeNode,0)
	stack=append(stack,root)
	for len(stack)>0{
		top:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		s:=[]int{top.Val}
		res=append(s,res...)
		if top.Left!=nil{
			stack=append(stack,top.Left)
		}
		if top.Right!=nil{
			stack=append(stack,top.Right)
		}
	}
	return res
}

func helper(root *TreeNode,res *[]int)  {
	if root!=nil{
		helper(root.Left,res)
		helper(root.Right,res)
		*res=append(*res,root.Val)
	}
}
func main() {
	
}
