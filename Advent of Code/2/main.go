package main

import (
	"bufio"
	"os"
	"strconv"
)

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

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func startsWith(s string, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func lastDigit(s string) int {
	return parseInt(s[len(s)-1:])
}

func main() {
	var commands = readFile("input.txt")

	var horz = 0
	var vert = 0
	var aim = 0
	for _, command := range commands {
		switch true {
		case startsWith(command, "forward"):
			horz += lastDigit(command)
			vert += lastDigit(command) * aim
		case startsWith(command, "up"):
			aim -= lastDigit(command)
		case startsWith(command, "down"):
			aim += lastDigit(command)
		}
	}

	println(vert * horz)
}
