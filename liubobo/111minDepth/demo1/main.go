package main

import (
	"math"
)

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
	res:=math.MaxInt32
	if root.Left!=nil{
		res=min(minDepth(root.Left),res)
	}
	if root.Right!=nil{
		res=min(minDepth(root.Right),res)
	}
	return res+1
}

func min(i,j int)int  {
	if i>j{
		return j
	}
	return i
}

func main() {
	
}
