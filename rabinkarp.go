package main

func genHash(str string, base int64, primeMod int64) int64 {
	var r int64 = 0
	for _, c := range str {
		r = (r*base + int64(c)) % primeMod
	}
	return r
}

func rollHash(oldHash int64, base int64, primeMod int64, subtractedChar byte, addedChar byte, strLen int, pow int64) int64 {
	newHash := (base*(oldHash-int64(subtractedChar)*pow) + int64(addedChar)) % primeMod
	if newHash < 0 {
		newHash += primeMod
	}
	return newHash
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

	//pow := int64(math.Pow(float64(base), float64(sLen-1)))
	var pow int64 = 1
	for i := 0; i < sLen-1; i++ {
		pow = (pow * base) % primeMod
	}

	i := 0
	for i+sLen < bLen {
		if hash == subStringHash && smallerString == biggerString[i:i+sLen] {
			rSlice = append(rSlice, i)
		}

		i++
		subStringHash = rollHash(subStringHash, base, primeMod, biggerString[i-1], biggerString[i-1+sLen], sLen, pow)
	}

	//Checking last place as to avoid more checks in the loop
	if biggerString[bLen-sLen:] == smallerString {
		rSlice = append(rSlice, bLen-sLen)
	}

	return rSlice
}
