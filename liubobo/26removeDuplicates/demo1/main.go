package main

func removeDuplicates(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}
	k:=1 //[0..k]为不重复
	pre:=nums[0]
	for i:=1;i<n;i++{
		if nums[i]!=pre{
			pre=nums[i]
			nums[i],nums[k]=nums[k],nums[i]
			k++
		}
	}
	return k
}

func main() {
	
}
