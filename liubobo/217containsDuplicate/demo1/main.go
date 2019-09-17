package main


func containsDuplicate(nums []int) bool {
	record:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		if _,ok:=record[nums[i]];ok{
			return true
		}
		record[nums[i]]=i
	}
	return false
}


func main() {
	
}
