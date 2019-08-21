package main

/**
474. 一和零
在计算机界中，我们总是追求用有限的资源获取最大的收益。

现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。

你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。

注意:

给定 0 和 1 的数量都不会超过 100。
给定字符串数组的长度不会超过 600。
示例 1:

输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
输出: 4

解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。
示例 2:

输入: Array = {"10", "0", "1"}, m = 1, n = 1
输出: 2

解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ones-and-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
dp[m][n]=max(dp[m][n],dp[m-strs[i]m][n-strs[i]n]+1)
*/



func findMaxForm(strs []string, m int, n int) int {
	if len(strs)==0{
		return 0
	}
	dp:=make([][]int,m+1)
	for i:=0;i<m+1;i++{
		dp[i]=make([]int,n+1)
	}
	for i:=0;i<len(strs);i++{
		ones:=0
		zeros:=0
		for j:=0;j<len(strs[i]);j++{
			if strs[i][j]=='1'{
				ones++
			}else{
				zeros++
			}
		}

		for l:=m;l>=zeros;l--{
			for k:=n;k>=ones;k--{
				dp[l][k]=max(dp[l][k],dp[l-zeros][k-ones]+1)
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
