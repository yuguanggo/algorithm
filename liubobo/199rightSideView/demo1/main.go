package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func rightSideView(root *TreeNode) []int {
	res:=make([]int,0)
	if root==nil{
		return res
	}
	q:=make([]*TreeNode,0)
	q=append(q,root)
	for len(q)>0{
		lr:=make([]*TreeNode,0)
		//存储每层的最后一个节点值
		res=append(res,q[len(q)-1].Val)
		for len(q)>0{
			node:=q[0]
			q=q[1:]
			if node.Left!=nil{
				lr=append(lr,node.Left)
			}
			if node.Right!=nil{
				lr=append(lr,node.Right)
			}
		}
		q=lr
	}
	return res
}


func main() {
	
}
