package currency

import "strconv"

func FormatRupiah(amount float64) string {
	formatted := strconv.FormatFloat(amount, 'f', 0, 64)
	length := len(formatted)
	result := ""
	
	for i := 0; i < length; i++ {
		result += string(formatted[i])
		if (length-i-1)%3 == 0 && i != length-1 {
			result += "."
		}
	}

	return "Rp. " + result
}
