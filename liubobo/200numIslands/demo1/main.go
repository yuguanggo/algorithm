package main


func numIslands(grid [][]byte) int {
	res:=0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]=='1'{
				dfs(grid,i,j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte,x,y int)  {
	grid[x][y]=0
	d:=[4][2]int{{-1,0},{1,0},{0,1},{0,-1}}
	for i:=0;i<4;i++{
		newx:=d[i][0]+x
		newy:=d[i][1]+y
		if newx>=0&&newx<len(grid)&&newy>=0&&newy<len(grid[0])&&grid[newx][newy]=='1'{
			dfs(grid,newx,newy)
		}
	}
}

func main() {
	
}
