package main

import "sort"

func threeSum(nums []int) [][]int {
	n:=len(nums)
	sort.Ints(nums)
	res:=make([][]int,0)
	for i:=0;i<n;i++{
		if nums[i]>0{
			break
		}
		if i>=1&&nums[i]==nums[i-1]{
			continue
		}
		target:=0-nums[i]
		l:=i+1
		r:=n-1
		for l<r{
			if nums[l]+nums[r]==target{
				res=append(res,[]int{nums[i],nums[l],nums[r]})
				for l<r&&nums[l+1]==nums[l]{
					l++
				}
				for l<r&&nums[r-1]==nums[i]{
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
