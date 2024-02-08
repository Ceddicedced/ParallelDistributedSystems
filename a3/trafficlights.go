/*
 trafficlights.go

 A specification of a set of four traffic lights at an intersection.

 Copyright (c) 2019-2023 HS Emden-Leer
 All Rights Reserved.

 @version 3.00 - 17 Apr 2023 - GJV - new nomenclature: Phase, Stage, Cycle
                                     CardinalDirection => Phase, Axis => Stage
 @version 2.10 - 08 May 2022 - GJV - optimized match to mCRL2 specification
 @version 2.00 - 07 May 2022 - GJV - adapt to new communication standard
 @version 1.30 - 26 May 2021 - GJV - adapt for module structure
 @version 1.20 - 16 May 2020 - GJV - make code match mCRL2 specification better
 @version 1.10 - 02 Jan 2020 - GJV - streamline code
 @version 1.00 - 17 Dec 2019 - GJV - initial version
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const deltaSleep = 1  // wiggle time to break rigid switching order
const maxRunTime = 50 // time to run the program, in milliseconds

// -----------------------------------------------------
// data section

// Phase defines the different phases in the solution
type Phase string

// representations for cardinal directions
const (
	phaseA Phase = "phaseA"
	phaseB Phase = "phaseB"
	phaseC Phase = "phaseC"
	phaseD Phase = "phaseD"
)

// Stage defines the different stages in the solution
type Stage string

// generate numerical representations for stages automatically
const (
	stage1 Stage = "stage1"
	stage2 Stage = "stage2"
)

// determines the stage based on the phase
var stage = map[Phase]Stage{
	phaseA: stage1,
	phaseB: stage2,
	phaseC: stage1,
	phaseD: stage2,
}

// Colour defines the different traffic light colours
type Colour string

// generate numerical representations for colours automatically
const (
	red    Colour = "red"
	green  Colour = "green"
	yellow Colour = "yellow"
)

// determines the next colour in the order of colours shown by a traffic light
var next = map[Colour]Colour{
	red:    green,
	green:  yellow,
	yellow: red,
}

// -----------------------------------------------------
// process section

// action: show
// shows the identification and the current state (=colour) of a traffic light
func show(p Phase, c Colour) {
	fmt.Printf("TL '%s' shows '%s' \n", p, c)
	//time.Sleep(time.Second) // delay to be able to check output
}

// action: yieldControl
// communication: yieldControl | seizeControl -> handover: ()
func yieldControl() {
	handover <- true // send a sync-message
}

// receives the control over the intersection from another traffic light
// communication: yieldControl | seizeControl -> handover: ()
func seizeControl() {
	<-handover // receive a sync-message
}

// synchronizes symmetrically between two traffic lights
// communication: synchronize | synchronize -> synchronization: (Stage)
// N.B.: the communication is undirected, i.e. both go routines can initiate the synchronization
func synchronize(s Stage) {
	select {
	case synchronization[s] <- true: // send a sync-message
	case <-synchronization[s]: // receive a sync-message
	}
}

// TrafficLight defines the main goroutine modelling one specific traffic light of the intersection
// p: the phase of the traffic light
// s: the starting stage
func TrafficLight(p Phase, s Stage) {
	if stage[p] != s { // one stage starts
		seizeControl() // the other has to wait to seize the control
	}
	TrafficLight_(p, red)
}

// TrafficLight_ is the auxiliary function containing the main loop
// p: the phase of the traffic light
// c: the current colour of the traffic light
func TrafficLight_(p Phase, c Colour) {
	for { // main loop replacing tail recursion
		show(p, c)            // show the current status
		synchronize(stage[p]) // synchronize along the stage
		if c == red {         // showing red?
			yieldControl()        // give up control
			synchronize(stage[p]) // synchronize along the stage
			seizeControl()        // regain control
		}

		// TWEAK: try to break the regular switching pattern
		time.Sleep(time.Duration(rand.Intn(deltaSleep)) * time.Millisecond)

		c = next[c] // replaces tail recursion: TrafficLight_(p, next(c)); p is unchanged
	}
}

// -----------------------------------------------------------------------------------
// communication section

// the channel type used to exchange booleans
type boolChannel chan bool

// communication: yieldControl | seizeControl -> handover: ()
var handover boolChannel = make(chan bool)

// the type to represent a set of boolean channels, one per stage
type StageChannel map[Stage]boolChannel

// communication: synchronize | synchronize -> synchronization: (Stage)
var synchronization = StageChannel{
	stage1: make(chan bool),
	stage2: make(chan bool),
}

// -----------------------------------------------------------------------------------
// start of the application
func main() {

	go TrafficLight(phaseA, stage1)
	go TrafficLight(phaseB, stage1)
	go TrafficLight(phaseC, stage1)
	go TrafficLight(phaseD, stage1)

	time.Sleep(maxRunTime * time.Millisecond)
}
