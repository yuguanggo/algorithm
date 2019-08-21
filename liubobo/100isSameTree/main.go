package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil&&q==nil{
		return true
	}
	if p==nil||q==nil{
		return false
	}
	if p.Val!=q.Val{
		return false
	}
	return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	queue:=make([][2]*TreeNode,0)
	queue=append(queue,[2]*TreeNode{p,q})
	for len(queue)>0{
		item:=queue[0]
		queue=queue[1:]
		if !check(item[0],item[1]){
			return false
		}
		if item[0]!=nil{
			queue=append(queue,[2]*TreeNode{item[0].Left,item[1].Left})
			queue=append(queue,[2]*TreeNode{item[0].Right,item[1].Right})
		}

	}
	return true
}

func check(p *TreeNode,q *TreeNode) bool  {
	if p==nil&&q==nil{
		return true
	}
	if p==nil||q==nil{
		return false
	}
	if p.Val!=q.Val{
		return false
	}
	return true
}

func main() {
	
}
