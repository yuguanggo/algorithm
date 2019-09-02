package main

import (
	"sort"
	"fmt"
)

/**
俄罗斯套娃信封问题
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

示例:

输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 */

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes)==0{
		return 0
	}
	//将信封按宽升序排，高度降序排
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0]<envelopes[j][0]{
			return true
		}else if envelopes[i][0]==envelopes[j][0]{
			if envelopes[i][1]>envelopes[j][1]{
				return true
			}
			return false
		}else{
			return false
		}
	})
	//遍历信封，将高度组成一个数组
	height:=make([]int,len(envelopes))
	for i:=0;i<len(envelopes);i++{
		height[i]=envelopes[i][1]
	}
	//查找高度的最长上升子序列的长度就是该问题的答案
	return lenghtOfLIS(height)
}

func lenghtOfLIS(nums []int) int  {
	//二分查找法
	top:=make([]int,len(nums))
	//初始化堆
	piles:=0
	for i:=0;i<len(nums);i++{
		left:=0
		right:=piles
		//找到top的最左边界
		for left<right{
			mid:=left+(right-left)/2
			if top[mid]<nums[i]{
				left=mid+1
			}else if top[mid]>nums[i]{
				right=mid
			}else {
				right=mid
			}
		}
		if left==piles{
			piles++
		}
		top[left]=nums[i]
	}
	return piles
}

func lenghtOfLIS2(nums []int) int  {
	//动态规划解法 时间复杂度o(n2)
	n:=len(nums)
	dp:=make([]int,n)
	//将数组初始化成1 dp[i]代表 nums[0..i]的最长上升子序列的长度
	for i:=0;i<n;i++{
		dp[i]=1
	}
	//动态方程 if nums[i]>num[j] j为0...j中的一位，dp[i]=max(dp[i],dp[j]+1)
	for i:=1;i<n;i++{
		for j:=0;j<i;j++{
			if nums[i]>nums[j]{
				dp[i]=max(dp[i],dp[j]+1)
			}
		}
	}
	res:=1
	//遍历数组，找出其中的最大值
	for i:=0;i<len(dp);i++{
		res=max(res,dp[i])
	}
	return res
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	nums:=[]int{10,9,2,5,3,7,101,18}
	fmt.Println(lenghtOfLIS(nums))
}
