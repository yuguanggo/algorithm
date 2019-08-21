package main

/**
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？


 */
func isAnagram(s string, t string) bool {
	hashS := [128]int{}
	if len(s)==0&&len(t)==0{
		return true
	}
	if len(s)!=len(t){
		return false
	}
	cs := 0
	ct :=0
	for _, v := range s {
		if hashS[v] == 0 {
			cs++
		}
		hashS[v]++
	}
	for _,v:=range t{
		if hashS[v]>0{
			hashS[v]--
			if hashS[v]==0{
				ct++
			}
		}
	}
	if cs==ct&&cs!=0&&ct!=0{
		return true
	}
	return false
}

func main() {

}
