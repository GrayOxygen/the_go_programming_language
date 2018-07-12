package chapters

import (
	"bytes"
	"fmt"
)

func HasPrefix(s, prefix string) bool {
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

func HasPostfix(s, postfix string) bool {
	return len(s) > len(postfix) && s[len(s)-len(postfix):] == postfix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

//给最后三个字符前加上逗号
func Comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	return Comma(str[:n-3]) + "," + str[n-3:]
}

//从右往左每三个字符加一个逗号
func CommaNotRecursion(str string) string {
	//xyzdope  7%3=1
	var buf bytes.Buffer
	runes := []rune(str)
	n := len(runes)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Println("--->", runes[i], string(runes[i]))
		if i <= n%3-1 {
			buf.WriteRune(runes[i])
			continue
		}

		//12  0
		if i == n%3 {
			buf.WriteRune(rune(','))
			buf.WriteRune(runes[i])
			count++
			continue
		}
		fmt.Println(runes[i])
		buf.WriteRune(runes[i])
		count++
		if i > n%3-1 && i < n-1 && count%3 == 0 {
			buf.WriteRune(rune(','))
			continue
		}

	}
	return buf.String()
}
