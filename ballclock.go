package main

import (
	"fmt"
	"time"
)

const (
	MinMaxBalls     = 4
	FiveMinMaxBalls = 11
	HoursMaxBalls   = 11
)

type clock struct {
	balls            int
	noOfHalfDays     int
	ballFinalQueue   []int
	ballInitialQueue []int

	mins    []int
	fiveMin []int
	hours   []int
}

func (c *clock) init(ballCount int) {
	c.balls = ballCount
	c.noOfHalfDays = 0
	c.assignIntialBalls(ballCount)
}

func (c *clock) assignIntialBalls(ballCount int) {
	for i := 1; i <= ballCount; i++ {
		c.ballInitialQueue = append(c.ballInitialQueue, i)
		c.ballFinalQueue = append(c.ballFinalQueue, i)
	}
}

func (c *clock) startClock() (finalTimeInMills float64) {
	start := time.Now().UnixNano()
	c.addMinutes(c.ballInitialQueue[0])
	c.ballInitialQueue = c.ballInitialQueue[1:]

	for !c.checkIfEnded() {
		//fmt.Println("Values in the queue are :", c.ballInitialQueue, c.mins, c.fiveMin, c.hours)
		c.addMinutes(c.ballInitialQueue[0])
		c.ballInitialQueue = c.ballInitialQueue[1:]
	}
	finalTimeInMills = float64(time.Now().UnixNano()-start) / 1000000
	return
}

func (c *clock) runClockWithMin(minutes int) {
	for i := 1; i <= minutes; i++ {
		c.addMinutes(c.ballInitialQueue[0])
		c.ballInitialQueue = c.ballInitialQueue[1:]
	}
	c.toString()
}

func (c *clock) toString() {
	fmt.Println("{\"Min:\" ", c.mins, "\"FiveMin:\" ", c.fiveMin, "\"Hour:\" ", c.hours, "\"Main:\" ", c.ballInitialQueue, "}")
}
func (c *clock) checkIfEnded() bool {
	if len(c.ballInitialQueue) != len(c.ballFinalQueue) {
		return false
	}
	for ballindex := 0; ballindex < c.balls; ballindex++ {
		if c.ballFinalQueue[ballindex] != c.ballInitialQueue[ballindex] {
			return false
		}
	}
	return true
}
func (c *clock) addMinutes(ball int) {
	//fmt.Println("Adding min ball", ball)
	if len(c.mins) < MinMaxBalls {
		c.mins = append(c.mins, ball)
	} else {
		c.emptyMinutesQueue()
		c.addFiveMinutes(ball)
	}
}
func (c *clock) emptyMinutesQueue() {
	for len(c.mins) > 0 {
		c.ballInitialQueue = append(c.ballInitialQueue, c.mins[len(c.mins)-1])
		c.mins = c.mins[:len(c.mins)-1]
	}
	c.mins = append(c.mins[:])
}
func (c *clock) addFiveMinutes(ball int) {
	//fmt.Println("Adding 5 min Ball", ball)
	if len(c.fiveMin) < FiveMinMaxBalls {
		c.fiveMin = append(c.fiveMin, ball)
	} else {
		c.emptyFiveMinutesQueue()
		c.addHours(ball)
	}
}
func (c *clock) emptyFiveMinutesQueue() {
	for len(c.fiveMin) > 0 {
		c.ballInitialQueue = append(c.ballInitialQueue, c.fiveMin[len(c.fiveMin)-1])
		c.fiveMin = c.fiveMin[:len(c.fiveMin)-1]
	}
}

func (c *clock) addHours(ball int) {
	//fmt.Println("Adding hours Ball", ball)
	if len(c.hours) < HoursMaxBalls {
		c.hours = append(c.hours, ball)
	} else {
		c.emptyHoursQueue()
		c.ballInitialQueue = append(c.ballInitialQueue, ball)
		c.noOfHalfDays++
	}
}

func (c *clock) emptyHoursQueue() {
	for len(c.hours) > 0 {
		c.ballInitialQueue = append(c.ballInitialQueue, c.hours[len(c.hours)-1])
		c.hours = c.hours[:len(c.hours)-1]
	}
}
