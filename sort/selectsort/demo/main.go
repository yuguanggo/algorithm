package main

import "fmt"

func selectSort(nums []int)  {
	n:=len(nums)
	for i:=0;i<n;i++{
		minIndex:=i
		for j:=i+1;j<n;j++{
			if nums[j]<nums[minIndex]{
				minIndex=j
			}
		}
		nums[minIndex],nums[i]=nums[i],nums[minIndex]
	}
}
func main() {
	nums:=[]int{3,21,3,1,0}
	selectSort(nums)
	fmt.Println(nums)
}
