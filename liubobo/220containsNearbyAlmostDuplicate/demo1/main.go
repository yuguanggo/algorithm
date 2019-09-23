package main

import "math"

//获取桶的id
func getNumId(x,w int) int {
	if x<0{
		return (x+1)/w-1
	}
	return x/w
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t<0{
		return false
	}
	//桶的宽度
	w:=t+1
	record:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		m:=getNumId(nums[i],w)
		if _,ok:=record[m];ok{
			return true
		}
		if _,ok:=record[m-1];ok{
			if math.Abs(float64(nums[i])-float64(record[m-1]))<float64(w){
				return true
			}
		}
		if _,ok:=record[m+1];ok{
			if math.Abs(float64(nums[i])-float64(record[m+1]))<float64(w){
				return true
			}
		}
		record[m]=nums[i]
		if len(record)>k{
			delete(record,getNumId(nums[i-k],w))
		}
	}
	return false
}

func main() {
	
}
