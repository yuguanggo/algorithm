package main

import (
	"fmt"
	"strings"
)

/**
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func isPalindrome(s string) bool {
	s =strings.ToLower(s)
	n:=len(s)
	i:=0
	j:=n-1
	for i<j{
		if !(s[i]>='a'&&s[i]<='z'||s[i]>='0'&&s[i]<='9'){
			i++
			continue
		}
		if !(s[j]>='a'&&s[j]<='z'||s[j]>='0'&&s[j]<='9'){
			j--
			continue
		}
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
}



func main() {
	s:="azAZ09"
	b:=[]byte(s)

	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
	fmt.Println(b[3])
	fmt.Println(b[4])
	fmt.Println(b[5])
}
