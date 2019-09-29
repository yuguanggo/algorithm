package main

func isValid(s string) bool {
	var match byte
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if s[i] == ')' {
				match='('
			} else if s[i] == ']' {
				match='['
			} else {
				match='{'
			}
			top := stack[len(stack)-1]
			if top != match {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack)>0{
		return false
	}
	return true
}

func main() {

}
