package main

import "fmt"

/**
437. 路径总和 III
给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root==nil{
		return 0
	}
	res:=findPath(root,sum)
	res+=pathSum(root.Left,sum)
	res+=pathSum(root.Right,sum)
	return res
}

func findPath(root *TreeNode,num int) int  {
	if root==nil{
		return 0
	}
	res:=0
	if num==root.Val{
		res+=1
	}
	res+=findPath(root.Left,num-root.Val)
	res+=findPath(root.Right,num-root.Val)
	return res
}

func main() {
	tree:=&TreeNode{
		Val:10,
		Left:&TreeNode{
			Val:5,
			Left:&TreeNode{
				Val:3,
				Left:&TreeNode{
					Val:3,
					Left:nil,
					Right:nil,
				},
				Right:&TreeNode{
					Val:-2,
					Left:nil,
					Right:nil,
				},
			},
			Right:&TreeNode{
				Val:2,
				Left:nil,
				Right:&TreeNode{
					Val:1,
					Left:nil,
					Right:nil,
				},
			},
		},
		Right:&TreeNode{
			Val:-3,
			Left:nil,
			Right:&TreeNode{
				Val:11,
				Left:nil,
				Right:nil,
			},
		},
	}
	fmt.Println(pathSum(tree,8))
}
