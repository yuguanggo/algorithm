package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals)==0{
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[j][1]>=intervals[i][1]{
			return true
		}
		return false
	})
	res:=1
	pre:=0
	for i:=1;i<len(intervals);i++{
		if intervals[i][0]>=intervals[pre][1]{
			res++
			pre=i
		}
	}
	return len(intervals)-res
}



func main() {
	
}
