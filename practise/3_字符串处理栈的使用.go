package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：

	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	每个右括号都有一个对应的相同类型的左括号。

示例 1：输入：s = "()"输出：true
示例 2：输入：s = "()[]{}"输出：true
示例 3：输入：s = "(]"输出：false
示例 4：输入：s = "([])"输出：true
示例 5：输入：s = "([)]"输出：false

提示：

	1 <= s.length <= 10的4次方
	s 仅由括号 '()[]{}' 组成
*/
func isValid(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	matching := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	for _, ch := range str {
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else if matched, ok := matching[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != matched {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}
func main() {
	s := "(){}{{{}}}"
	fmt.Println(isValid(s))
}
