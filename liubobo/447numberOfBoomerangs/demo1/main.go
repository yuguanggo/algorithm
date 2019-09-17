package main


func numberOfBoomerangs(points [][]int) int {
	n:=len(points)
	res:=0
	for i:=0;i<n;i++{
		record:=make(map[int]int)
		for j:=0;j<n;j++{
			if i!=j{
				record[dis(points[i],points[j])]++
			}
		}
		for _,v:=range record{
			res+=v*(v-1)
		}
	}
	return res
}
func dis(p,q []int) int  {
	return (p[0]-q[0])*(p[0]-q[0])+(p[1]-q[1])*(p[1]-q[1])
}

func main() {
	
}
