package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	if root.Left==nil&&root.Right==nil{
		return 1
	}
	mind:=math.MaxInt32
	if root.Left!=nil{
		mind=min(minDepth(root.Left),mind)
	}
	if root.Right!=nil{
		mind=min(minDepth(root.Right),mind)
	}
	return mind+1
}
func min(i,j int) int{
	if i>j{
		return j
	}
	return i
}
func main() {
	
}
