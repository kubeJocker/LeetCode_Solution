package binary_number_to_string

import "strings"

func printBin(num float64) string {
	res := strings.Builder{}
	res.WriteString("0.")
	for res.Len() <= 32 && num != 0 {
		num *= 2
		if num == 0 {
			return res.String()
		}
		if num < 1 {
			res.WriteString("0")
		} else {
			res.WriteString("1")
			num--
		}
	}
	return "ERROR"
}
