package helpers

import (
	"strconv"
	"strings"
)

// รับตัวเลข, จำนวนทศนิยม และ การบังคับให้แสดงทศนิยมแม้ว่าตัวเลขจะเป็นจำนวนเต็ม
func ConvertFloatformat(num float64, decimal int, forceShowDecimal bool) string {
	numStr := strconv.FormatFloat(num, 'f', decimal, 64)
	intPart, decPart := splitNumber(numStr)

	intPart = formatIntegerPart(intPart)

	if decimal == 0 {
		return intPart
	}

	if forceShowDecimal {
		return intPart + "." + decPart
	} else {
		if len(decPart) == 0 || decPart == strings.Repeat("0", len(decPart)) {
			return intPart
		}
		return intPart + "." + decPart
	}
}

func splitNumber(numStr string) (string, string) {
	parts := strings.Split(numStr, ".")
	intPart := parts[0]
	decPart := ""
	if len(parts) > 1 {
		decPart = parts[1]
	}
	return intPart, decPart
}

func formatIntegerPart(intPart string) string {
	formattedIntPart := ""
	isNegative := false

	if intPart[0] == '-' {
		isNegative = true
		intPart = intPart[1:]
	}

	for i, digit := range reverseString(intPart) {
		if i != 0 && i%3 == 0 {
			formattedIntPart += ","
		}
		formattedIntPart += string(digit)
	}

	if isNegative {
		formattedIntPart += "-"
	}

	return reverseString(formattedIntPart)
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
