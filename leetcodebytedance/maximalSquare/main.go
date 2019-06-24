package main

import "fmt"

/**
最大正方形
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 1 1 1 1

输出: 4
dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
 */
func maximalSquare2(matrix [][]byte) int {
	if len(matrix)==0{
		return 0
	}
	m:=len(matrix)
	n:=len(matrix[0])
	sz:=0
	var dp [][]int
	//初始化
	for i:=0;i<m;i++{
		temp:=make([]int,n)
		dp = append(dp,temp)
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0||j==0||matrix[i][j]==0{
				dp[i][j]=int(matrix[i][j])
			}else{
				dp[i][j]=Min(dp[i-1][j-1],Min(dp[i][j-1],dp[i-1][j]))+1
			}
			sz=Max(dp[i][j],sz)
		}
	}
	return sz*sz
}
func maximalSquare(matrix [][]byte) int {
	if len(matrix)==0{
		return 0
	}
	m:=len(matrix)
	n:=len(matrix[0])
	sz:=0
	pre:=0
	var cur =make([]int,n)
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			tmp:=cur[j]
			if i==0||j==0||matrix[i][j]==0{
				cur[j]=int(matrix[i][j])
			}else{
				cur[j]=Min(pre,Min(cur[j-1],cur[j]))+1
			}
			sz=Max(cur[j],sz)
			pre=tmp
		}
	}
	return sz*sz
}

func Min(i,j int) int  {
	if i<j{
		return i
	}
	return j
}

func Max(i,j int) int  {
	if i<j{
		return j
	}
	return i
}

func main() {

	var matrix =[][]byte{{1,0,1,0,0},{1,0,1,1,1},{1,1,1,1,1},{1,1,1,1,1}}
	fmt.Println(maximalSquare(matrix))
	//m:=5
	//dp:=make(map[int][m]int)
	//dp[1][1]=2
	//fmt.Println(dp)
}
