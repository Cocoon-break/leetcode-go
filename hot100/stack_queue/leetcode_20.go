package stack_queue

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

输入: "()"
输出: true

输入: "(]"
输出: false

输入: "()[]{}"
输出: true
*/
func isValid(s string) bool {
	maps := map[uint8]uint8{
		')': '(',
		']': '[',
		'}': '{',
	}
	var chars []uint8
	for i := 0; i < len(s); i++ {
		char := s[i] //得到uint8
		if char == '(' || char == '[' || char == '{' {
			chars = append(chars, char)
		} else {
			//第一个元素不是)]}中的一种
			if len(chars) == 0 {
				return false
			}
			//存在不相等情况
			if maps[char] != chars[len(chars)-1] {
				return false
			}
			//存在相等则移除栈顶元素
			chars = chars[:len(chars)-1]
		}
	}
	return len(chars) == 0
}
