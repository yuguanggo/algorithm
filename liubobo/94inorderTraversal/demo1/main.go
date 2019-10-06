package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func inorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	if root==nil{
		return res
	}
	stack:=make([]*TreeNode,0)
	stack=append(stack,root)
	cur:=root
	for len(stack)>0{
		for cur!=nil&&cur.Left!=nil{
			stack=append(stack,cur.Left)
			cur=cur.Left
		}
		top:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		res=append(res,top.Val)
		if top.Right!=nil{
			stack=append(stack,top.Right)
		}
		cur=top.Right
	}
	return res
}



func helper(root *TreeNode,res *[]int)  {
	if root!=nil{
		helper(root.Left,res)
		*res=append(*res,root.Val)
		helper(root.Right,res)
	}
}
func main() {
	l3:=&TreeNode{
		Val:3,
		Left:nil,
		Right:nil,
	}
	l2:=&TreeNode{
		Val:2,
		Left:l3,
		Right:nil,
	}
	l1:=&TreeNode{
		Val:1,
		Left:nil,
		Right:l2,
	}
	fmt.Println(inorderTraversal(l1))
}
