package main

import "fmt"

/**
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
 */

func isValid(s string) bool {
	bytes:=[]byte(s)
	stack:=make([]byte,0)
	for i:=0;i<len(bytes);i++{
		v:=bytes[i]
		if v=='('||v=='{'||v=='['{
			stack = append(stack,bytes[i])
		}else {
			if len(stack)==0{
				return false
			}
			top:=stack[len(stack)-1]
			var match byte
			if v==')'{
				match='('
			}else if v=='}'{
				match='{'
			}else if v==']'{
				match='['
			}else{
				return false
			}
			if top==match{
				stack=stack[0:len(stack)-1]
			}else{
				return false
			}
		}
	}
	if len(stack)!=0{
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
}
