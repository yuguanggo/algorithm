package main

func wordBreak(s string, wordDict []string) bool {
	//dp[i]表示从[0..i]是否可以拆分 dp[i]=dp[j]+in(s[i-j])
	n:=len(s)
	dp:=make([]bool,n+1)
	dp[0]=true
	for i:=1;i<=n;i++{
		for j:=0;j<i;j++{
			if dp[j]&&iscontain(s[j:i],wordDict){
				dp[i]=true
				break
			}
		}
	}
	return dp[n]
}

func iscontain(s string,wordDict []string) bool  {
	for _,v:=range wordDict{
		if v==s{
			return true
		}
	}
	return false
}



func main() {
	
}
