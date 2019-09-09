package main
func lengthOfLongestSubstring(s string) int {
	//滑动窗口[l..r]
	l,r:=0,-1
	n:=len(s)
	res:=0
	record:=[256]int{}
	for l<n{
		if r+1<n&&record[s[r+1]]==0{
			record[s[r+1]]++
			r++
		}else{
			record[s[l]]--
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
