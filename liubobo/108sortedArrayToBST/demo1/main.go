package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums,0,len(nums)-1)
}

func helper(nums []int,l,r int) *TreeNode  {
	if l>r{
		return nil
	}
	mid:=l+(r-l)/2
	root:=&TreeNode{
		Val:nums[mid],
		Left:nil,
		Right:nil,
	}
	root.Left=helper(nums,l,mid-1)
	root.Right=helper(nums,mid+1,r)
	return root
}


func main() {
	
}
