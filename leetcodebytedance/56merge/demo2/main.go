package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0]<intervals[j][0]{
			return true
		}
		return false
	})
	n:=len(intervals)
	res:=make([][]int,0)
	i:=0
	for i<n{
		l:=intervals[i][0]
		r:=intervals[i][1]
		j:=i+1
		for j<n&&intervals[j][0]<=r{
			if intervals[j][1]>r{
				r=intervals[j][1]
			}
			j++
		}
		res=append(res,[]int{l,r})
		i=j
	}
	return res
}

func main() {
	
}
