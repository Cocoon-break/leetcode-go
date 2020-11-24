package stack_queue

import (
	"fmt"
	"strings"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

输入：s = "3[a]2[bc]"
输出："aaabcbc"

输入：s = "3[a2[c]]"
输出："accaccacc"
*/

// 思路： 将数字和字符放入一个数据结构中
type DecodeStringStruct struct {
	Times int
	Res   string
}

func decodeString(s string) string {
	stack := make([]DecodeStringStruct, 0)
	var times int
	var res string

	for _, char := range s {
		if char >= '0' && char <= '9' {
			times = times*10 + int(char-'0')
		} else if char == '[' {
			st := DecodeStringStruct{Times: times, Res: res}
			stack = append(stack, st)
			res = ""
			times = 0
		} else if char == ']' {
			st := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			fmt.Printf("res= %+v times=%+v \n", res, st.Times)
			res = st.Res + strings.Repeat(res, st.Times)
		} else {
			res = res + string(char)
		}
	}
	return res
}
