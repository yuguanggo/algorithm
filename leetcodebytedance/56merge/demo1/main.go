package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	res:=make([][]int,0)
	n:=len(intervals)
	if n==0{
		return res
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0]<intervals[j][0]{
			return true
		}
		return false
	})
	start:=intervals[0][0]
	end:=intervals[0][1]
	for i:=1;i<n;i++{
		if intervals[i][0]<=end{
			end=max(intervals[i][1],end)
		}else{
			res=append(res,[]int{start,end})
			start=intervals[i][0]
			end=intervals[i][1]
		}
	}
	res=append(res,[]int{start,end})
	return res
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}


func main() {
	
}
