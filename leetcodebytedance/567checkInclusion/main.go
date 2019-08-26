package main

/**
567. 字符串的排列

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
	sm1:=[26]int{}
	sm2:=[26]int{}
	for j := 0; j < n1; j++ {
		sm1[s1[j]-'a']++
		sm2[s2[j]-'a']++

	}
	//维持n1长度的窗口
	for i := 0; i < n2-n1; i++ {
		if match(sm1,sm2){
			return true
		}
		sm2[s2[i+n1]-'a']++
		sm2[s2[i]-'a']--
	}
	return match(sm1,sm2)
}

func match(sm1,sm2 [26]int)bool  {
	for i:=0;i<26;i++{
		if sm1[i]!=sm2[i]{
			return false
		}
	}
	return true
}



func main() {
}

