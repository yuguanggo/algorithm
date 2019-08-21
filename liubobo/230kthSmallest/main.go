package main

import (
	"math"
	"fmt"
)

/**
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func kthSmallest(root *TreeNode, k int) int {
	var res int
	res= math.MaxInt32
	inOrder(root,&k,&res)
	return res
}

func inOrder(root *TreeNode,k,res *int)  {
	if root!=nil{
		if root.Left!=nil{
			inOrder(root.Left,k,res)
		}
		if *res!=math.MaxInt32{
			return
		}
		*k--
		if *k==0{
			*res=root.Val
		}
		if root.Right!=nil{
			inOrder(root.Right,k,res)
		}
	}
}

func kthSmallest2(root *TreeNode, k int) int  {
	stack:=make([]*TreeNode,0)
	cur:=root
	for cur!=nil||len(stack)>0{
		for cur!=nil{
			stack=append(stack,cur)
			cur=cur.Left
		}
		node:=stack[len(stack)-1]
		stack=stack[0:len(stack)-1]
		k--
		if k==0{
			return node.Val
		}
		cur=node.Right
	}
	return -1
}

func main() {
	var res int
	res=1
	fmt.Println(res)
}
