package main

import (
	"strconv"
)

/**
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func binaryTreePaths(root *TreeNode) []string {
	queue:=make([]*TreeNode,0)
	res:=make([]string,0)
	if root==nil{
		return res
	}
	path:=make([]string,0)
	queue=append(queue,root)
	path=append(path,strconv.Itoa(root.Val))
	for len(queue)>0{
		node:=queue[0]
		queue=queue[1:]
		p:=path[0]
		path=path[1:]
		if node.Left==nil&&node.Right==nil{
			res=append(res,p)
		}
		if node.Left!=nil{
			path=append(path,p+"->"+strconv.Itoa(node.Left.Val))
			queue=append(queue,root.Left)
		}
		if node.Right!=nil{
			path=append(path,p+"->"+strconv.Itoa(node.Right.Val))
			queue=append(queue,root.Right)
		}

	}
	return res
}

func binaryTreePaths2(root *TreeNode) []string {
	var res =make([]string,0)
	helper(root,"",&res)
	return res
}

func helper(root *TreeNode,path string,res *[]string)  {
	if root!=nil{
		path=path+strconv.Itoa(root.Val)
		if root.Left==nil&&root.Right==nil{
			*res=append(*res,path)
		}else{
			path=path+"->"
			helper(root.Left,path,res)
			helper(root.Right,path,res)
		}
	}
}
func main() {
	
}
