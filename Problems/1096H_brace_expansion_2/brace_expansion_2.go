package _096H_brace_expansion_2

import (
	"sort"
	"unicode"
)

/*
花括号展开的表达式可以看作一个由 花括号、逗号 和 小写英文字母 组成的字符串，定义下面几条语法规则：
	如果只给出单一的元素 x，那么表达式表示的字符串就只有 "x"。R(x) = {x}
		例如，表达式 "a" 表示字符串 "a"。
		而表达式 "w" 就表示字符串 "w"。
	当两个或多个表达式并列，以逗号分隔，我们取这些表达式中元素的并集。R({e_1,e_2,...}) = R(e_1) ∪ R(e_2) ∪ ...
		例如，表达式 "{a,b,c}" 表示字符串 "a","b","c"。
		而表达式 "{{a,b},{b,c}}" 也可以表示字符串 "a","b","c"。
	要是两个或多个表达式相接，中间没有隔开时，我们从这些表达式中各取一个元素依次连接形成字符串。R(e_1 + e_2) = {a + b for (a, b) in R(e_1) × R(e_2)}
		例如，表达式 "{a,b}{c,d}" 表示字符串 "ac","ad","bc","bd"。
		表达式之间允许嵌套，单一元素与表达式的连接也是允许的。
例如，表达式 "a{b,c,d}" 表示字符串 "ab","ac","ad"​​​​​​。
例如，表达式 "a{b,c}{d,e}f{g,h}" 可以表示字符串 "abdfg", "abdfh", "abefg", "abefh", "acdfg", "acdfh", "acefg", "acefh"。
*/

/*
采用中缀表达式的思想，利用栈解决问题。用到两个栈，一个符号栈，一个运算对象栈
	遇到‘{’时，当前面是‘}’或者小写英文字母时，需要添加乘号。然后将{入栈
	遇到‘}’时， 不断弹出栈顶元素并运算，直到栈顶为‘{’
	遇到‘,’时，当栈顶元素为‘*’时，计算至栈顶不为‘*’，然后将+入栈
	遇到小写字母，当前面元素为‘}’或小写字母时添加‘*’，然后构建以当前字母为对象的集合放入运算对象栈中

*/

func braceExpansionII(expression string) []string {
	opStack, itemStack := []rune{}, [][]string{}

	operator := func() {
		item1, item2 := itemStack[len(itemStack)-1], itemStack[len(itemStack)-2]
		op := opStack[len(opStack)-1]
		if op == '*' {
			temp := []string{}
			for _, left := range item2 {
				for _, right := range item1 {
					temp = append(temp, left+right)
				}
			}
			itemStack[len(itemStack)-2] = temp
		} else {
			itemStack[len(itemStack)-2] = append(item2, item1...)
		}
		opStack = opStack[:len(opStack)-1]
		itemStack = itemStack[:len(itemStack)-1]
	}

	for i, char := range expression {
		switch char {
		case '{':
			if i > 0 && (expression[i-1] == '}' || unicode.IsLower(rune(expression[i-1]))) {
				opStack = append(opStack, '*')
			}
			opStack = append(opStack, char)
		case '}':
			for opStack[len(opStack)-1] != '{' {
				operator()
			}
			opStack = opStack[:len(opStack)-1]
		case ',':
			for len(opStack) > 0 && opStack[len(opStack)-1] == '*' {
				operator()
			}
			opStack = append(opStack, '+')
		default:
			if i > 0 && (expression[i-1] == '}') || unicode.IsLower(rune(expression[i-1])) {
				opStack = append(opStack, '*')
			}
			itemStack = append(itemStack, []string{expression[i : i+1]})
		}
	}
	for len(opStack) > 0 {
		operator()
	}
	return removeDuplication(itemStack)
}

func removeDuplication(itemStack [][]string) []string {
	hashMap := map[string]bool{}
	res := []string{}
	for _, list := range itemStack {
		for _, str := range list {
			hashMap[str] = true
		}
	}
	for k, _ := range hashMap {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}
