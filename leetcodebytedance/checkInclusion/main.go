package main

import (
	"strings"
	"fmt"
)

/**
字符串的排列

给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False


注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间
 */
func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1>n2 {
		return false
	}
	m := make(map[string]int)
	//存储s1每个字符的个数
	sm1 := make(map[uint8]int)
	for j := 0; j < n1; j++ {
		if v, ok := sm1[s1[j]-'a']; ok {
			sm1[s1[j]-'a'] = v + 1
		} else {
			sm1[s1[j]-'a'] = 1
		}

	}
	count:=0
	for i := 0; i < n2; i++ {
		if strings.Index(s1, ss2[i]) != -1 {
			if v, ok := m[ss2[i]]; ok {
				//判断同一个字符是否超过sm1里的个数
				if v < sm1[ss2[i]] {
					m[ss2[i]] = v + 1
				} else {
					count=0
					m = make(map[string]int)
				}
			} else {
				m[ss2[i]] = 1
			}
			count++
			if count==n1{
				return true
			}
		} else {
			m = make(map[string]int)
		}
	}
	return false
}

func main() {
	s1 := "a"
	//s2 := "ab"
	fmt.Println((s1[0]-'a'))
	//fmt.Println(checkInclusion(s1, s2))
}

