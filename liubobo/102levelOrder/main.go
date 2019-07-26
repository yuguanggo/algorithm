package main

import "fmt"

/**
102. 二叉树的层次遍历
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levels:=make([][]int,0)
	if root==nil{
		return levels
	}
	helper(root,0,&levels)
	return levels
}
func helper(node *TreeNode,level int,levels *[][]int){
	if level==len(*levels){
		*levels=append(*levels,[]int{})
	}
	(*levels)[level]=append((*levels)[level],node.Val)
	if node.Left!=nil{
		helper(node.Left,level+1,levels)
	}
	if node.Right!=nil{
		helper(node.Right,level+1,levels)
	}
}

func levelOrder2(root *TreeNode) [][]int {
	levels:=make([][]int,0)
	if root==nil{
		return levels
	}
	queue:=make([]*TreeNode,0)
	level:=0
	queue=append(queue,root)
	for len(queue)>0{
		levels=append(levels,[]int{})
		ql:=len(queue)
		for i:=0;i<ql ;i++  {
			node:=queue[0]
			queue=queue[1:]
			levels[level]=append(levels[level],node.Val)
			if node.Left!=nil{
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

func test(arr []int)  {
	fmt.Printf("%p \n",&arr)
	arr[1]=1
}

func main() {
	arr:=make([]int,0)
	arr = append(arr,1)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println(arr[1:])
}
