package main


/**
78. 子集
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func subsets(nums []int) [][]int {
	var res=make([][]int,0)
	helper(nums,0,[]int{},&res)
	return res
}

func helper(nums []int,index int,path []int,res *[][]int)  {
	var p=make([]int,len(path))
	copy(p,path)
	*res=append(*res,p)
	for i:=index;i<len(nums);i++{
		path=append(path,nums[i])
		helper(nums,i+1,path,res)
		path=path[0:len(path)-1]
	}
}

func main() {
	
}
