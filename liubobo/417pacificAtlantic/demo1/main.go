package main

func pacificAtlantic(matrix [][]int) [][]int {
	res:=make([][]int,0)
	row:=len(matrix)
	if row==0{
		return res
	}
	col:=len(matrix[0])
	visitedP:=make([][]bool,row)
	visitedA:=make([][]bool,row)
	for i:=0;i<row;i++{
		visitedP[i]=make([]bool,col)
		visitedA[i]=make([]bool,col)
	}
	//从边缘开始遍历
	for i:=0;i<row;i++{
		//第一列
		dfs(matrix,i,0,visitedP)
		//最后一列
		dfs(matrix,i,col-1,visitedA)
	}
	for i:=0;i<col;i++{
		//第一行
		dfs(matrix,0,i,visitedP)
		//最后一行
		dfs(matrix,row-1,i,visitedA)
	}
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if visitedP[i][j]&&visitedA[i][j]{
				res=append(res,[]int{i,j})
			}
		}
	}
	return res
}

func dfs(matrix [][]int,x,y int,visited [][]bool)  {
	visited[x][y]=true
	d:=[4][2]int{{-1,0},{1,0},{0,1},{0,-1}}
	for i:=0;i<4;i++{
		newx:=d[i][0]+x
		newy:=d[i][1]+y
		if newx>=0&&newx<len(matrix)&&newy>=0&&newy<len(matrix[0])&&!visited[newx][newy]&&matrix[newx][newy]>=matrix[x][y]{
			dfs(matrix,newx,newy,visited)
		}
	}
}


func main() {
	
}
