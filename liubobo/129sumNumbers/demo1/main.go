package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}




func sumNumbers(root *TreeNode) int {
	sum:=0
	helper(root,[]int{},&sum)
	return sum
}

func helper(root *TreeNode,path []int,sum *int)  {
	if root==nil{
		return
	}
	path=append(path,root.Val)
	if root.Left==nil&&root.Right==nil{
		for i:=0;i<len(path);i++{
			m:=1
			for j:=i;j<len(path)-1;j++{
				m=m*10
			}
			*sum=*sum+path[i]*m
		}
		return
	}
	helper(root.Left,path,sum)
	helper(root.Right,path,sum)
}


func main() {
	
}
