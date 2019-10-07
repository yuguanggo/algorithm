package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func isSymmetric2(root *TreeNode) bool {
	return helper(root,root)
}

func helper(p,q *TreeNode) bool  {
	if p==nil&&q==nil{
		return true
	}
	if p==nil||q==nil{
		return false
	}
	if q.Val!=p.Val{
		return false
	}
	return helper(p.Left,q.Right)&&helper(p.Right,q.Left)
}

func isSymmetric(root *TreeNode) bool  {
	q:=make([]*TreeNode,0)
	q=append(q,root,root)
	for len(q)>0{
		n1:=q[0]
		n2:=q[1]
		q=q[2:]
		if n1==nil&&n2==nil{
			continue
		}
		if n1==nil||n2==nil{
			return false
		}
		if n1.Val!=n2.Val{
			return false
		}
		q=append(q,n1.Left)
		q=append(q,n2.Right)
		q=append(q,n1.Right)
		q=append(q,n2.Left)
	}
	return true
}

func main() {
	
}
