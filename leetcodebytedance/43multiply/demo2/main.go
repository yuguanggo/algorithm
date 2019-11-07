package main

import "strconv"

func multiply(num1 string, num2 string) string {
	n1:=len(num1)
	n2:=len(num2)
	if num1=="0"||num2=="0"{
		return "0"
	}
	if n1==0||n2==0{
		return "0"
	}
	res:=make([]int,n1+n2-1)
	for i:=n1-1;i>=0;i--{
		for j:=n2-1;j>=0;j--{
			res[i+j]+=int(num1[i]-'0')*int(num2[j]-'0')
			if res[i+j]>=10&&i+j!=0{
				res[i+j-1]+=res[i+j]/10
				res[i+j]=res[i+j]%10
			}
		}
	}
	str:=""
	for i:=0;i<len(res);i++{
		str+=strconv.Itoa(res[i])
	}
	return str
}


func main() {
	
}
