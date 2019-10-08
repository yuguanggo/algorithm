package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==q||root==p||root==nil{
		return root
	}
	left:=lowestCommonAncestor(root.Left,p,q)
	right:=lowestCommonAncestor(root.Right,p,q)
	if left!=nil&&right!=nil{
		return root
	}else if left!=nil{
		return left
	}else{
		return right
	}
}
func main() {
	
}
