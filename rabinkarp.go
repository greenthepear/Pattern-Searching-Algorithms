package main

import "math"

func genHash(str string, base uint64, primeMod uint64) uint64 {
	var r uint64 = 0
	for _, c := range str {
		r = (r*base + uint64(c)) % primeMod
	}
	return r
}

func rollHash(oldHash uint64, base uint64, primeMod uint64, subtractedChar byte, addedChar byte, strLen int) uint64 {
	oldHash = (oldHash + primeMod - uint64(subtractedChar)*uint64(math.Pow(float64(base), float64(strLen-1)))%primeMod) % primeMod
	return (oldHash*base + uint64(addedChar)) % primeMod
}

func rabinkarp(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	if sLen > bLen {
		return rSlice
	}

	const base = 256
	const primeMod = 101

	hash := genHash(smallerString, base, primeMod)
	subStringHash := genHash(biggerString[0:sLen], base, primeMod)

	i := 0
	for i+sLen < bLen {
		if hash == subStringHash && smallerString == biggerString[i:i+sLen] {
			rSlice = append(rSlice, i)
		}

		i++
		subStringHash = rollHash(subStringHash, base, primeMod, biggerString[i-1], biggerString[i-1+sLen], sLen)
	}

	//Checking last place as to avoid more checks in the loop
	if biggerString[bLen-sLen:] == smallerString {
		rSlice = append(rSlice, bLen-sLen)
	}

	return rSlice
}
