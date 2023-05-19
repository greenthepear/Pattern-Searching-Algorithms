package main

func kmpGenTable(word string) []int {
	wLen := len(word)
	rSlice := make([]int, wLen)

	pos := 1
	candidate := 0

	rSlice[0] = -1

	for pos < wLen {
		if word[pos] == word[candidate] {
			rSlice[pos] = rSlice[candidate]
		} else {
			rSlice[pos] = candidate
			for candidate >= 0 && word[pos] != word[candidate] {
				candidate = rSlice[candidate]
			}
		}
		pos++
		candidate++
	}

	rSlice = append(rSlice, candidate)

	return rSlice
}

func kmp(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	if sLen > bLen {
		return rSlice
	}

	table := kmpGenTable(smallerString)

	bi := 0 //biggerString iterator
	si := 0 //smallerString iterator

	for bi < bLen {
		if biggerString[bi] == smallerString[si] {
			bi++
			si++

			if si == sLen {
				rSlice = append(rSlice, bi-si)
				si = table[si]
			}
		} else {
			si = table[si]
			if si < 0 {
				bi++
				si++
			}
		}
	}

	return rSlice
}
