package main

import "fmt"

/**
404. 左叶子之和
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	queue:=make([]*TreeNode,0)
	t:=make([]int,0)//1表示是左子树，0表示其他
	if root==nil{
		return 0
	}

	sum:=0
	queue=append(queue,root)
	t=append(t,0)
	for len(queue)>0{
		node:=queue[0]
		queue=queue[1:]
		tt:=t[0]
		t=t[1:]
		if tt==1&&node.Left==nil&&node.Right==nil{
			sum+=node.Val
		}
		if node.Left!=nil{
			t=append(t,1)
			queue=append(queue,node.Left)
		}
		if node.Right!=nil{
			t=append(t,0)
			queue=append(queue,node.Right)
		}
	}
	return sum
}
func sumOfLeftLeaves2(root *TreeNode) int {
	var sum int
	dfs(root,&sum)
	return sum
}

func dfs(root *TreeNode,sum *int)  {
	if root==nil{
		return
	}
	if root.Left!=nil&&root.Left.Left==nil&&root.Left.Right==nil{
		*sum+=root.Left.Val
	}
	dfs(root.Left,sum)
	dfs(root.Right,sum)
}
func main() {
	root:=&TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:4,
				Left:nil,
				Right:nil,
			},
			Right:&TreeNode{
				Val:5,
				Left:nil,
				Right:nil,
			},
		},
		Right:&TreeNode{
			Val:3,
			Left:nil,
			Right:nil,
		},
	}
	fmt.Println(sumOfLeftLeaves(root))
}
