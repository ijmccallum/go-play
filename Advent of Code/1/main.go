package main

import (
	"bufio"
	"os"
	"strconv"
)

func readFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, parseInt(scanner.Text()))
	}
	return lines
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func convertToTriples(input []int) []int {
	var triples []int
	for i := 0; i < len(input)-2; i++ {
		triples = append(triples, input[i]+input[i+1]+input[i+2])
	}
	return triples
}

func main() {
	var lines = readFile("input.txt")
	var triples = convertToTriples(lines)

	var total = 0
	var previous = 0
	for _, triple := range triples {
		var current = triple
		if current > previous {
			total++
		}
		previous = current
	}

	println(total)
}
