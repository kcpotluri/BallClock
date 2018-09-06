package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func runBallClock(ballCount int) {
	//fmt.Println("Clock is invoked")
	bc := clock{}
	bc.init(ballCount)
	finalTimeInMills := bc.startClock()
	fmt.Printf("%v balls cycle after %v days", bc.balls, bc.noOfHalfDays/2)
	fmt.Printf("\nCompleted in %v milliseconds (%f seconds)", finalTimeInMills, float64(finalTimeInMills/1000))
}
func runBallClockTime(ballCount int, mins int) {
	//fmt.Println("Time is invoked")
	bc := clock{}
	bc.init(ballCount)
	bc.runClockWithMin(mins)
}

func validateInput(ballCount, timeinMin string) (totalBalls int, time int, isvalid bool) {
	if ballCount != "" {
		if balls, err := strconv.Atoi(ballCount); err == nil {
			if balls < 27 || balls > 127 {
				fmt.Println("Ball count should be passed in between 27 to 127")
				isvalid = false
				return
			}
			totalBalls = balls
			isvalid = true
		} else {
			fmt.Println("You should pass only integer for ballCount")
		}
	}
	if timeinMin != "" {
		if timeInMin, err := strconv.Atoi(timeinMin); err == nil && timeInMin > 0 {
			time = timeInMin
		} else {
			fmt.Println("You should pass a valid Integer for mins and should be greater than 0")
			isvalid = false
		}
	}
	return
}
func main() {

	reader := bufio.NewReader(os.Stdin)

	for
	{
		fmt.Println("\n**************************************** PLEASE ENTER YOUR INPUT **********************************")
		fmt.Println("The Ball Clock has two modes to communicate. Below are the ones to choose from")
		fmt.Println("1. For the first mode enter the number of balls to be used in the Clock (should be between 27 to 127) (Example:30). The result will be displayed saying how many days it will take all the balls to return to its original positions along with the time it took to process them. ")
		fmt.Println("2. For the second mode you can enter no of balls(must be 27 to 127) and the number of minutes (must be greater than 0) you want the clock to run(EX 30 325). this will display the  values of the balls in min,fivemins and hours queue in json format")
		fmt.Println("Enter 'exit' to exit the program:")

		text, _ := reader.ReadString('\n')
		input := strings.Split(strings.TrimSpace(text), " ")
		if input[0] == "exit" {
			break
		} else {
			fmt.Println("length is : ", len(input))
			switch len(input) {
			case 1:
				ballcount, _, isvalid := validateInput(input[0], "")
				if isvalid {
					runBallClock(ballcount)
				}
			case 2:
				ballcount, mins, isvalid := validateInput(input[0], input[1])
				if isvalid {
					runBallClockTime(ballcount, mins)
				}

			default:
				fmt.Println("You should enter 1 or 2 arguments or exit.")
			}
		}
	}
}
