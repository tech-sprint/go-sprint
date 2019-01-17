package p00008

import (
	"strconv"

	. "github.com/weineel/go-sprint/kit"
)

// StringtoInteger ...
func StringtoInteger(str string) int {
	result := 0

	start := 0
	sign := 1
	for start < len(str) && str[start] == ' ' {
		start++
	}

	if start == len(str) {
		return 0
	}

	if str[start] == '-' {
		sign = -1
		start++
	} else if str[start] == '+' {
		sign = 1
		start++
	}

	// len(strconv.Itoa(result)) <= len(strconv.Itoa(INT32MAX)) 判断长度防止 result int溢出
	for ; start < len(str) && str[start] >= '0' && str[start] <= '9' && len(strconv.Itoa(result)) <= len(strconv.Itoa(INT32MAX)); start++ {
		result = result*10 + int(str[start]-'0')
	}

	if sign == 1 && result > INT32MAX {
		return INT32MAX
	}
	if sign == -1 && (result*sign < INT32MIN) {
		return INT32MIN
	}
	return result * sign
}
