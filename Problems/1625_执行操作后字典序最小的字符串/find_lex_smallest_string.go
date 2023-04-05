package _625_执行操作后字典序最小的字符串

/*
 */
func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	res := s
	s = s + s
	g := gcd(b, n)
	add := func(str *string, start int) {
		minVal, times := 10, 0
		for i := 0; i < 10; i++ {
			added := (int((*str)[start]-'0') + i*a) % 10
			if added < minVal {
				minVal = added
				times = i
			}
		}
		temp := []byte(*str)
		for i := start; i < n; i += 2 {
			temp[i] = byte('0' + (int((*str)[i]-'0')+times*a)%10)
		}
		*str = string(temp)
	}

	for i := 0; i < n; i += g {
		temp := s[i : i+n]
		add(&temp, 1)
		if b%2 == 1 {
			add(&temp, 0)
		}
		if res > temp {
			res = temp
		}
	}
	return res
}

func gcd(num1, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}
