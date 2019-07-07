package main

import "fmt"

/**
438. 找到字符串中所有字母异位词
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：

字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。
示例 1:

输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 示例 2:

输入:
s: "abab" p: "ab"

输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。

 */
func findAnagrams(s string, p string) []int {
	//s[l..r]为滑动窗口
	l, r := 0, 0
	needs := make(map[byte]int)
	windows := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		needs[p[i]]++
	}
	res := make([]int, 0)
	for r < len(s) {
		windows[s[r]]++
		r++
		if isEqual(windows, needs) {
			if r-l == len(p) {
				res = append(res, l)
			}
		}
		if r-l >= len(p) {
			windows[s[l]]--
			l++
		}
	}
	return res
}

func findAnagrams3(s string, p string) []int {
	hashP := make(map[byte]int)
	hashS := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		hashP[p[i]]++
	}
	res := make([]int, 0)
	m, n := len(s), len(p)
	for i := 0; i < m; i++ {
		if i >= n {
			hashS[s[i-n]]--
		}
		hashS[s[i]]++
		if isEqual(hashS, hashP) {
			res = append(res, i-n+1)
		}
	}
	return res
}
func isEqual(a, b map[byte]int) bool {
	for k, v := range b {
		if a[k] != v {
			return false
		}
	}
	return true
}

func findAnagrams2(s string, p string) []int {
	//s[l..r]为滑动窗口
	l, r := 0, 0
	needs := make(map[string]int)
	windows := make(map[string]int)
	for i := 0; i < len(p); i++ {
		needs[string(p[i])]++
	}
	match := 0
	res := []int{}
	for r < len(s) {
		if _, ok := needs[string(s[r])]; ok {
			windows[string(s[r])]++
			if windows[string(s[r])] == needs[string(s[r])] {
				match++
			}
		}
		r++
		for match == len(needs) {
			if r-l == len(p) {
				res = append(res, l)
			}
			if _, ok := needs[string(s[l])]; ok {
				windows[string(s[l])]--
				if windows[string(s[l])] < needs[string(s[l])] {
					match--
				}
			}
			l++
		}
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams3(s, p))

}
