package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res:=make([][]int,0)
	for i:=0;i<len(nums);i++{
		if nums[i]>0{
			break
		}
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		target:=0-nums[i]
		//在剩下的区间里寻找2数之和等于target
		l:=i+1
		r:=len(nums)-1
		for l<r{
			if nums[l]+nums[r]==target{
				res=append(res,[]int{nums[i],nums[l],nums[r]})
				for l<r&&nums[l]==nums[l+1]{
					l++
				}
				for l<r&&nums[r]==nums[r-1]{
					r--
				}
				l++
				r--
			}else if nums[l]+nums[r]>target{
				r--
			}else{
				l++
			}
		}
	}
	return res
}





func main() {
	
}
