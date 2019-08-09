package main


/**
52. N皇后 II
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回 n 皇后不同的解决方案的数量。

示例:

输入: 4
输出: 2
解释: 4 皇后问题存在如下两个不同的解法。
[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func totalNQueens(n int) int {
	col:=make([]bool,n)//存放第几列已经被占
	dia1:=make([]bool,2*n-1) //存放对角线1 的占用情况
	dia2:=make([]bool,2*n-1) //存放对角线2 的占用情况
	res:=0
	putQueen(n,0,col,dia1,dia2,[]int{},&res)
	return res
}
//尝试在一个n皇后的问题中，在index行摆放皇后
func putQueen(n int,index int,col []bool,dia1 []bool,dia2 []bool,row []int,res *int)  {
	if n==index{
		*res++
	}
	for i:=0;i<n;i++{
		if !col[i]&&!dia1[index+i]&&!dia2[index-i+n-1]{
			col[i]=true
			dia1[index+i]=true
			dia2[index-i+n-1]=true
			row=append(row,i)
			putQueen(n,index+1,col,dia1,dia2,row,res)
			col[i]=false
			dia1[index+i]=false
			dia2[index-i+n-1]=false
			row=row[:len(row)-1]
		}
	}
	return
}
func main() {
	
}
