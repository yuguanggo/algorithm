package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func pathSum2(root *TreeNode, sum int) int {
	if root==nil{
		return 0
	}
	res:=helper(root,sum)
	res+=pathSum2(root.Left,sum)
	res+=pathSum2(root.Right,sum)
	return res
}

func helper(root *TreeNode,sum int) int  {
	if root==nil{
		return 0
	}
	res:=0
	if root.Val==sum{
		res++
	}
	res+=helper(root.Left,sum-root.Val)
	res+=helper(root.Right,sum-root.Val)
	return res
}

func pathSum(root *TreeNode, sum int) int {
	count:=0
	nodes:=make([]int,0)
	path(root,sum,&count,nodes)
	return count
}

func path(root *TreeNode,sum int,count *int,nodes []int)  {
	if root==nil{
		return
	}
	if sum==root.Val{
		*count++
	}
	n:=len(nodes)
	s:=root.Val
	for i:=n-1;i>=0;i--{
		s+=nodes[i]
		if s==sum{
			*count++
		}
	}
	nodes=append(nodes,root.Val)
	path(root.Left,sum,count,nodes)
	path(root.Right,sum,count,nodes)
}


func main() {
	
}
