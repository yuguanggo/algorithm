package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func isValidBST(root *TreeNode) bool {
	nodes:=make([]int,0)
	inOrder(root,&nodes)
	for i:=1;i<len(nodes);i++{
		if nodes[i]<=nodes[i-1]{
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode,nodes *[]int)  {
	if root!=nil{
		inOrder(root.Left,nodes)
		*nodes=append(*nodes,root.Val)
		inOrder(root.Right,nodes)
	}
}


func main() {
	
}
