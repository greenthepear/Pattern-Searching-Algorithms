package main

import "fmt"

func genHash(str string, base uint64, primeMod uint64) uint64 {
	var r uint64 = 0
	for _, c := range str {
		r = (r*base + uint64(c)) % primeMod
	}
	return r
}

func rollHash(oldHash uint64, oldChar byte, newChar byte, base uint64, primeMod uint64) uint64 {
	//fmt.Printf("[(%d + %d - %d * [(%d mod %d) * %d] mod %d) * %d + %d] mod %d", oldHash, primeMod, int(oldChar), base, primeMod, base, primeMod, base, int(newChar), primeMod)
	fmt.Printf("ASCII of %c = %d\n", newChar, uint64(newChar))
	return ((oldHash+primeMod-uint64(oldChar)*((base%primeMod)*base)%primeMod)*base + uint64(newChar)) % primeMod
}

func rabinkarp(smallerString string, biggerString string) []int {
	sLen, bLen := len(smallerString), len(biggerString)
	rSlice := make([]int, 0)

	const base = 256
	const primeMod = 101

	hash := genHash(smallerString, base, primeMod)
	subStringHash := genHash(biggerString[0:sLen], base, primeMod)

	var power uint64 = 1
	for range smallerString {
		power = (power * base) % primeMod
	}

	i := 0
	for i+sLen < bLen {
		if hash == subStringHash {
			if smallerString == biggerString[i:i+sLen] {
				rSlice = append(rSlice, i)
			}
		}
		//fmt.Printf("\n\nCurrent hash of %s: %d (s hash: %d)", biggerString[i:i+sLen], subStringHash, hash)
		//fmt.Printf("\nRolling %d, removing %c, adding %c", subStringHash, biggerString[i-1], biggerString[i+sLen-1])
		i++
		subStringHash = (subStringHash*base + uint64(biggerString[i])) % primeMod
		if i >= sLen {
			subStringHash -= power * uint64(biggerString[i-sLen]) % primeMod
		}

		//subStringHash = rollHash(subStringHash, biggerString[i], biggerString[i+sLen], base, primeMod) //problem probably here, rewrite, OB1 errors galore
		//fmt.Printf("\tNew hash: %d", subStringHash)
	}

	//Checking last place as to avoid more checks in the loop
	if biggerString[bLen-sLen:] == smallerString {
		rSlice = append(rSlice, bLen-sLen)
	}

	return rSlice
}
