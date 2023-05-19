package main

import (
	"fmt"
)

func bruteforce(smallerString string, biggerString string, printInfo bool) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)
	if sLen > bLen {
		if printInfo {
			fmt.Printf("Wrong sizes!\n")
		}
		return rSlice
	}
	foundAny := false

	for bIndex := range biggerString {
		if bIndex+sLen > bLen {
			break
		}

		patternFound := true
		for sIndex := range smallerString {
			//fmt.Printf("Comparing %c and %c...\n", sChar, biggerString[bIndex+sIndex])
			if smallerString[sIndex] != biggerString[bIndex+sIndex] {
				patternFound = false
				break
			}
		}

		if patternFound {
			if printInfo {
				fmt.Printf("Found pattern '%s' at position %d.\n", smallerString, bIndex)
			}
			foundAny = true
			rSlice = append(rSlice, bIndex)
		}
	}

	if !foundAny && printInfo {
		fmt.Printf("Pattern '%s' not found anywhere in '%s'.\n", smallerString, biggerString)
	}

	return rSlice
}
