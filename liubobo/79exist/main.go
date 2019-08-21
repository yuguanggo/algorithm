package main

import "fmt"

/**
79. 单词搜索
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
 var(
 	m int
 	n int
 	d [4][2]int
 	visited [][]bool
 )

func exist(board [][]byte, word string) bool {
	m=len(board)
	if m==0{
		return false
	}
	n=len(board[0])
	d =[4][2]int{{-1,0},{0,1},{1,0},{0,-1}}
	//初始化visited
	visited=make([][]bool,m)
	for i:=0;i<m;i++{
		visited[i]=make([]bool,n)
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if searchWord(board,i,j,word,0){
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte,startx,starty int,word string,index int) bool  {
	if index==len(word)-1{
		return board[startx][starty]==word[index]
	}
	if board[startx][starty]==word[index]{
		visited[startx][starty]=true
		for i:=0;i<4;i++{
			newx:=startx+d[i][0]
			newy:=starty+d[i][1]
			if inArea(newx,newy)&&!visited[newx][newy]&&searchWord(board,newx,newy,word,index+1){
				return true
			}
		}
		visited[startx][starty]=false
	}
	return false
}

func inArea(x,y int) bool {
	return x>=0&&x<m&&y>=0&&y<n
}


func main() {
	//初始化

	var bord=[][]byte{{'A','B','C','D','E'},{'T','S','R','Q','F'},{'M','N','O','P','G'},{'L','K','J','I','H'}}
	fmt.Println(exist(bord,"ABCDEFGHIJKLMNOPQRST"))
}
