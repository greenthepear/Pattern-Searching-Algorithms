package main

func genHash(str string, base uint64, primeMod uint64) uint64 {
	var r uint64 = 0
	for _, c := range str {
		r = (r*base + uint64(c)) % primeMod
	}
	return r
}

func power(base, exp, mod uint64) uint64 {
	result := uint64(1)
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

func rollHash(oldHash uint64, base uint64, primeMod uint64, subtractedChar byte, addedChar byte, strLen int) uint64 {
	// Subtract the contribution of the removed character from the original hash
	oldHash = (oldHash - uint64(subtractedChar)*power(base, uint64(strLen-1), primeMod)) % primeMod

	// Add the contribution of the added character to the updated hash
	newHash := (oldHash*base + uint64(addedChar)) % primeMod
	return newHash
}

// Helper function to calculate base^exp % mod efficiently

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

		i++
		//oldSubstringHash := subStringHash
		subStringHash = genHash(biggerString[i:i+sLen], base, primeMod)
		//fmt.Printf("\nOldS\tNewS\tOldH\tNewH\tDiff\n%s\t%s\t%d\t%d\t%d\n",
		//	biggerString[i-1:i-1+sLen], biggerString[i:i+sLen],
		//	oldSubstringHash, subStringHash, oldSubstringHash-subStringHash)
	}

	//Checking last place as to avoid more checks in the loop
	if biggerString[bLen-sLen:] == smallerString {
		rSlice = append(rSlice, bLen-sLen)
	}

	return rSlice
}
