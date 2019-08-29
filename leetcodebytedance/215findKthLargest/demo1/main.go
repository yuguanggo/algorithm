package main

import (
	"algorithm/sort/sorttesthelper"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	return helper(nums,0,len(nums)-1,k)
}

func helper(nums []int,l,r int,k int) int  {

	//partion
	mid:=l+(r-l)/2
	nums[l],nums[mid]=nums[mid],nums[l]
	v:=nums[l]

	//三路快排 [l..lt]>v [lt+1..i]==v [gt..r]<v
	lt:=l
	i:=l+1
	gt:=r+1
	for i<gt{
		if nums[i]<v{
			nums[i],nums[gt-1]=nums[gt-1],nums[i]
			gt--
		}else if nums[i]>v{
			nums[i],nums[lt+1]=nums[lt+1],nums[i]
			lt++
			i++
		}else{
			i++
		}
	}
	nums[l],nums[lt]=nums[lt],nums[l]
	if k-1<=lt-1{
		return helper(nums,l,lt-1,k)
	}else if k-1>=gt{
		return helper(nums,gt,r,k)
	}else{
		return nums[k-1]
	}
}


func main() {
	nums:=sorttesthelper.GenerateRandomArray(10,1,10)
	fmt.Println(nums)
	helper(nums,0,len(nums)-1)
	fmt.Println(nums)
}
