package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root==nil{
		return res
	}
	q:=make([]*TreeNode,0)
	level:=0
	q=append(q,root)
	for len(q)>0{
		n:=len(q)
		s:=make([]int,n)
		var node *TreeNode
		for i:=0;i<n;i++{
			node=q[0]
			q=q[1:]
			if level%2==0{
				s[i]=node.Val
			}else{
				s[n-i-1]=node.Val
			}
			if node.Left!=nil{
				q=append(q,node.Left)
			}
			if node.Right!=nil{
				q=append(q,node.Right)
			}
		}
		level++
		res=append(res,s)
	}
	return res
}



func main() {
	t7:=&TreeNode{
		Val:7,
		Left:nil,
		Right:nil,
	}
	t15:=&TreeNode{
		Val:15,
		Left:nil,
		Right:nil,
	}
	t20:=&TreeNode{
		Val:20,
		Left:t15,
		Right:t7,
	}
	t9:=&TreeNode{
		Val:9,
		Left:nil,
		Right:nil,
	}
	t3:=&TreeNode{
		Val:3,
		Left:t9,
		Right:t20,
	}
	fmt.Println(zigzagLevelOrder(t3))
}
