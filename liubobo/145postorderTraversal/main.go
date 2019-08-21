package main


/**
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	stack:=make([]*TreeNode,0)
	res:=make([]int,0)
	if root==nil{
		return res
	}
	stack=append(stack,root)
	for len(stack)>0{
		node:=stack[len(stack)-1]
		stack=stack[0:len(stack)-1]
		s:=[]int{node.Val}
		res=append(s,res...)
		if node.Left!=nil{
			stack=append(stack,node.Left)
		}
		if node.Right!=nil{
			stack=append(stack,node.Right)
		}
	}
	return res
}


func postorderTraversal2(root *TreeNode) []int {
	res:=make([]int,0)
	helper(root,&res)
	return res
}

func helper(root *TreeNode,res *[]int)  {
	if root!=nil{
		if root.Left!=nil{
			helper(root.Left,res)
		}
		if root.Right!=nil{
			helper(root.Right,res)
		}
		*res=append(*res,root.Val)
	}
}

func main() {
	
}
