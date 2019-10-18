package main


func findMaxForm(strs []string, m int, n int) int {
 	dp:=make([][]int,m+1)
 	for i:=0;i<=m;i++{
 		dp[i]=make([]int,n+1)
	}
	for i:=0;i<len(strs);i++{
		//算出字符串消耗了多少个0和1
		zeros:=0
		ones:=0
		for j:=0;j<len(strs[i]);j++{
			if strs[i][j]=='0'{
				zeros++
			}else{
				ones++
			}
		}
		for k:=m;k>=zeros;k--{
			for l:=n;l>=ones;l--{
				dp[k][l]=max(dp[k][l],dp[k-zeros][l-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}


func main() {
	
}
