package main

func findAnagrams(s string, p string) []int {
	//滑动窗口[l,r]
	l,r:=0,0
	res:=make([]int,0)
	needs:=make(map[byte]int)
	if len(p)==0{
		return res
	}
	for i:=0;i<len(p);i++{
		needs[p[i]]++
	}
	windows:=make(map[byte]int)
	n:=len(s)

	for r<n{
		windows[s[r]]++
		if isEqu(windows,needs){
			if r-l+1==len(p){
				res=append(res,l)
			}
		}
		if r-l+1>=len(p){
			windows[s[l]]--
			l++
		}
		r++
	}
	return res
}

func isEqu(windows,needs map[byte]int) bool {
	for k,_:=range needs{
		if windows[k]!=needs[k]{
			return false
		}
	}
	return true
}


func main() {
	
}
