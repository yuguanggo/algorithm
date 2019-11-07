package main

import "fmt"

func lengthOfLIS2(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}
	dp:=make([]int,n+1)
	for i:=0;i<n;i++{
		dp[i]=1
	}
	for i:=1;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[i]>nums[j]{
				dp[i]=max(dp[i],dp[j]+1)
			}
		}
	}
	res:=1
	for i:=0;i<len(dp);i++{
		res=max(res,dp[i])
	}
	return res
}

func lengthOfLIS(nums []int) int {
	if len(nums)==0{
		return 0
	}
	q:=make([]int,0)
	q=append(q,nums[0])
	for i:=1;i<len(nums);i++{
		if nums[i]>q[len(q)-1]{
			q=append(q,nums[i])
			continue
		}
		//找到左边界
		l:=0
		r:=len(q)-1
		for l<=r{
			mid:=l+(r-l)/2
			if q[mid]==nums[i]{
				r=mid-1
			}else if q[mid]<nums[i]{
				l=mid+1
			}else{
				r=mid-1
			}
		}
		q[l]=nums[i]
	}
	return len(q)
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	nums:=[]int{10,9,2,5,3,4}
	fmt.Println(lengthOfLIS(nums))
}
