package p00012

func makeTable(rts string) [10]string {
	return [10]string{
		"",
		rts[0:1],
		rts[0:1] + rts[0:1],
		rts[0:1] + rts[0:1] + rts[0:1],
		rts[0:1] + rts[1:2],
		rts[1:2],
		rts[1:2] + rts[0:1],
		rts[1:2] + rts[0:1] + rts[0:1],
		rts[1:2] + rts[0:1] + rts[0:1] + rts[0:1],
		rts[0:1] + rts[2:3],
	}
}

// IntToRoman1 ...
// 28ms
func IntToRoman1(num int) string {
	result := ""
	romanTokensList := [4]string{"IVX", "XLC", "CDM", "M  "}
	for index := 0; index < len(romanTokensList) && num > 0; index++ {
		result = makeTable(romanTokensList[index])[num%10] + result
		num = num / 10
	}
	return result
}

// IntToRoman 方法二 直接构造表。
// 24ms
func IntToRoman(num int) string {
	M := [4]string{"", "M", "MM", "MMM"}
	C := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}
