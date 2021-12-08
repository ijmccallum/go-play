package main

import (
	"bufio"
	"fmt"
	"os"
)

type line struct {
	x1, y1, x2, y2 int
}

type Grid [][]int

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
	println("returning")
	return lines
}

func parseLine(s string) (int, int, int, int) {
	var x1, y1, x2, y2 int
	_, err := fmt.Sscanf(s, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
	if err != nil {
		panic(err)
	}
	//make x1 smaller than x2 and y1 smaller than y2
	if x1 == x2 {
		//for vert
		if y1 > y2 {
			y1, y2 = y2, y1
		}
	} else if y1 == y2 {
		//for horz
		if x1 > x2 {
			x1, x2 = x2, x1
		}
	} else if (x1 + y1) > (x2 + y2) {
		//for decreasing
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	} else if (x2 < x1) && (y2 > y1) && (x1+y1 == x2+y2) {
		//for increasing
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	return x1, y1, x2, y2
}

func convertStringsToLines(strings []string) []line {
	var lines []line
	for _, s := range strings {
		var line line
		line.x1, line.y1, line.x2, line.y2 = parseLine(s)
		lines = append(lines, line)
	}
	return lines
}

func get90s(lines []line) []line {
	var ninty []line
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			ninty = append(ninty, line)
		}
	}
	return ninty
}

func getIncreasing(lines []line) []line {
	var increasing []line
	for _, line := range lines {
		if line.x1 < line.x2 && line.y1 > line.y2 {
			println("increasing got")
			increasing = append(increasing, line)
		}
	}
	return increasing
}
func getDecreasing(lines []line) []line {
	var decreasing []line
	for _, line := range lines {
		if line.x1 < line.x2 && line.y1 < line.y2 {
			decreasing = append(decreasing, line)
		}
	}
	return decreasing
}

//get the largest x1 value from array of lines
func getMaxX(lines []line) int {
	var max int
	for _, line := range lines {
		if line.x1 > max {
			max = line.x1
		}
		if line.x2 > max {
			max = line.x2
		}
	}
	return max
}
func getMaxY(lines []line) int {
	var max int
	for _, line := range lines {
		if line.y1 > max {
			max = line.y1
		}
		if line.y2 > max {
			max = line.y2
		}
	}
	return max
}

func makeGrid(width, height int) Grid {
	var grid Grid
	for i := 0; i < width; i++ {
		var row []int
		for j := 0; j < height; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}
	return grid
}

func part1(lines []line) {
	var ninty = get90s(lines)
	var increasing = getIncreasing(lines)
	var decreasing = getDecreasing(lines)
	println("ALL")
	println(len(lines)) //assuming the right 6 lines
	println("INV")
	println(len(ninty))
	println(len(increasing))
	println(len(decreasing))
	println(len(ninty) + len(increasing) + len(decreasing))
	var grid = makeGrid(getMaxX(lines)+2, getMaxY(lines)+2)

	//tick the points in the grid by line (the ninty!)
	for _, line := range ninty {
		for x := line.x1; x <= line.x2; x++ {
			for y := line.y1; y <= line.y2; y++ {
				grid[x][y]++
			}
		}
	}

	for _, line := range increasing {
		var ycounter = line.y1
		for x := line.x1; x <= line.x2; x++ {
			grid[x][ycounter]++
			ycounter--
		}
	}
	for _, line := range decreasing {
		var ycounter = line.y1
		for x := line.x1; x <= line.x2; x++ {
			println(x, ycounter)
			grid[x][ycounter]++
			ycounter++
		}
	}

	//count the number of points in the grid that are on
	var count int
	for x := range grid {
		// var rowString string
		for y := range grid[x] {
			if grid[x][y] > 1 {
				count++
			}
			// rowString += " " + strconv.Itoa(grid[x][y])
		}
		// println(rowString)
	}
	println(count)
}

func main() {
	lines := convertStringsToLines(readFile("input-real.txt"))
	// println(len(lines))
	// println(lines[0].y1)
	part1(lines)
	// part2(lines)
}

//955571 too high
//952947 too high
