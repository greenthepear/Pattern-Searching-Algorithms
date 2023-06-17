package main

import (
	"strings"
)

func sunday(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	if sLen > bLen {
		return rSlice
	}

	pos := 0
	for pos+sLen <= bLen {
		if biggerString[pos:pos+sLen] == smallerString {
			rSlice = append(rSlice, pos)
		}

		if pos+sLen < bLen && !strings.ContainsRune(smallerString, rune(biggerString[pos+sLen])) {
			pos += sLen + 1
		} else {
			pos++
		}
	}

	return rSlice
}
