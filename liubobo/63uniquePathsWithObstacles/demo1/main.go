package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m:=len(obstacleGrid)
	if m==0{
		return 0
	}
	n:=len(obstacleGrid[0])
	dp:=make([]int,n)
	//初始化第一行
	for i:=0;i<n;i++{
		if obstacleGrid[0][i]==1{
			dp[i]=0
			break
		}else{
			dp[i]=1
		}
	}
	iscol:=false
	for i:=1;i<m;i++{
		if obstacleGrid[i][0]==1||iscol{
			//初始化第一列的值
			dp[0]=0
			iscol=true
		}
		for j:=1;j<n;j++{
			if obstacleGrid[i][j]==1{
				dp[j]=0
			}else{
				dp[j]=dp[j-1]+dp[j]
			}
		}
	}
	return dp[n-1]
}


func main() {
	
}
