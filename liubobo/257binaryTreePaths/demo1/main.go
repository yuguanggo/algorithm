package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res:=make([]string,0)
	helper(root,"",&res)
	return res
}

func helper(root *TreeNode,path string,res *[]string)  {
	if root==nil{
		return
	}
	if path==""{
		path=strconv.Itoa(root.Val)
	}else{
		path=path+"->"+strconv.Itoa(root.Val)
	}

	if root.Left==nil&&root.Right==nil{
		*res=append(*res,path)
		return
	}
	helper(root.Left,path,res)
	helper(root.Right,path,res)
}


func main() {
	
}
