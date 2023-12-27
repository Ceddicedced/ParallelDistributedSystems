/*
 * controller.go
 *
 * The controller provides several check routines for a simulated critical section.
 *
 * Copyright (c) 2019-2019 HS Emden/Leer
 * All Rights Reserved.
 *
 * version 2.00 - 30 Oct 2022 - GJV - transform into workspace version
 * version 1.00 - 21 Oct 2019 - GJV - initial version
 *
 * author: Gert Veltink, gert.veltink@hs-emden-leer.de (GJV)
 */

package controller

import (
	"log"
	"math/rand"
	"time"
)

/* global synchronization variable */
var currentProcess = 0

// EnterCriticalSection is called by process 'process' before entering the critical section.
// If another process is already occupying the critical section an error message is generated.
func EnterCriticalSection(process int) {
	if currentProcess != 0 {
		log.Fatalf("Process %d tried to enter the critical section while it was already occupied by: %d\n", process, currentProcess)
	} else {
		currentProcess = process
		log.Printf("entered CS: %d\n", process)
	}
}

// LeaveCriticalSection is called by process 'process' before leaving the critical section.
// If this process is not currently occupying the critical section an error message is generated.
func LeaveCriticalSection(process int) {
	if currentProcess != process {
		if currentProcess == 0 {
			log.Fatalf("Process %d tried to leave the critical section while it was not occupied\n", process)
		} else {
			log.Fatalf("Process %d tried to leave the critical section while it was occupied by: %d\n", process, currentProcess)
		}
	} else {
		currentProcess = 0
		log.Printf("left CS: %d\n", process)
	}
}

// InsideCriticalSection is called by process 'process' to simulate it is performing some tasks
// inside the critical section for msec milliseconds.
func InsideCriticalSection(process int, msec time.Duration) {
	if process != currentProcess {
		log.Fatalf("Process %d tried to work inside the critical section while it has not occupied it\n", process)
	}
	log.Printf("inside CS: %d (%d msecs)\n", process, msec)
	time.Sleep(msec * time.Millisecond)
}

// OutsideCriticalSection is called by process 'process' to simulate it is performing some tasks
// outside the critical section for msec milliseconds.
func OutsideCriticalSection(process int, msec time.Duration) {
	if process == currentProcess {
		log.Fatalf("Process %d tried to work outside the critical section while it occupies the critical section\n", process)
	}
	log.Printf("outside CS: %d (%d msecs)\n", process, msec)
	time.Sleep(msec * time.Millisecond)
}

// ProcessCrashed determines whether the process crashes.
// The parameter probability determines the chance that a process crashes.
// If the process crashes true is returned, false otherwise.
func ProcessCrashed(probability float32) bool {
	return rand.Float32() < probability
}
