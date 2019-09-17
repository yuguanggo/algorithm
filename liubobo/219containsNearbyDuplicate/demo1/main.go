package main

func containsNearbyDuplicate(nums []int, k int) bool {
	n:=len(nums)
	record:=make(map[int]int)
	for i:=0;i<n;i++{
		if _,ok:=record[nums[i]];ok{
			return true
		}
		record[nums[i]]=i
		if len(record)>k{
			delete(record,nums[i-k])
		}
	}
	return false
}


func main() {
	
}
