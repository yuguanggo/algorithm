package main

import "fmt"

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 */
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	m := make(map[string]int)
	i := 0
	for j := 0; j < n; j++ {
		if k, ok := m[string(s[j])]; ok {
			if k > i {
				i = k
			}
		}
		if ans < j-i+1 {
			ans = j - i + 1
		}
		m[string(s[j])] = j + 1
	}
	return ans

}

func main() {
	s:="abcabcbb"
	pos:=lengthOfLongestSubstring(s)
	fmt.Println(pos)
}
