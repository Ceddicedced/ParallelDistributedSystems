package main

import (
	"fmt"
	"sync"
)

// Enumerations for CardinalDirection and Colour. All as ints for minimal memory usage.
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
		<-axisFreeChan    // Wait for the axis to be free.
		c = nextColour(c) // When the axis is free, change the colour.
		fmt.Printf("Traffic Light %v: %v\n", d, c)

		if c == Red {
			redCountChan <- d // Notify that this light is red.
			<-redCountChan    // If this is the first red light, wait for the start of the other red light.
		} else {
			axisFreeChan <- true // Signal that the axis is free. So the other light can change colour.
		}
	}
}

func main() {
	axisFreeChanNorthSouth := make(chan bool) // Channels to signal that an axis is free.
	axisFreeChanEastWest := make(chan bool)   // Channels to signal that an axis is free.
	redCountChan := make(chan CardinalDirection)

	// Counters for red lights.
	redCountNorthSouth := 0
	redCountEastWest := 0

	currentAxis := NorthSouth // Start with North-South axis. Can be changed to East-West.

	var wg sync.WaitGroup
	wg.Add(4)

	// Start TrafficLight goroutines.
	go TrafficLight(North, axisFreeChanNorthSouth, redCountChan, &wg) // Pass in the channels and wait group.
	go TrafficLight(East, axisFreeChanEastWest, redCountChan, &wg)
	go TrafficLight(South, axisFreeChanNorthSouth, redCountChan, &wg)
	go TrafficLight(West, axisFreeChanEastWest, redCountChan, &wg)

	axisFreeChanNorthSouth <- true // Start with North-South axis. Can be changed to East-West.

	for { // Loop forever. This is the main loop. It will never exit.
		select { // Select statement to handle the channels.
		case direction := <-redCountChan: // If a red light is detected.
			if direction == North || direction == South { // If the red light is on the North-South axis.
				redCountNorthSouth++                                      // Increment the counter.
				if redCountNorthSouth == 2 && currentAxis == NorthSouth { // If this is the second red light and the current axis is North-South.
					redCountChan <- direction    // Bring the second red light into the waiting state.
					currentAxis = EastWest       // Switch to East-West axis.
					axisFreeChanEastWest <- true // Send signal to change colour on the East-West axis.
					redCountNorthSouth = 0       // Reset the counter.
				} else { // If this is the first red light
					axisFreeChanNorthSouth <- true // Start the waiting light on the North-South axis to change colour.
					redCountChan <- direction      // Bring the red light into the waiting state. This is implemented to avoid a race condition where a redlight would otherwise maybe be detect the change color signal.
				}
			} else { // Aquivelent to the above, but for the East-West axis.
				redCountEastWest++
				if redCountEastWest == 2 && currentAxis == EastWest {
					redCountChan <- direction
					currentAxis = NorthSouth
					axisFreeChanNorthSouth <- true
					redCountEastWest = 0
				} else {
					axisFreeChanEastWest <- true
					redCountChan <- direction
				}
			}
		}
	}

	close(axisFreeChanNorthSouth)
	close(axisFreeChanEastWest)
	close(redCountChan)
}
