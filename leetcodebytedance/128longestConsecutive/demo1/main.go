package main

import "fmt"

func longestConsecutive(nums []int) int {
	max:=0
	record:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		if record[nums[i]]==0{
			l:=record[nums[i]-1]
			r:=record[nums[i]+1]
			len:=l+r+1
			if max<len{
				max=len
			}
			record[nums[i]-l]=len
			record[nums[i]+r]=len
			record[nums[i]]=len
		}
	}
	return max
}


func main() {
	nums:=[]int{100,4,200,1,3,2}
	fmt.Println(longestConsecutive(nums))
}
