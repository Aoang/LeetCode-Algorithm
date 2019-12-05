package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	arr := make([]string, numRows)
	idx := 0
	l := len(s)

	for idx < l {
		i := 0

		for ; i < numRows; i++ {
			arr[i] += string(s[idx])
			idx++
			if idx == l {
				goto sum
			}
		}
		for i -= 2; i > 0; i-- {
			arr[i] += string(s[idx])
			idx++
			if idx == l {
				goto sum
			}
		}
	}

sum:
	return strings.Join(arr, "")
}
