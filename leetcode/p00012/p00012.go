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

// IntToRoman ...
// 28ms
func IntToRoman(num int) string {
	result := ""
	romanTokensList := [4]string{"IVX", "XLC", "CDM", "M  "}
	for index := 0; index < len(romanTokensList) && num > 0; index++ {
		result = makeTable(romanTokensList[index])[num%10] + result
		num = num / 10
	}
	return result
}
