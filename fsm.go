package main

func largestSuffixPrefix(str string, state int, symbol int) int {
	for istate := state; istate > 0; istate-- {
		if int(str[istate-1]) == symbol {
			for i := 0; i < istate-1; i++ {
				if str[i] != str[state-istate+1+i] {
					break
				}

				if i == istate-1 {
					return istate
				}
			}
		}
	}

	return 0
}

func genFSMmatrix(str string) [][]int {
	strLen := len(str)

	r := make([][]int, strLen+1)
	for i := 0; i < strLen+1; i++ {
		r[i] = make([]int, 256)
	}

	for state := 0; state <= strLen; state++ {
		for symbol := 0; symbol < 256; symbol++ {
			if state < strLen && symbol == int(str[state]) {
				r[state][symbol] = state + 1
			} else {
				r[state][symbol] = largestSuffixPrefix(str, state, symbol)
			}
		}
	}

	return r
}

func fsm(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	if sLen > bLen {
		return rSlice
	}

	matrix := genFSMmatrix(smallerString)

	i, state := 0, 0
	for i < bLen {
		state = matrix[state][biggerString[i]]
		if state == sLen {
			rSlice = append(rSlice, i-sLen+1)
		}
		i++
	}

	return rSlice
}
