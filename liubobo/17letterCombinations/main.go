package main

import (
	"fmt"
)

/**
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func letterCombinations(digits string) []string {
	var res=make([]string,0)
	if len(digits)==0{
		return res
	}
	var lettermap=[]string{
		"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz",
	}
	s:=""
	helper(digits,&res,lettermap,s,0)
	return res
}

func helper(digits string,res *[]string,lettermap []string,s string,index int)  {
	if index==len(digits){
		*res=append(*res,s)
		return
	}
	c:=digits[index]
	letter:=lettermap[c-'0']
	for i:=0;i<len(letter);i++{
		helper(digits,res,lettermap,s+string(byte(letter[i])),index+1)
	}
}

func main() {
	var lettermap=[]string{
		"","","adc","def","ghi","jkl","mno","pqrs","tuv","wxyz",
	}
	fmt.Println(string(byte(lettermap[2][1])))
}
