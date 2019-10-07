package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	if p.Val>root.Val&&q.Val>root.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}
	if p.Val<root.Val&&q.Val<root.Val{
		return lowestCommonAncestor(root.Left,p,q)
	}
	return root
}

func main() {
	
}
