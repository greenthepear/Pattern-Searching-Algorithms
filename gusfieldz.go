package main

func genZslice(str string) []int {
	strLen := len(str)
	rSlice := make([]int, strLen)

	r, l := 0, 0
	for i := 1; i < strLen; i++ {
		firstFlag := false

		if i > r {
			l, r = i, i
			firstFlag = true
		}

		if !firstFlag {
			if rSlice[i-l] < r-i+1 {
				rSlice[i] = rSlice[i-l]
				continue
			}

			l = i
		}

		for r < strLen && str[r-l] == str[r] {
			r++
		}

		rSlice[i] = r - l
		r--
	}
	return rSlice
}

func gusfieldz(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	if sLen > bLen {
		return rSlice
	}

	conStr := smallerString + "$" + biggerString //P$T
	zSlice := genZslice(conStr)

	for i := 0; i < len(conStr); i++ {
		if zSlice[i] == sLen {
			rSlice = append(rSlice, i-sLen-1)
		}
	}

	return rSlice
}
