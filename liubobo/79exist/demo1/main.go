package main


func exist(board [][]byte, word string) bool {
	row:=len(board)
	col:=len(board[0])
	visited:=make([][]bool,row)
	for i:=0;i<row;i++{
		visited[i]=make([]bool,col)
	}
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if dfs(board,visited,word,0,i,j){
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte,visited [][]bool,word string,index int,x,y int) bool   {
	if index==len(word)-1{
		return board[x][y]==word[index]
	}
	d:=[4][2]int{{-1,0},{0,1},{1,0},{0,-1}}
	if board[x][y]!=word[index]{
		return false
	}
	visited[x][y]=true
	for i:=0;i<4;i++{
		newx:=d[i][0]+x
		newy:=d[i][1]+y
		if newx>=0&&newx<len(board)&&newy>=0&&newy<len(board[0])&&!visited[newx][newy]&&dfs(board,visited,word,index+1,newx,newy){
			return true
		}
	}
	visited[x][y]=false
	return false
}



func main() {
	
}
