package main

func isSubsequence(s string, t string) bool {
	n:=len(s)
	if n==0{
		return true
	}
	j:=0
	for i:=0;i<len(t);i++{
		if s[j]==t[i]{
			j++
		}
		if j==len(s){
			return true
		}
	}
	return false
}

func main() {
	
}
