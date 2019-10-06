package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func levelOrder(root *TreeNode) [][]int {
	q:=make([]*TreeNode,0)
	res:=make([][]int,0)
	if root==nil{
		return res
	}
	q=append(q,root)
	for len(q)>0{
		level:=make([]*TreeNode,0)
		l:=make([]int,0)
		for len(q)>0{
			node:=q[0]
			q=q[1:]
			l=append(l,node.Val)
			if node.Left!=nil{
				level=append(level,node.Left)
			}
			if node.Right!=nil{
				level=append(level,node.Right)
			}
		}
		res=append(res,l)
		q=level
	}
	return res
}


func main() {
	
}
