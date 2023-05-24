package main

func bruteforce(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	if sLen > bLen {
		return rSlice
	}

	pos := 0
	for pos+sLen <= bLen {
		found := true
		for sPos := range smallerString {
			if smallerString[sPos] != biggerString[pos+sPos] {
				found = false
				break
			}
		}

		if found {
			rSlice = append(rSlice, pos)
		}
		pos++
	}

	return rSlice
}
