package main

import "fmt"

/**
200. 岛屿数量
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func numIslands(grid [][]byte) int {
	m:=len(grid)
	if m==0{
		return 0
	}
	n:=len(grid[0])
	d:=[4][2]int{{-1,0},{0,1},{1,0},{0,-1}}
	res:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid[i][j]=='1'{
				res++
				//遍历当前岛屿并标记
				dfs(grid,i,j,d)
			}
		}
	}
	return res
}

func dfs(grid [][]byte,startx,starty int,d [4][2]int)  {
	//从四个方向标记岛屿
	for i:=0;i<4;i++{
		newx:=startx+d[i][0]
		newy:=starty+d[i][1]
		if inArea(newx,newy,len(grid),len(grid[0]))&&grid[newx][newy]=='1'{
			grid[newx][newy]='0'
			dfs(grid,newx,newy,d)
		}
	}
}

func inArea(x,y,m,n int)bool  {
	return x>=0&&x<m&&y>=0&&y<n
}

func inarea(x,y,m,n int)bool  {
	return x>=0&&x<m&&y>=0&&y<n
}

func numIslands2(grid [][]byte) int {
	m:=len(grid)
	if m==0{
		return 0
	}
	n:=len(grid[0])
	d:=[4][2]int{{-1,0},{0,1},{1,0},{0,-1}}
	queue:=make([][2]int,0)
	res:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid[i][j]=='1'{
				res++
				//广度优先遍历
				grid[i][j]='0'
				queue=append(queue,[2]int{i,j})
				for len(queue)>0{
					q:=queue[0]
					queue=queue[1:]
					//将q点置成已访问过

					//开始遍历q节点的四周
					for k:=0;k<4;k++{
						newx:=q[0]+d[k][0]
						newy:=q[1]+d[k][1]
						if inarea(newx,newy,m,n)&&grid[newx][newy]=='1'{
							grid[newx][newy]='0'
							queue=append(queue,[2]int{newx,newy})
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	var grid=[][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}
	fmt.Println(numIslands(grid))
}
