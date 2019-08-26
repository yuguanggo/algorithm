package main

import "fmt"

/**
300. 最长上升子序列
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
dp[i]=max(dp[i],1+dp[j] if nums[i]>nums[j])
 */

 //动态规划
func lengthOfLIS(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}
	dp:=make([]int,n+1)
	for i:=0;i<n;i++{
		dp[i]=1
	}
	for i:=1;i<n;i++{
		for j:=0;j<i;j++{
			if nums[i]>nums[j]{
				dp[i]=max(dp[i],dp[j]+1)
			}
		}
	}
	res:=1
	for i:=0;i<n;i++{
		res=max(res,dp[i])
	}
	return res
}

func lengthOfLIS2(nums []int) int{
	n:=len(nums)
	if n<2{
		return n
	}
	tail:=make([]int,0)
	tail=append(tail,nums[0])
	for i:=1;i<n;i++{
		if nums[i]>tail[len(tail)-1]{
			tail=append(tail,nums[i])
			continue
		}
		left:=0
		right:=len(tail)-1
		for left<right{
			mid:=left+(right-left)/2
			if tail[mid]<nums[i]{
				left=mid+1
			}else{
				right=mid
			}
		}
		tail[left]=nums[i]
	}
	return len(tail)
}

func max(i,j int)int  {
	if i>j{
		return i
	}
	return j
}


func main() {
	nums:=[]int{3,5,6,2,5,4,5,19,5,6,7,12}
	fmt.Println(lengthOfLIS2(nums))
}
