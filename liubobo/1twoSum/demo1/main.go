package main

func twoSum(nums []int, target int) []int {
	record:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		k:=target-nums[i]
		if _,ok:=record[k];ok{
			return []int{i,record[k]}
		}
		record[nums[i]]=i
	}
	return []int{}
}


func main() {
	
}
