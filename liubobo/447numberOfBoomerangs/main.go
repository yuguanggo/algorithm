package main


/**
447. 回旋镖的数量
给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。

找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。

示例:

输入:
[[0,0],[1,0],[2,0]]

输出:
2

解释:
两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-boomerangs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
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

func dis(p,q []int) int {
	return (p[0]-q[0])*(p[0]-q[0])+(p[1]-q[1])*(p[1]-q[1])
}

func main() {
	
}
