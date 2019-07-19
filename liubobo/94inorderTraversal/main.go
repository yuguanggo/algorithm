package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack:=make([]*TreeNode,0)
	res:=make([]int,0)
	if root==nil{
		return res
	}
	cur:=root
	for cur!=nil||len(stack)>0{
		for cur!=nil{
			stack=append(stack,cur)
			cur=cur.Left
		}
		node:=stack[len(stack)-1]
		stack=stack[0:len(stack)-1]
		res = append(res,node.Val)
		cur=node.Right
	}
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	res:=make([]int,0)
	helper(root,&res)
	return res
}
func helper(root *TreeNode,res *[]int)   {
	if root!=nil{
		if root.Left!=nil{
			helper(root.Left,res)
		}
		*res=append(*res,root.Val)
		if root.Right!=nil{
			helper(root.Right,res)
		}
	}
}

func main() {
	res:=make([]int,0)
	res=append(res,1)
	res[0]=3
}
