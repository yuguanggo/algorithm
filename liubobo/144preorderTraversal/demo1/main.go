package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func preorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	if root==nil{
		return res
	}
	stack:=make([]*TreeNode,0)
	stack=append(stack,root)
	for len(stack)>0{
		top:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		res=append(res,top.Val)
		if top.Right!=nil{
			stack=append(stack,top.Right)
		}
		if top.Left!=nil{
			stack=append(stack,top.Left)
		}
	}
	return res
}




func helper(root *TreeNode,res *[]int)  {
	if root!=nil{
		*res=append(*res,root.Val)
		helper(root.Left,res)
		helper(root.Right,res)
	}
}

func main() {
	
}
