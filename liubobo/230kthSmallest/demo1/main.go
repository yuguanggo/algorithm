package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func kthSmallest(root *TreeNode, k int) int {
	res:=math.MaxInt32
	inOrder(root,&k,&res)
	return res
}

func inOrder(root *TreeNode,k *int,res *int)   {
	if root!=nil{
		inOrder(root.Left,k,res)
		if *res!=math.MaxInt32{
			return
		}
		*k--
		if *k==0{
			*res=root.Val
		}
		inOrder(root.Right,k,res)
	}
}

func main() {
	
}
