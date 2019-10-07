package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func pathSum(root *TreeNode, sum int) [][]int {
	res:=make([][]int,0)
	helper(root,[]int{},&res,sum)
	return res
}

func helper(root *TreeNode,path []int,res *[][]int,sum int)  {
	if root==nil{
		return
	}
	path=append(path,root.Val)
	if sum==root.Val&&root.Left==nil&&root.Right==nil{
		p:=make([]int,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	helper(root.Left,path,res,sum-root.Val)
	helper(root.Right,path,res,sum-root.Val)
}


func main() {
	
}
