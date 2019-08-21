package main

/**
107. 二叉树的层次遍历 II
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	levels:=make([][]int,0)
	if root==nil{
		return levels
	}
	queue:=make([]*TreeNode,0)
	queue=append(queue,root)
	for len(queue)>0{
		ql:=len(queue)
		values:=make([]int,0)
		for i:=0;i<ql;i++{
			node:=queue[0]
			queue=queue[1:]
			values=append(values,node.Val)
			if node.Left!=nil{
				queue=append(queue,node.Left)
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
			}
		}
		levels=append([][]int{values},levels...)
	}
	return levels
}

func main() {
	
}
