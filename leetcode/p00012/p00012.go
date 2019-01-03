package p00012

import "strings"

func makeTable(romanTokens string) [10]string {
	rts := strings.Split(romanTokens, "")
	return [10]string{
		"",
		rts[0],
		rts[0] + rts[0],
		rts[0] + rts[0] + rts[0],
		rts[0] + rts[1],
		rts[1],
		rts[1] + rts[0],
		rts[1] + rts[0] + rts[0],
		rts[1] + rts[0] + rts[0] + rts[0],
		rts[0] + rts[2],
	}
}

// IntToRoman ...
func IntToRoman(num int) string {
	result := ""
	romanTokensList := [4]string{"IVX", "XLC", "CDM", "M  "}
	for index := 0; index < len(romanTokensList) && num > 0; index++ {
		result = makeTable(romanTokensList[index])[num%10] + result
		num = num / 10
	}
	return result
}
