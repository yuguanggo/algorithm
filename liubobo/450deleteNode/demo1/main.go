package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val>key{
		root.Left=deleteNode(root.Left,key)
		return root
	}else if root.Val<key{
		root.Right=deleteNode(root.Right,key)
		return root
	}else{
		//没有左子树
		if root.Left==nil{
			right:=root.Right
			root=nil
			return right
		}
		//没有右子树
		if root.Right==nil{
			left:=root.Left
			root=nil
			return left
		}
		//左右子树都有
		//找到右子树的最小值
		successor:=minmum(root.Right)
		//删除最小值
		successor.Right=removeMin(root.Right)
		successor.Left=root.Left
		return successor
	}
}

func removeMin(root *TreeNode) *TreeNode  {
	if root.Left==nil{
		right:=root.Right
		root=nil
		return right
	}
	root.Left=removeMin(root.Left)
	return root
}

func minmum(root *TreeNode) *TreeNode  {
	if root.Left==nil{
		return root
	}
	return minmum(root.Left)
}


func main() {
	
}
