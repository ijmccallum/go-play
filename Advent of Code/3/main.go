package main

import (
	"bufio"
	"os"
	"strconv"
)

//read a given file and return the content as an array of strings for each line
func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

//get character from given position in string
func getChar(str string, pos int) string {
	return string(str[pos])
}

//convert binary to decimal
func binToDec(binary string) int {
	var decimal int
	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			decimal += (1 << uint(i))
		}
	}
	return decimal
}

//convert array of strings to string
func arrayToString(arr []string) string {
	var str string
	for _, v := range arr {
		str += v
	}
	return str
}

func getMostCommonBits(bitLength int, lines []string) {
	var commonBits = make([]string, bitLength)
	for i := 0; i < bitLength; i++ {
		var count0 = 0
		var count1 = 0
		for _, line := range lines {
			if getChar(line, i) == "0" {
				count0++
			} else {
				count1++
			}
		}
		if count0 > count1 {
			commonBits[i] = "0"
		} else {
			commonBits[i] = "1"
		}
	}
}

func getOxyRating(bitLength int, lines []string) int {
	for i := 0; i < bitLength; i++ {
		//get the most common bit in position i
		var count0 = 0
		var count1 = 0
		var mostCommonBit = ""
		for _, line := range lines {
			if getChar(line, i) == "0" {
				count0++
			} else {
				count1++
			}
		}
		if count0 > count1 {
			mostCommonBit = "0"
		} else if count1 > count0 {
			mostCommonBit = "1"
		} else {
			mostCommonBit = "1"
		}
		//filter out all lines that don't have that bit in position i
		var filteredLines = []string{}
		for _, line := range lines {
			if getChar(line, i) == mostCommonBit {
				filteredLines = append(filteredLines, line)
			}
		}
		//if there's only one line left, return that line
		if len(filteredLines) == 1 {
			var oxy, err = strconv.ParseInt(filteredLines[0], 2, 32)
			if err != nil {
				panic(err)
			}
			return int(oxy)
		}
		lines = filteredLines
	}
	return 0
}

func getCo2Rating(bitLength int, lines []string) int {
	for i := 0; i < bitLength; i++ {
		//get the most common bit in position i
		var count0 = 0
		var count1 = 0
		var leastCommonBit = ""
		for _, line := range lines {
			if getChar(line, i) == "0" {
				count0++
			} else {
				count1++
			}
		}
		if count0 > count1 {
			leastCommonBit = "1"
		} else if count1 > count0 {
			leastCommonBit = "0"
		} else {
			leastCommonBit = "0"
		}
		//filter out all lines that don't have that bit in position i
		var filteredLines = []string{}
		for _, line := range lines {
			if getChar(line, i) == leastCommonBit {
				filteredLines = append(filteredLines, line)
			}
		}
		//if there's only one line left, return that line
		if len(filteredLines) == 1 {
			var oxy, err = strconv.ParseInt(filteredLines[0], 2, 32)
			if err != nil {
				panic(err)
			}
			return int(oxy)
		}
		lines = filteredLines
	}
	return 0
}

func main() {
	var lines = readFile("input.txt")
	var bitLength = len(lines[0])

	// var commonBitString = getMostCommonBits(bitLength, lines)

	var oxy = getOxyRating(bitLength, lines)
	var co2 = getCo2Rating(bitLength, lines)
	var lifeSup = oxy * co2

	println(lifeSup)
}
