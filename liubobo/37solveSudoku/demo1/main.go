package main

func solveSudoku(board [][]byte)  {
	xUsed:=[9][10]bool{}
	yUsed:=[9][10]bool{}
	boxUsed:=[3][3][10]bool{}
	//初始化
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]!='.'{
				num:=int(board[i][j]-'0')
				xUsed[i][num]=true
				yUsed[j][num]=true
				boxUsed[i/3][j/3][num]=true
			}
		}
	}
	dfs(board,0,0,xUsed,yUsed,boxUsed)
}

func dfs(board [][]byte,x,y int,xUsed,yUsed [9][10]bool,boxUsed [3][3][10]bool) bool {
	if y==9{
		y=0
		x++
		if x==9{
			return true
		}
	}
	if board[x][y]=='.'{
		//开始填数
		for i:=1;i<9;i++{
			if !xUsed[x][i]&&!yUsed[y][i]&&!boxUsed[x/3][y/3][i]{
				xUsed[x][i]=true
				yUsed[y][i]=true
				boxUsed[x/3][y/3][i]=true
				board[x][y]=byte(i+'0')
				if dfs(board,x,y+1,xUsed,yUsed,boxUsed){
					return true
				}
				board[x][y]='.'
				xUsed[x][i]=false
				yUsed[y][i]=false
				boxUsed[x/3][y/3][i]=false
			}
		}
	}else{
		return dfs(board,x,y+1,xUsed,yUsed,boxUsed)
	}
	return false
}
func main() {
	
}
