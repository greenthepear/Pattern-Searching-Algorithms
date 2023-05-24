package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func generateSurroundingText(smallerString string, biggerString string, position int) string {
	rString := ""
	surroundingLetters := 30

	biggerString = strings.ReplaceAll(biggerString, "\n", " ")

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
		indexes = bruteforce(smallerString, biggerString)
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

func testAll(smallerString string, biggerString string) {
	fmt.Printf("\nBruteforce:\n")
	patternAlgorithmRead(smallerString, biggerString, 0)
	fmt.Printf("\nSunday:\n")
	patternAlgorithmRead(smallerString, biggerString, 1)
	fmt.Printf("\nKMP:\n")
	patternAlgorithmRead(smallerString, biggerString, 2)
	fmt.Printf("\nRabin-Karp:\n")
	patternAlgorithmRead(smallerString, biggerString, 3)
	fmt.Printf("\nFSM:\n")
	patternAlgorithmRead(smallerString, biggerString, 4)
	fmt.Printf("\nGusfield Z:\n")
	patternAlgorithmRead(smallerString, biggerString, 5)
}

func readLinesFromFile(filepath string, numLines int) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := ""
	lineCount := 0

	for scanner.Scan() {
		lines += scanner.Text() + "\n"
		lineCount++

		if lineCount >= numLines {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return lines, nil
}

func testAllFileLines(numLines int, filepath string, smallerString string) {
	biggerString, err := readLinesFromFile(filepath, numLines)
	if err != nil {
		fmt.Printf("Error:%s", err)
		return
	}

	testAll(smallerString, biggerString)
}

func benchmark(maxLines int, times int, smallerString string, filepath string) {
	timeElapsed := make([]time.Duration, 6)
	timeElapsedStart := make([]time.Time, 6)
	indexes := make([][]int, 6)

	fmt.Printf("Lines\tBf\tSd\tKMP\tRK\tFSM\tZ\n")
	for ilines := 12; ilines < maxLines; ilines += 3 {
		biggerString, err := readLinesFromFile(filepath, ilines)
		if err != nil {
			fmt.Printf("Error:%s", err)
			return
		}

		timeElapsedTotal := make([]time.Duration, 6)
		for iavg := 0; iavg < times; iavg++ {
			//Bruteforce
			timeElapsedStart[0] = time.Now()
			indexes[0] = bruteforce(smallerString, biggerString)
			timeElapsed[0] = time.Since(timeElapsedStart[0])
			timeElapsedTotal[0] += timeElapsed[0]

			//Sunday
			timeElapsedStart[1] = time.Now()
			indexes[1] = sunday(smallerString, biggerString)
			timeElapsed[1] = time.Since(timeElapsedStart[1])
			timeElapsedTotal[1] += timeElapsed[1]

			//KMP
			timeElapsedStart[2] = time.Now()
			indexes[2] = kmp(smallerString, biggerString)
			timeElapsed[2] = time.Since(timeElapsedStart[2])
			timeElapsedTotal[2] += timeElapsed[2]

			//R-K
			timeElapsedStart[3] = time.Now()
			indexes[3] = rabinkarp(smallerString, biggerString)
			timeElapsed[3] = time.Since(timeElapsedStart[3])
			timeElapsedTotal[3] += timeElapsed[3]

			//FSM
			timeElapsedStart[4] = time.Now()
			indexes[4] = rabinkarp(smallerString, biggerString)
			timeElapsed[4] = time.Since(timeElapsedStart[4])
			timeElapsedTotal[4] += timeElapsed[4]

			//Z
			timeElapsedStart[5] = time.Now()
			indexes[5] = rabinkarp(smallerString, biggerString)
			timeElapsed[5] = time.Since(timeElapsedStart[3])
			timeElapsedTotal[5] += timeElapsed[3]
		}

		fmt.Printf("\n%d\t", ilines)
		for i := range timeElapsedTotal {
			timeElapsedTotal[i] /= time.Duration(times)
			fmt.Printf("%d\t", timeElapsedTotal[i])
		}
	}

}

func main() {
	//text := "The King's Indian Defence (or KID) is a common chess opening. It is defined by the following moves: 1. d4 Nf6 2. c4 g6. Black intends to follow up with 3...Bg7 and 4...d6 (the Grünfeld Defence arises when Black plays 3...d5 instead and is considered a separate opening). White's major third move options are 3.Nc3, 3.Nf3 or 3.g3, with both the King's Indian and Grünfeld playable against these moves. The Encyclopaedia of Chess Openings classifies the King's Indian Defence under the codes E60 through E99."
	//baseString := "abracadabra"
	//pattern := "he"
	filepath := "hhgttg.txt"
	pattern := "th"
	//text := "ththththththth"

	//testAll(pattern, text)

	//testAll(pattern, text)

	testAllFileLines(50, filepath, pattern)

	//benchmark(500, 50, pattern, filepath)
}
