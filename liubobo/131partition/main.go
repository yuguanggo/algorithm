package main

/**
131. 分割回文串
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func partition(s string) [][]string {
	var res =make([][]string,0)
	helper(s,0,[]string{},&res)
	return res
}

func helper(s string,pos int,path []string,res *[][]string)  {
	if pos==len(s){
		var p =make([]string,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=1;i<=len(s);i++{
		if pos+i>len(s){
			break;
		}
		sub:=s[pos:pos+i]
		if !isPalindrome(sub){
			continue
		}
		path=append(path,sub)
		helper(s,pos+i,path,res)
		path=path[0:len(path)-1]
	}
}

func isPalindrome(s string) bool {
	n:=len(s)
	l:=0
	r:=n-1
	for l<r{
		if s[l]!=s[r]{
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	
}
