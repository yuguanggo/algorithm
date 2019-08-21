package main

import "fmt"

/**
417. 太平洋大西洋水流问题
给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。

规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。

请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。


提示：

输出坐标的顺序不重要
m 和 n 都小于150
 

示例：

 

给定下面的 5x5 矩阵:

  太平洋 ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * 大西洋

返回:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pacific-atlantic-water-flow
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func pacificAtlantic(matrix [][]int) [][]int {
	m:=len(matrix)
	res:=make([][]int,0)
	if m==0{
		return res
	}
	n:=len(matrix[0])
	d:=[4][2]int{{-1,0},{0,1},{1,0},{0,-1}}

	visitedP:=make([][]bool,m)
	visitedX:=make([][]bool,m)
	for k:=0;k<m;k++{
		visitedP[k]=make([]bool,n)
		visitedX[k]=make([]bool,n)
	}
	//从边缘开始遍历
	for i:=0;i<n;i++{
		dfs(matrix,0,i,d,visitedP)
		dfs(matrix,m-1,i,d,visitedX)
	}
	for i:=0;i<m;i++{
		dfs(matrix,i,0,d,visitedP)
		dfs(matrix,i,n-1,d,visitedX)
	}
	for x:=0;x<m;x++{
		for y:=0;y<n;y++{
			if visitedP[x][y]&&visitedX[x][y]{
				res=append(res,[]int{x,y})
			}
		}
	}
	return res
}

func dfs(matrix [][]int,startx,starty int,d [4][2]int,visited [][]bool) {
	visited[startx][starty]=true
	//开始遍历四个方向
	for i:=0;i<4;i++{
		newx:=startx+d[i][0]
		newy:=starty+d[i][1]
		if inarea(newx,newy,len(matrix),len(matrix[0]))&&!visited[newx][newy]&&matrix[startx][starty]<=matrix[newx][newy]{
			dfs(matrix,newx,newy,d,visited)
		}
	}
}

func inarea(x,y,m,n int)bool  {
	return x>=0&&x<m&&y>=0&&y<n
}

func main() {
	var matrix=[][]int{{1,1},{1,1},{1,1}}
	fmt.Println(pacificAtlantic(matrix))
}
