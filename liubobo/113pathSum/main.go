package main

import "fmt"

/**
113. 路径总和 II
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	queue:=make([]*TreeNode,0)
	res:=make([][]int,0)
	path:=make([][]int,0)
	if root==nil{
		return res
	}
	queue=append(queue,root)
	path=append(path,[]int{root.Val})
	for len(queue)>0{
		node:=queue[0]
		queue=queue[1:]
		p:=path[0]
		path=path[1:]
		if node.Left==nil&&node.Right==nil&&checkSum(p,sum){
			res=append(res,p)
		}
		if node.Left!=nil{
			queue=append(queue,node.Left)
			var p1=make([]int,len(p))
			copy(p1,p)
			p1=append(p1,node.Left.Val)
			path=append(path,p1)
		}
		if node.Right!=nil{
			queue=append(queue,node.Right)
			var p2 =make([]int,len(p))
			copy(p2,p)
			p2=append(p2,node.Right.Val)
			path=append(path,p2)
		}
	}
	return res
}
func checkSum(path []int,sum int)bool  {
	for _,v:=range path{
		sum-=v
	}
	return sum==0
}
func pathSum2(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int
	return dfs(root,path,res,sum)
}

func dfs(root *TreeNode,path []int,res [][]int,sum int)[][]int{
	if root==nil{
		return res
	}
	sum-=root.Val
	path=append(path,root.Val)
	if root.Left==nil&&root.Right==nil{
		if sum==0{
			s:=make([]int,len(path))
			copy(s,path)
			res=append(res,s)
		}
		return res
	}
	if root.Left!=nil{
		res=dfs(root.Left,path,res,sum)
	}
	if root.Right!=nil{
		res=dfs(root.Right,path,res,sum)
	}
	return res
}

func main() {
	var p1 =make([]int,len([]int{1,2,4}))
	copy(p1,[]int{2,3,4})
	fmt.Println(p1)
}
