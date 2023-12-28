package main

import (
	"fmt"
	"sync"
)

// Enumerations for CardinalDirection and Colour.
type CardinalDirection int
type Colour int
type Axis int

const (
	North CardinalDirection = iota
	East
	South
	West
)

const (
	NorthSouth Axis = iota
	EastWest
)

const (
	Red Colour = iota
	Green
	Yellow
)

// Function to get the next Colour.
func nextColour(c Colour) Colour {
	return (c + 1) % 3
}

// TrafficLight goroutine.
func TrafficLight(d CardinalDirection, axisFreeChan chan bool, redCountChan chan CardinalDirection, wg *sync.WaitGroup) {
	defer wg.Done()
	c := Red // Starting colour, initially set to Red.
	for {
		<-axisFreeChan // Wait for the axis to be free.
		c = nextColour(c)
		fmt.Printf("Traffic Light %v: %v\n", d, c)

		if c == Red {
			redCountChan <- d // Notify that this light is red.
			<-redCountChan    // Wait for the other light to be red.
		} else {
			axisFreeChan <- true // Signal that the axis is free.
		}
	}
}

func main() {
	axisFreeChanNorthSouth := make(chan bool)
	axisFreeChanEastWest := make(chan bool)
	redCountChan := make(chan CardinalDirection)

	// Counters for red lights.
	redCountNorthSouth := 0
	redCountEastWest := 0

	currentAxis := NorthSouth // Start with North-South axis.

	var wg sync.WaitGroup
	wg.Add(4)

	// Start TrafficLight goroutines.
	go TrafficLight(North, axisFreeChanNorthSouth, redCountChan, &wg)
	go TrafficLight(East, axisFreeChanEastWest, redCountChan, &wg)
	go TrafficLight(South, axisFreeChanNorthSouth, redCountChan, &wg)
	go TrafficLight(West, axisFreeChanEastWest, redCountChan, &wg)

	axisFreeChanNorthSouth <- true

	for {
		select {
		case direction := <-redCountChan:
			if direction == North || direction == South {
				redCountNorthSouth++
				if redCountNorthSouth == 2 && currentAxis == NorthSouth {
					redCountChan <- direction
					currentAxis = EastWest
					axisFreeChanEastWest <- true // Switch to East-West axis.
					redCountNorthSouth = 0
				} else {
					axisFreeChanNorthSouth <- true
					redCountChan <- direction
				}
			} else {
				redCountEastWest++
				if redCountEastWest == 2 && currentAxis == EastWest {
					redCountChan <- direction
					currentAxis = NorthSouth
					axisFreeChanNorthSouth <- true // Switch to North-South axis.
					redCountEastWest = 0
				} else {
					axisFreeChanEastWest <- true
					redCountChan <- direction
				}
			}
		}
	}
}
