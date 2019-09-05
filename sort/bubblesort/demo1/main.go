package main

import "fmt"

func bubblesSort(nums []int)  {
	n:=len(nums)
	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			if nums[j]>nums[j+1]{
				nums[j],nums[j+1]=nums[j+1],nums[j]
			}
		}
	}
}

func main() {
	nums:=[]int{3,21,3,1,0}
	bubblesSort(nums)
	fmt.Println(nums)
}
