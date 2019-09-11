package main

import "fmt"

func isHappy(n int) bool {
	record:=make(map[int]bool)
	return help(n,record)
}

func help(n int,record map[int]bool) bool  {
	if n==1{
		return true
	}
	res:=0
	for n!=0{
		a:=n%10
		res+=a*a
		if a!=0{
			n=n-a
		}else{
			n=n/10
		}
	}
	if _,ok:=record[res];ok{
		return false
	}else{
		record[res]=true
	}
	return help(res,record)
}


func main() {
	fmt.Println(isHappy(19))
}
