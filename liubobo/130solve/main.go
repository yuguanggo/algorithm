package main

import "fmt"

/**
130. 被围绕的区域
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func solve(board [][]byte)  {
	var m=len(board)
	if m==0{
		return
	}
	var n=len(board[0])
	var d=[4][2]int{{-1,0},{0,1},{1,0},{0,-1}}

	//搜索第一行和最后一行
	for i:=0;i<n;i++{
		if board[0][i]=='O'{
			dfs(board,0,i,d)
		}
		if board[m-1][i]=='O'{
			dfs(board,m-1,i,d)
		}
	}
	//搜索第一列和最后一列
	for j:=0;j<m;j++{
		if board[j][0]=='O'{
			dfs(board,j,0,d)
		}
		if board[j][n-1]=='O'{
			dfs(board,j,n-1,d)
		}
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if board[i][j]=='O'{
				board[i][j]='X'
			}
			if board[i][j]=='#'{
				board[i][j]='O'
			}
		}
	}
}
func dfs(board [][]byte,startx,starty int,d [4][2]int) {
	board[startx][starty]='#'
	for i:=0;i<4;i++{
		newx:=startx+d[i][0]
		newy:=starty+d[i][1]
		if inarea(newx,newy,len(board),len(board[0]))&&board[newx][newy]=='O'{
			//先将边界的改成#
			dfs(board,newx,newy,d)
		}
	}

}
func inarea(x,y,m,n int) bool {
	return x>=0&&x<m&&y>=0&&y<n
}

func main() {
	var bord=[][]byte{{'O','O','O'},{'O','O','O'},{'O','O','O'}}
	solve(bord)
	fmt.Println(bord)
}
