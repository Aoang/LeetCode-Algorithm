package main

func lengthOfLongestSubstring(s string) int {
	loc := [256]int{}
	max, left := 0, 0

	for i := 0; i < len(s); i++ {
		if loc[s[i]] >= left+1 {
			left = loc[s[i]]
		} else if i+1-left > max {
			max = i + 1 - left
		}
		loc[s[i]] = i + 1
	}

	return max
}
