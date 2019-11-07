package main


//[l..r]窗口
func lengthOfLongestSubstring(s string) int {
	record:=[128]int{}
	l:=0
	r:=0
	res:=0
	for r<len(s){
		for record[s[r]]!=0{
			record[s[l]]--
			l++
		}
		record[s[r]]++
		res=max(res,r-l+1)
		r++
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
