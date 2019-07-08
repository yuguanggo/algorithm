package main

import "fmt"

/**
149. 直线上最多的点数
给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。

示例 1:

输入: [[1,1],[2,2],[3,3]]
输出: 3
解释:
^
|
|        o
|     o
|  o  
+------------->
0  1  2  3  4
示例 2:

输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4
解释:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6


 */
func addLines(i,j int,points [][]int,count ,duplicates,horisontal *int,lines map[float64]int) {
	if points[i][0]==points[j][0]&&points[i][1]==points[j][1]{
		*duplicates+=1
	}else if points[i][1]==points[j][1]{
		*horisontal+=1
		*count = max(*horisontal,*count)
	}else {
		slope:= float64((points[i][0]-points[j][0]))/float64((points[i][1]-points[j][1]))
		lines[slope]++
		*count = max(lines[slope]+1,*count)
	}
}

//包含i点 的直线上最多的点数
func maxPointsContainsI(i int,points [][]int) int {
	lines:=make(map[float64]int)//记录i与points[i+1...n-1]两点之间的斜率相同的个数
	count:=1 //初始化
	duplicates:=0 //重复的点得个数
	horisontal:=1//处理和i处于同一水平的点
	for j:=i+1;j<len(points);j++{
		addLines(i,j,points,&count,&duplicates,&horisontal,lines)
	}
	return count+duplicates
}
func maxPoints(points [][]int) int {
	n:=len(points)
	if n<3{
		return n
	}
	maxcount:=1
	for i:=0;i<n;i++{
		maxcount = max(maxPointsContainsI(i,points),maxcount)
	}
	return maxcount
}

func max(i,j int) int  {
	if i<j{
		return j
	}
	return i
}

func main() {
	points:=[][]int{{0,0},{94911151,94911150},{94911152,94911151}}
	fmt.Println(maxPoints(points))
}
