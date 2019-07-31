package main


/**
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}
	return insert(nums,0,len(nums)-1)
}

//二分法
func insert(nums []int,l,r int) *TreeNode  {
	if l>r{
		return nil
	}
	mid:=l+(r-l)/2
	root:=&TreeNode{
		Val:nums[mid],
		Left:nil,
		Right:nil,
	}
	root.Left=insert(nums,l,mid-1)
	root.Right=insert(nums,mid+1,r)
	return root
}

func main() {
	
}
