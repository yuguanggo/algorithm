package main

func addLine(i,j int)  {

}


func maxPointsContains(i int,points [][]int)int  {

	count:=1 //斜率相同的点
	dum:=0 //重复的点
	horn:=1 //水平的点
	record:=make(map[float64]int)
	for j:=i+1;j<len(points);j++{
		if points[i][0]==points[j][0]&&points[i][1]==points[j][1]{
			dum++
		}else if points[i][1]==points[j][1]{
			horn++
			count=max(horn,count)
		}else{
			k:=float64((points[i][0]-points[j][0]))/float64((points[i][1]-points[j][1]))
			record[k]++
			count=max(record[k]+1,count)
		}
	}
	return count+dum
}

func maxPoints(points [][]int) int {
	n:=len(points)
	if n<3{
		return n
	}
	maxCount:=0
	for i:=0;i<n;i++{
		maxCount=max(maxPointsContains(i,points),maxCount)
	}
	return maxCount
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	
}
