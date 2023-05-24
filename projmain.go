package main

import "fmt"

func generateSurroundingText(smallerString string, biggerString string, position int) string {
	rString := ""
	surroundingLetters := 30
	if position-surroundingLetters < 0 {
		rString += biggerString[:position]
	} else {
		rString += "..." + biggerString[position-surroundingLetters:position]
	}

	rString += "‹" + smallerString + "›"

	if position+surroundingLetters+len(smallerString) > len(biggerString) {
		rString += biggerString[position+len(smallerString):]
	} else {
		rString += biggerString[position+len(smallerString):position+len(smallerString)+surroundingLetters] + "..."
	}

	return rString
}

func patternAlgorithmRead(smallerString string, biggerString string, algorithm int) {
	var indexes []int
	switch algorithm {
	case 0:
		indexes = bruteforce(smallerString, biggerString, false)
	case 1:
		indexes = sunday(smallerString, biggerString)
	case 2:
		indexes = kmp(smallerString, biggerString)
	case 3:
		indexes = rabinkarp(smallerString, biggerString)
	case 4:
		indexes = fsm(smallerString, biggerString)
	case 5:
		indexes = gusfieldz(smallerString, biggerString)
	default:
		fmt.Printf("Bad input!\n")
		return
	}

	if len(indexes) == 0 {
		fmt.Printf("Pattern ‹%s› not found anywhere.\n", smallerString)
		return
	}

	fmt.Printf("Pattern ‹%s› found at %d position(s): %v\n", smallerString, len(indexes), indexes)

	for i, pos := range indexes {
		fmt.Printf("%d.\t%d:\t%s\n", i+1, pos, generateSurroundingText(smallerString, biggerString, pos))
	}
}

func main() {
	baseString := "The King's Indian Defence (or KID) is a common chess opening. It is defined by the following moves: 1. d4 Nf6 2. c4 g6. Black intends to follow up with 3...Bg7 and 4...d6 (the Grünfeld Defence arises when Black plays 3...d5 instead and is considered a separate opening). White's major third move options are 3.Nc3, 3.Nf3 or 3.g3, with both the King's Indian and Grünfeld playable against these moves. The Encyclopaedia of Chess Openings classifies the King's Indian Defence under the codes E60 through E99."
	//baseString := "abracadabra"
	toFindString := "he"
	//bruteforce(toFindString, baseString, true)
	fmt.Printf("\nBruteforce:\n")
	patternAlgorithmRead(toFindString, baseString, 0)
	fmt.Printf("\nSunday:\n")
	patternAlgorithmRead(toFindString, baseString, 1)
	fmt.Printf("\nKMP:\n")
	patternAlgorithmRead(toFindString, baseString, 2)
	fmt.Printf("\nRabin-Karp:\n")
	patternAlgorithmRead(toFindString, baseString, 3)
	fmt.Printf("\nFSM:\n")
	patternAlgorithmRead(toFindString, baseString, 4)
	fmt.Printf("\nGusfield Z:\n")
	patternAlgorithmRead(toFindString, baseString, 5)

	//zTestString := "aabaacd"
	//fmt.Printf("\nZ test for '%s': %v", zTestString, genZslice(zTestString))
}
