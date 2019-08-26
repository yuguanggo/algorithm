package main

func lengthOfLongestSubstring(s string) int {
	n:=len(s)
	//s[l..r]作为滑动窗口
	l:=0
	r:=-1
	freq:=[256]int{}
	res:=0
	for l<n{
		if r+1<n&&freq[s[r+1]]==0{
			freq[s[r+1]]++
			r++
		}else{
			freq[s[l]]--
			l++
		}
		res=max(res,r-l+1)
	}
	return res
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}
func main() {
	
}
