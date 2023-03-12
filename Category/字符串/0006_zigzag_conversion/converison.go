package _006_zigzag_conversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	zigList := make([]strings.Builder, numRows)
	row, flag := 0, -1
	for _, c := range s {
		zigList[row].WriteString(string(c))
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		row += flag
	}
	res := strings.Builder{}
	for _, str := range zigList {
		res.WriteString(str.String())
	}
	return res.String()
}
