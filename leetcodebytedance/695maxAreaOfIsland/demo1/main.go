package main

func maxAreaOfIsland(grid [][]int) int {
	row:=len(grid)
	if row==0{
		return 0
	}
	col:=len(grid[0])
	res:=0
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if grid[i][j]!=0{
				res = max(res,dfs(grid,i,j))
			}
		}
	}
	return res
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func dfs(grid [][]int,x,y int) int  {
	row:=len(grid)
	col:=len(grid[0])
	d:=[4][2]int{{-1,0},{0,1},{1,0},{0,-1}}
	sum:=1
	grid[x][y]=0
	for i:=0;i<4;i++{
		newx:=x+d[i][0]
		newy:=y+d[i][1]
		if inArea(newx,newy,row,col)&&grid[newx][newy]==1{
			sum+=dfs(grid,newx,newy)
		}
	}
	return sum
}

func inArea(x,y,row,col int) bool  {
	return x>=0&&x<row&&y>=0&&y<col
}




func main() {
	
}
