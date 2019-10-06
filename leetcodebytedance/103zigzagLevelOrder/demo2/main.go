package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root==nil{
		return res
	}
	q:=make([]*TreeNode,0)
	q=append(q,root)
	level:=0
	for len(q)>0{
		lq:=make([]*TreeNode,0)
		lr:=make([]int,0)
		for len(q)>0{
			top:=q[0]
			q=q[1:]
			if top.Left!=nil{
				lq=append(lq,top.Left)
			}
			if top.Right!=nil{
				lq=append(lq,top.Right)
			}
			if level%2==0{
				//从左往右
				lr=append(lr,top.Val)
			}else{
				//从右往左
				lr=append([]int{top.Val},lr...)
			}
		}
		level++
		q=lq
		res=append(res,lr)
	}
	return res
}
func main() {
	
}
