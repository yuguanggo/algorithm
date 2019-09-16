package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n:=len(nums)
	res:=make([][]int,0)
	for i:=0;i<n;i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		target1:=target-nums[i]
		for j:=i+1;j<n;j++{

			if j>i+1&&nums[j]==nums[j-1]{
				continue
			}
			target2:=target1-nums[j]
			l:=j+1
			r:=n-1
			for l<r{
				if nums[l]+nums[r]==target2{
					res=append(res,[]int{nums[i],nums[j],nums[l],nums[r]})
					for l<r&&nums[l+1]==nums[l]{
						l++
					}
					for l<r&&nums[r-1]==nums[r]{
						r--
					}
					l++
					r--
				}else if nums[l]+nums[r]>target2{
					r--
				}else{
					l++
				}
			}
		}
	}
	return res
}


func main() {
	
}
