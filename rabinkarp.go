package main

import "fmt"

func genHash(str string, base int, primeMod int) int {
	r := 0
	for _, c := range str {
		r = (r*base + int(c)) % primeMod
	}
	return r
}

func rollHash(oldHash int, oldChar byte, newChar byte, base int, primeMod int) int {
	//fmt.Printf("[(%d + %d - %d * [(%d mod %d) * %d] mod %d) * %d + %d] mod %d", oldHash, primeMod, int(oldChar), base, primeMod, base, primeMod, base, int(newChar), primeMod)
	return ((oldHash+primeMod-int(oldChar)*((base%primeMod)*base)%primeMod)*base + int(newChar)) % primeMod
}

func rabinkarp(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	hash := genHash(smallerString, 256, 101)
	subStringHash := genHash(biggerString[0:sLen], 256, 101)

	i := 1
	for i+sLen < bLen {
		fmt.Printf("\nCurrent hash of %s: %d\n", biggerString[i-1:i-1+sLen], subStringHash)
		subStringHash = rollHash(subStringHash, biggerString[i-1], biggerString[i+sLen], 256, 101) //problem probably here, rewrite, OB1 errors galore
		if hash == subStringHash {
			if smallerString == biggerString[i:i+sLen] {
				rSlice = append(rSlice, i)
			}
		}
		i++
	}

	return rSlice
}
