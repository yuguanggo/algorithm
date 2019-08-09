package main

/**
37. 解数独
编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。



一个数独。



答案被标成红色。

Note:

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
00-22
03-25
06-28
30-52
33-55
36-58
60-82
63-85
66-88
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sudoku-solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func solveSudoku(board [][]byte) {

	boxUsed := [3][3][10]bool{}
	xUsed:=[9][10]bool{}
	yUsed:=[9][10]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				c := int(board[i][j] - '0')
				xUsed[i][c]=true
				yUsed[j][c]=true
				boxUsed[i/3][j/3][c]=true
			}
		}
	}
	dfs(board, 0, 0, boxUsed,xUsed,yUsed)
}

func dfs(board [][]byte, x, y int, boxUsed [3][3][10]bool,xUsed [9][10]bool,yUsed [9][10]bool) bool {
	if y==9{
		y=0
		x++
		if x==9{
			return true
		}
	}
	if board[x][y]=='.'{
		for j := 1; j <= 9; j++ {
			//开始填数字
			if !boxUsed[x/3][y/3][j]&&!xUsed[x][j]&&!yUsed[y][j] {
				boxUsed[x/3][y/3][j]=true
				xUsed[x][j]=true
				yUsed[y][j]=true
				board[x][y]=byte(j+'0')
				if dfs(board, x, y+1, boxUsed,xUsed,yUsed){
					return true
				}
				board[x][y]='.'
				boxUsed[x/3][y/3][j]=false
				xUsed[x][j]=false
				yUsed[y][j]=false

			}
		}
	}else{
		return dfs(board,x,y+1,boxUsed,xUsed,yUsed)
	}
	return false
}




func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)

}
