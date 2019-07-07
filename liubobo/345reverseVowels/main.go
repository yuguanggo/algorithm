package main

/**
反转字符串中的元音字母
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func reverseVowels(s string) string {
	b:=[]byte(s)
	n:=len(s)
	i:=0
	j:=n-1
	for i<j{
		if b[i]!='a'&&b[i]!='e'&&b[i]!='i'&&b[i]!='o'&&b[i]!='u'&&b[i]!='A'&&b[i]!='E'&&b[i]!='I'&&b[i]!='O'&&b[i]!='U'{
			i++
			continue
		}
		if b[j]!='a'&&b[j]!='e'&&b[j]!='i'&&b[j]!='o'&&b[j]!='u'&&b[j]!='A'&&b[j]!='E'&&b[j]!='I'&&b[j]!='O'&&b[j]!='U'{
			j--
			continue
		}
		b[i],b[j]=b[j],b[i]
		i++
		j--
	}
	return string(b)
}

func main() {
	
}
