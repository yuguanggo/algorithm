package main

/**
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val>key{
		root.Left=deleteNode(root.Left,key)
		return root
	}else if root.Val<key{
		root.Right=deleteNode(root.Right,key)
		return root
	}else{
		//没有左子树
		if root.Left==nil{
			right:=root.Right
			root=nil
			return right
		}
		//没有右子树
		if root.Right==nil{
			left:=root.Left
			root=nil
			return left
		}
		//左右子树都在的情况下，当前节点用右子树的最小值替代
		successor:=minmum(root.Right)
		successor.Right=removeMin(root.Right)
		successor.Left=root.Left
		root=nil
		return successor
	}
}
//删除已node为根节点的最小值并返回根节点
func removeMin(node *TreeNode) *TreeNode  {
	if node.Left==nil{
		right:=node.Right
		node=nil
		return right
	}
	node.Left=removeMin(node.Left)
	return node
}

//找出以node为根节点的最小值
func minmum(node *TreeNode) *TreeNode {
	if node.Left==nil{
		return node
	}
	return minmum(node.Left)
}

func main() {
	
}
