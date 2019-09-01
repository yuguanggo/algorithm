package main

import "fmt"

/**
* 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil {
		return right
	} else if right == nil {
		return left
	}
	return nil
}
func main() {
	t4:=&TreeNode{
		Val:4,
		Left:nil,
		Right:nil,
	}
	t7:=&TreeNode{
		Val:7,
		Left:nil,
		Right:nil,
	}
	t2:=&TreeNode{
		Val:2,
		Left:t7,
		Right:t4,
	}
	t6:=&TreeNode{
		Val:6,
		Left:nil,
		Right:nil,
	}
	t5:=&TreeNode{
		Val:5,
		Left:t6,
		Right:t2,
	}
	t0:=&TreeNode{
		Val:0,
		Left:nil,
		Right:nil,
	}
	t8:=&TreeNode{
		Val:8,
		Left:nil,
		Right:nil,
	}
	t1:=&TreeNode{
		Val:1,
		Left:t0,
		Right:t8,
	}
	t3:=&TreeNode{
		Val:3,
		Left:t5,
		Right:t1,
	}

	fmt.Println(lowestCommonAncestor(t3,t4,t5))

}
