package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root==nil{
		return 0
	}
	q:=make([]*TreeNode,0)
	q=append(q,root)
	oddsum:=0 //偶数层的和
	evensum:=0 //
	for len(q)>0{

	}
}

func main() {

}
