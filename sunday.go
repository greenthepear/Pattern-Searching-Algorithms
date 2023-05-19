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
	for pos+sLen < bLen {
		if biggerString[pos:pos+sLen] == smallerString {
			rSlice = append(rSlice, pos)
		}

		if !strings.ContainsRune(smallerString, rune(biggerString[pos+sLen])) {
			pos += sLen
		} else {
			pos++
		}
	}

	return rSlice
}
