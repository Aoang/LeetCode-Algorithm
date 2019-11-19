package main

import "bytes"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	str := changeStr(s)
	lenStr := len(str)
	prGroup := make([]int, lenStr)
	tempCentre, tempRight, prMax, prMaxIndex := 0, 0, 0, 1
	for i := 1; i < len(str)-1; i++ {

		if i <= tempRight {
			prGroup[i] = minInt(tempRight-i, prGroup[2*tempCentre-i])
		} else {
			prGroup[i] = 0
		}
		for str[i-prGroup[i]-1] == str[i+prGroup[i]+1] {
			prGroup[i]++
		}
		if tempRight < prGroup[i]+i {
			tempCentre = i
			tempRight = i + prGroup[i]
		}
		if prMax < prGroup[i] {
			prMax = prGroup[i]
			prMaxIndex = i
		}
	}
	return s[(prMaxIndex-1-prMax)/2 : (prMaxIndex-1+prMax)/2]

}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func changeStr(source string) string {
	b := bytes.Buffer{}
	b.WriteString("$#")
	for i := 0; i < len(source); i++ {
		b.WriteString(string(source[i]))
		b.WriteString("#")
	}
	b.WriteString("&")
	return b.String()
}
