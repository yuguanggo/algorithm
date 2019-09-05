package main

import "fmt"

func insertSort(nums []int)  {
	n:=len(nums)
	for i:=1;i<n;i++{
		for j:=i;j>0;j--{
			if nums[j]<nums[j-1]{
				nums[j],nums[j-1]=nums[j-1],nums[j]
			}
		}
	}
}
func main() {
	nums:=[]int{3,21,3,1,0}
	insertSort(nums)
	fmt.Println(nums)
}
