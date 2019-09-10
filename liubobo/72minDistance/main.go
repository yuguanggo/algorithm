package main

/**
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
示例 1:

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2:

输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

dp[i][j] 表示word1[0..i]替换成 word2[0..j]的最小距离

有三种状态
0 什么都不做 dp[i][j]=dp[i-1][j-1]
1 插入   dp[i][j-1]+
2 删除 dp[i-1][j]+
3 替换 dp[i-1][j-1]+

 */
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	//初始化
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	//worrd1 变成空需要的步骤
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	//空变成word2需要的步骤
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j]=min(dp[i][j-1]+1,min(dp[i-1][j]+1,dp[i-1][j-1]+1))
			}
		}
	}
	return dp[m][n]
}

func min(i,j int)int  {
	if i>j{
		return j
	}
	return i
}

func main() {

}
