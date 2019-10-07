package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func countNodes(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return countNodes(root.Left)+countNodes(root.Right)+1
}

func countNodes2(root *TreeNode) int {
	q:=make([]*TreeNode,0)
	res:=0
	if root==nil{
		return res
	}
	q=append(q,root)
	for len(q)>0{
		node:=q[0]
		q=q[1:]
		res++
		if node.Left!=nil{
			q=append(q,node.Left)
		}
		if node.Right!=nil{
			q=append(q,node.Right)
		}
	}
	return res
}
func main() {
	
}
