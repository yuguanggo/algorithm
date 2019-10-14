package main

import "fmt"

func solve(board [][]byte)  {
	if len(board)==0{
		return
	}
	//从边缘向里探测
	for i:=0;i<len(board);i++{
		//处理第一列和最后一列
		if board[i][0]=='O'{
			dfs(board,i,0)
		}
		if board[i][len(board[0])-1]=='O'{
			dfs(board,i,len(board[0])-1)
		}
	}
	for i:=0;i<len(board[0]);i++{
		//处理第一行和最后一行
		if board[0][i]=='O'{
			dfs(board,0,i)
		}
		if board[len(board)-1][i]=='O'{
			dfs(board,len(board)-1,i)
		}
	}
	//将没有更改的O变成X #变回O
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if board[i][j]=='O'{
				board[i][j]='X'
			}
			if board[i][j]=='#'{
				board[i][j]='O'
			}
		}
	}
}

func dfs(board [][]byte,x,y int)  {
	//现将边缘的O都填成#
	board[x][y]='#'
	d:=[4][2]int{{-1,0},{1,0},{0,1},{0,-1}}
	for i:=0;i<4;i++{
		newx:=d[i][0]+x
		newy:=d[i][1]+y
		if newx>=0&&newx<len(board)&&newy>=0&&newy<len(board[0])&&board[newx][newy]=='O'{
			dfs(board,newx,newy)
		}
	}
}

func main() {
	var bord=[][]byte{{'X','O','X','X'},{'O','X','O','X'},{'X','O','X','O'},{'O','X','O','X'},{'X','O','X','O'},{'O','X','O','X'}}
	solve(bord)
	fmt.Println(bord)
}
