package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n:=len(nums)
	sort.Ints(nums)
	ans:=nums[0]+nums[1]+nums[2]
	for i:=0;i<n;i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		l:=i+1
		r:=n-1
		for l<r{
			sum:=nums[i]+nums[l]+nums[r]
			if am
			for nums[l]+nums[r]-target1<gap{
				gap=nums[l]+nums[r]-target1
				sum=nums[i]+nums[l]+nums[r]
				l++
			}
			if nums[l]+nums[r]-target1>=gap{
				r--
			}
		}
	}
	return sum
}


func main() {
	
}
