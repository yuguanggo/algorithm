package main

import (
	"fmt"
)

/**
动态方程
dp[i,j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
dp[j]=min(dp[j-1],pre,dp[j])
 */

func maximalSquare(matrix [][]byte) int {
	//优化后的动态规划
	row:=len(matrix)
	if row==0{
		return 0
	}
	col:=len(matrix[0])
	dp:=make([]int,col+1)
	maxl:=0
	pre:=0
	for i:=1;i<=row;i++{
		for j:=1;j<=col;j++{
			temp:=dp[j]
			if matrix[i-1][j-1]=='1'{
				dp[j]=min(dp[j],min(dp[j-1],pre))+1
				if maxl<dp[j]{
					maxl=dp[j]
				}
			}else{
				dp[j]=0
			}
			pre=temp
		}
	}
	return maxl*maxl
}
func maximalSquare2(matrix [][]byte) int {
	row:=len(matrix)
	if row==0{
		return 0
	}
	col:=len(matrix[0])
	//初始化
	dp:=make([][]int,row+1)
	for i:=0;i<=row;i++{
		dp[i]=make([]int,col+1)
	}
	maxlen:=0
	for i:=1;i<=row;i++{
		for j:=1;j<=col;j++{
			if matrix[i-1][j-1]=='1'{
				dp[i][j]=min(dp[i-1][j-1],min(dp[i-1][j],dp[i][j-1]))+1
				if maxlen<dp[i][j]{
					maxlen=dp[i][j]
				}
			}
		}
	}
	return maxlen*maxlen
}

func  min(x,y int)int  {
	if x>y{
		return y
	}
	return x
}

func main() {
	matrix:=[][]byte{{'1'}}
	fmt.Println(maximalSquare(matrix))
}
