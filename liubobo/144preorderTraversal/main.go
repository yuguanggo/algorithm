package main

import "fmt"

/**
144. 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Command struct {
	s string // go ,print
	node *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	if root==nil{
		return res
	}
	stack:=make([]Command,0)
	stack=append(stack,Command{"go",root})
	for len(stack)>0{
		command:=stack[len(stack)-1]
		stack=stack[0:len(stack)-1]
		if command.s=="print"{
			res = append(res,command.node.Val)
		}else{
			if command.node.Right!=nil{
				stack=append(stack,Command{"go",command.node.Right})
			}
			if command.node.Left!=nil{
				stack=append(stack,Command{"go",command.node.Left})
			}
			stack=append(stack,Command{"print",command.node})
		}
	}
	return res
}

func preorderTraversal3(root *TreeNode) []int {
	stack:=make([]*TreeNode,0)
	res:=make([]int,0)
	if root==nil{
		return res
	}
	stack = append(stack,root)
	for len(stack)>0{
		node:=stack[len(stack)-1]
		stack=stack[0:len(stack)-1]
		res = append(res,node.Val)
		if node.Right!=nil{
			stack=append(stack,node.Right)
		}
		if node.Left!=nil{
			stack=append(stack,node.Left)
		}

	}
	return res
}










func preorderTraversal2(root *TreeNode) []int {
	var res = make([]int,0)
	helper(root,&res)
	return res
}

func helper(root *TreeNode,res *[]int)  {
	if root!=nil{
		*res=append(*res,root.Val)
		if root.Left!=nil{
			helper(root.Left,res)
		}
		if root.Right!=nil{
			helper(root.Right,res)
		}
	}
}



func main() {
	var root *TreeNode=nil
	fmt.Println(preorderTraversal(root))
}
