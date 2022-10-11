package main

import "fmt"

func main() {
	fmt.Printf("isValid(\"()\"): %v\n", isValid("()"))
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 	1、左括号必须用相同类型的右括号闭合。
// 	2、左括号必须以正确的顺序闭合。
// 	3、每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
    hash := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{'}
    stack := make([]byte, 0)

    if s == "" {
        return true
    }

    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, s[i])
        } else if len(stack) > 0 && stack[len(stack) - 1] == hash[s[i]] {
            stack = stack[:len(stack) - 1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}