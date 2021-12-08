package main

import "fmt"

var realFish = []int{4, 1, 3, 2, 4, 3, 1, 4, 4, 1, 1, 1, 5, 2, 4, 4, 2, 1, 2, 3, 4, 1, 2, 4, 3, 4, 5, 1, 1, 3, 1, 2, 1, 4, 1, 1, 3, 4, 1, 2, 5, 1, 4, 2, 2, 1, 1, 1, 3, 1, 5, 3, 1, 2, 1, 1, 1, 1, 4, 1, 1, 1, 2, 2, 1, 3, 1, 3, 1, 3, 4, 5, 1, 2, 2, 1, 1, 1, 4, 1, 5, 1, 3, 1, 3, 4, 1, 3, 2, 3, 4, 4, 4, 3, 4, 5, 1, 3, 1, 3, 5, 1, 1, 1, 1, 1, 2, 4, 1, 2, 1, 1, 1, 5, 1, 1, 2, 1, 3, 1, 4, 2, 3, 4, 4, 3, 1, 1, 3, 5, 3, 1, 1, 5, 2, 4, 1, 1, 3, 5, 1, 4, 3, 1, 1, 4, 2, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 4, 5, 1, 2, 5, 3, 1, 1, 3, 1, 1, 1, 1, 5, 1, 2, 5, 1, 1, 1, 1, 1, 1, 3, 5, 1, 3, 2, 1, 1, 1, 1, 1, 1, 1, 4, 5, 1, 1, 3, 1, 5, 1, 1, 1, 1, 3, 3, 1, 1, 1, 4, 4, 1, 1, 4, 1, 2, 1, 4, 4, 1, 1, 3, 4, 3, 5, 4, 1, 1, 4, 1, 3, 1, 1, 5, 5, 1, 2, 1, 2, 1, 2, 3, 1, 1, 3, 1, 1, 2, 1, 1, 3, 4, 3, 1, 1, 3, 3, 5, 1, 2, 1, 4, 1, 1, 2, 1, 3, 1, 1, 1, 1, 1, 1, 1, 4, 5, 5, 1, 1, 1, 4, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 1, 1, 1, 1, 1, 1, 5}

var fish = realFish

var testFish = []int{3, 4, 3, 1, 2}

// var fish = testFish

//each item in array is the number of fish in that state
var fishStateCount [9]int

func newIncFish() {
	fishStateCount[8], fishStateCount[7], fishStateCount[6], fishStateCount[5], fishStateCount[4], fishStateCount[3], fishStateCount[2], fishStateCount[1], fishStateCount[0] = fishStateCount[0], fishStateCount[8], fishStateCount[7]+fishStateCount[0], fishStateCount[6], fishStateCount[5], fishStateCount[4], fishStateCount[3], fishStateCount[2], fishStateCount[1]
}

func incrementFish() {
	var fishToAppend []int
	//print fish
	// fmt.Println(fish)

	for i := 0; i < len(fish); i++ {
		if fish[i] == 0 {
			//append a number to the end of the fish array
			fishToAppend = append(fishToAppend, 8)
			fish[i] = 6
		} else {
			fish[i]--
		}
	}
	//append the fishToAppend to the end of the fish array
	fish = append(fish, fishToAppend...)
}

func increaseByPercent(value float64, percent float64, times int) float64 {
	for i := 0; i < times; i++ {
		value = value * percent
	}
	return value
}

func main() {
	//2 extra days for first cycle: 8
	//6 for the rest
	//how many fish

	// //increment
	// var percentSteps []string
	// var lastCount = float64(len(fish))
	// var thisCount float64
	// //loop 18 times
	// for i := 0; i < 80; i++ {
	// 	//print length of fish
	// 	incrementFish()
	// 	thisCount = float64(len(fish))
	// 	var percentStep = thisCount / lastCount
	// 	percentSteps = append(percentSteps, fmt.Sprintf("%f", percentStep))
	// 	lastCount = thisCount

	// }
	// // fmt.Println(percentSteps)
	// // fmt.Printf("%v", percentSteps)
	// fmt.Println(strings.Join(percentSteps[:], ","))

	//set init fish state
	//loop through fish
	for i := 0; i < len(fish); i++ {
		//increment fish state
		fishStateCount[fish[i]]++
	}
	//loop 80 times
	for i := 0; i < 256; i++ {
		newIncFish()
	}

	//count fish
	var totalFish = 0
	for i := 0; i < len(fishStateCount); i++ {
		totalFish += fishStateCount[i]
	}
	fmt.Println(totalFish)
}
