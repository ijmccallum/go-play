package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//real
var draws = []int{37, 60, 87, 13, 34, 72, 45, 49, 61, 27, 97, 88, 50, 30, 76, 40, 63, 9, 38, 67, 82, 6, 59, 90, 99, 54, 11, 66, 98, 23, 64, 14, 18, 4, 10, 89, 46, 32, 19, 5, 1, 53, 25, 96, 2, 12, 86, 58, 41, 68, 95, 8, 7, 3, 85, 70, 35, 55, 77, 44, 36, 51, 15, 52, 56, 57, 91, 16, 71, 92, 84, 17, 33, 29, 47, 75, 80, 39, 83, 74, 73, 65, 78, 69, 21, 42, 31, 93, 22, 62, 24, 48, 81, 0, 26, 43, 20, 28, 94, 79}

//test
//var draws = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

type slot struct {
	number int
	used   bool
}

type board struct {
	name     string
	score    int
	didBingo bool
	rows     [5][5]slot
}

var boards = make([]board, 100) //test 3 real 100
var weHaveBingo = false

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

func convertStringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		panic("aaaargh")
	}
	return i
}

func fillBoards() {
	var lines = readFile("input-boards-real.txt")
	var boardCounter = 0
	var rowCounter = 0
	var colCounter = 0
	for _, line := range lines {
		//if line is blank
		println(line)
		if line == "" {
			boards[boardCounter].name = "board" + strconv.Itoa(boardCounter)
			println("board " + boards[boardCounter].name + " has score " + strconv.Itoa(boards[boardCounter].score))
			boards[boardCounter].score = 0
			boardCounter++
			rowCounter = 0
		} else {
			var slotsArray = strings.Split(line, " ")
			for _, number := range slotsArray {
				if number != "" {
					var slot = slot{}
					slot.number = convertStringToInt(number)
					slot.used = false
					// println("board:" + strconv.Itoa(boardCounter) + " col:" + strconv.Itoa(colCounter) + " row:" + strconv.Itoa(rowCounter))
					boards[boardCounter].rows[rowCounter][colCounter] = slot
					colCounter++
				}
			}
			colCounter = 0
			rowCounter++
		}
	}
}

func checkBingo(board board) bool {
	var didBingo = false
	for _, row := range board.rows {
		//if every slot in the row is used then bingo
		var usedSlots = 0
		for _, slot := range row {
			if slot.used {
				usedSlots++
			}
		}
		if usedSlots == 5 {
			didBingo = true
		}
	}
	//check columns
	for c := 0; c < 5; c++ {
		var usedSlots = 0
		for r := 0; r < 5; r++ {
			if board.rows[r][c].used {
				usedSlots++
			}
		}
		if usedSlots == 5 {
			didBingo = true
		}
	}
	return didBingo
}

func GetScore(board board, draw int) int {
	//sum all the unmarked numbers
	println("getting score for board " + board.name + " at draw " + strconv.Itoa(draw))
	var sum = 0
	for r, row := range board.rows {
		for c := range row {
			if board.rows[r][c].used == false {
				sum += board.rows[r][c].number
			}
		}
	}
	println("sum is " + strconv.Itoa(sum))
	return sum * draw
}

var boardBingoCountdown = 100

func setBingoCount() {
	for d, draw := range draws {
		if weHaveBingo == false {
			println("d:" + strconv.Itoa(d) + " draw:" + strconv.Itoa(draw))
			for b, board := range boards {
				for r, row := range board.rows {
					for c, col := range row {
						if col.number == draw {
							boards[b].rows[r][c].used = true
						}
					}
				}
				if checkBingo(boards[b]) {
					if boardBingoCountdown == 1 {
						println("final bingo on board " + board.name + " at draw " + strconv.Itoa(draw) + " with score " + strconv.Itoa(GetScore(boards[b], draws[d])))
					}
					if boards[b].didBingo == false {
						boardBingoCountdown--
					}
					boards[b].didBingo = true

				}
			}
		}
	}
}

func main() {
	//get the boards from the input
	fillBoards()
	// println("---")
	// println(boards[0].rows[1][1].number)
	//get the board that bingos first
	// var bingoBoard = getBingoBoard(boards)
	setBingoCount()
	//get the score of the bingoBoard
	// var bingoScore = getBingoScore(bingoBoard)
}

//12924 is wrong too low
