package main


/**
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	return helper(root,root)

}
func helper(p *TreeNode,q *TreeNode) bool  {
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
func isSymmetric2(root *TreeNode) bool {
	queue:=make([]*TreeNode,0)
	queue=append(queue,root,root)
	for len(queue)>0 {
		q:=queue[0]
		p:=queue[1]
		queue=queue[2:]
		if q==nil&&p==nil{
			continue
		}
		if q==nil||p==nil{
			return false
		}
		if q.Val!=p.Val{
			return false
		}
		queue=append(queue,q.Left)
		queue=append(queue,p.Right)
		queue=append(queue,q.Right)
		queue=append(queue,p.Left)
	}
	return true
}



func main() {
	
}
