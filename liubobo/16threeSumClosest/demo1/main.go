package main

import (
	"sort"
	"math"
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
			if math.Abs(float64(target-sum))<math.Abs(float64(target-ans)){
				ans=sum
			}
			if sum>target{
				r--
			}else if sum<target{
				l++
			}else{
				return ans
			}

		}
	}
	return ans
}


func main() {
	
}
