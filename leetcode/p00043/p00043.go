package p00043

import "strconv"

// MultiplyStrings ...
func MultiplyStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	a1, a0 := num1[0:m-m/2], num1[m-m/2:]
	b1, b0 := num2[0:n-n/2], num2[n-n/2:]

	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if m == 1 && n == 1 {
		return strconv.Itoa(int(a1[0]-'0') * int(b1[0]-'0'))
	}
	if m == 1 && n > 1 {
		return AddStrings(appendZero(MultiplyStrings(a1, b1), n/2), MultiplyStrings(a1, b0))
	}
	if m > 1 && n == 1 {
		return AddStrings(appendZero(MultiplyStrings(a1, b1), m/2), MultiplyStrings(a0, b1))
	}
	return AddStrings(AddStrings(appendZero(MultiplyStrings(a1, b1), m/2+n/2), appendZero(MultiplyStrings(a1, b0), m/2)), AddStrings(appendZero(MultiplyStrings(a0, b1), n/2), MultiplyStrings(a0, b0)))
}

// AddStrings 大数加法
func AddStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	carry := 0
	result := ""
	for l1 > 0 || l2 > 0 || carry > 0 {
		n1 := 0
		if l1 > 0 {
			l1--
			n1 = int(num1[l1] - '0')
		}
		n2 := 0
		if l2 > 0 {
			l2--
			n2 = int(num2[l2] - '0')
		}
		n := n1 + n2 + carry
		carry = n / 10
		result = string(n%10+'0') + result
	}
	return result
}

func appendZero(num string, length int) string {
	s := make([]byte, length)
	for i := range s {
		s[i] = '0'
	}
	return num + string(s)
}
