package main

import "strconv"

func evalRPN(tokens []string) int {
	n:=len(tokens)
	if n==0{
		return 0
	}
	if n==1{
		num,_:=strconv.Atoi(tokens[0])
		return num
	}
	stack:=make([]int,0)
	for i:=0;i<n;i++{
		if tokens[i]!="+"&&tokens[i]!="-"&&tokens[i]!="*"&&tokens[i]!="/"{
			num,_:=strconv.Atoi(tokens[i])
			stack=append(stack,num)
		}else{
			num1:=stack[len(stack)-1]
			num2:=stack[len(stack)-2]
			stack=stack[:len(stack)-2]
			sum:=0
			switch tokens[i] {
			case "+":
				sum=num1+num2
			case "*":
				sum=num1*num2
			case "-":
				sum=num2-num1
			case "/":
				sum=num2/num1
			}
			stack=append(stack,sum)
		}
	}
	return stack[0]
}


func main() {
	
}
