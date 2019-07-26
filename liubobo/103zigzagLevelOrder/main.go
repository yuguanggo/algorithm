package main

/**
103. 二叉树的锯齿形层次遍历
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue:=make([]*TreeNode,0)
	levels:=make([][]int,0)
	if root==nil{
		return levels
	}
	level:=0
	queue=append(queue,root)
	for len(queue)>0{
		levels=append(levels,[]int{})
		ql:=len(queue)
		node:=new(TreeNode)

		for i:=0;i<ql;i++{
			node=queue[0]
			queue=queue[1:]
			if level%2==0{
				levels[level]=append(levels[level],node.Val)
			}else{
				levels[level]=append([]int{node.Val},levels[level]...)
			}
			if node.Left!=nil {
				queue=append(queue,node.Left)
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
			}
		}
		level++
	}
	return levels
}

func main() {
	
}
