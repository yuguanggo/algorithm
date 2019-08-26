package main

import (
	"sort"
	"fmt"
)

func threeSum(nums []int) [][]int {
	var (
		res    [][]int
		n      int
		k      int
		i      int
		j      int
		target int
	)
	sort.Ints(nums)
	fmt.Println(nums)
	n = len(nums)
	for k = 0; k < n; k++ {
		if nums[k] > 0 {
			break
		}
		if k>0&&nums[k] == nums[k-1] {
			continue
		}
		target = 0 - nums[k]
		i = k + 1
		j = n - 1
		for i < j {
			if nums[i]+nums[j] == target {
				res = append(res,[]int{nums[k],nums[i],nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}
		}

	}
	return res
}
func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	nenums :=threeSum(nums)
	for _,v:=range nenums{
		fmt.Println(v)
	}
	fmt.Println()
}
