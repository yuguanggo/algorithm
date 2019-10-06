package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func levelOrderBottom(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root==nil{
		return res
	}
	q:=make([]*TreeNode,0)
	q=append(q,root)
	for len(q)>0{
		level:=make([]*TreeNode,0)
		lres:=make([]int,0)
		for len(q)>0{
			top:=q[0]
			q=q[1:]
			lres=append(lres,top.Val)
			if top.Left!=nil{
				level=append(level,top.Left)
			}
			if top.Right!=nil{
				level=append(level,top.Right)
			}
		}
		res=append([][]int{lres},res...)
		q=level
	}
	return res
}


func main() {
	
}
