package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res:=make([][]int,0)
	sort.Ints(nums)
	n:=len(nums)
	for k:=0;k<n;k++{
		if nums[k]>0{
			break
		}
		if k>0&&nums[k]==nums[k-1]{
			continue
		}
		//双指针
		target:=0-nums[k]
		l:=k+1
		r:=n-1
		for l<r{
			if nums[l]+nums[r]==target{
				res=append(res,[]int{nums[k],nums[l],nums[r]})
				for l<r&&nums[l]==nums[l+1]{
					l++
				}
				for l<r&&nums[r]==nums[r-1]{
					r--
				}
				l++
				r--
			}else if nums[l]+nums[r]<target{
				l++
			}else{
				r--
			}
		}
	}
	return res
}




func main() {
	nums:=[]int{-1,0,1,2,-1,-4};
	fmt.Println(threeSum(nums))
}
