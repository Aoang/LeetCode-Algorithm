package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	str := strings.TrimSpace(s)
	l := len(str)
	if l == 0 {
		return 0
	}
	if !isCode(str[0]) && !isDigit(str[0]) {
		return 0
	}
	f := 1
	if isCode(str[0]) {
		if str[0] == '-' {
			f = -1
		}
		str = str[1:]
		l = len(str)

	}
	res := 0
	for i := 0; i < l; i++ {
		if !isDigit(str[i]) {
			break
		}
		res += int(str[i] - '0')
		if f*res > math.MaxInt32 {
			return math.MaxInt32
		}
		if f*res < math.MinInt32 {
			return math.MinInt32
		}
		res *= 10
	}
	res = f * res / 10
	return res
}

func isDigit(c uint8) bool {
	return '0' <= c && c <= '9'
}

func isCode(c uint8) bool {
	return c == '+' || c == '-'
}
