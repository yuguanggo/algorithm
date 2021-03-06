package main



/**
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root==nil{
		return 0
	}
	queue:=make([]*TreeNode,0)
	queue=append(queue,root)
	res:=0
	for len(queue)>0{
		q:=queue[0]
		queue=queue[1:]
		res++
		if q.Left!=nil{
			queue=append(queue,q.Left)
		}
		if q.Right!=nil{
			queue=append(queue,q.Right)
		}
	}
	return res
}

func countNodes2(root *TreeNode) int {
	if root==nil{
		return 0
	}

	return countNodes2(root.Left)+countNodes2(root.Right)+1
}

func main() {
	
}
