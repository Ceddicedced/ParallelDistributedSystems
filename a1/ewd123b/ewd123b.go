/*
 * ewd123b.go
 *
 * A program to represent the fourth mutex strategy, as described in EWD123.
 *
 * Copyright (c) 2023 Cedric Busacker
 * Original Copyright (c) 2019-2022 HS Emden/Leer
 * All Rights Reserved.
 *
 * version 2.01 - 03 Nov 2023 - CB - adding required functions
 * version 2.00 - 30 Oct 2022 - GJV - transform into workspace version
 * version 1.00 - 21 Oct 2019 - GJV - initial version
 *
 * author: Cedric Busacker, cedric@busacker.dev (CB)
 * original author: Gert Veltink, gert.veltink@hs-emden-leer.de (GJV)
 */

package ewd123b

import (
	"log"

	"busacker.dev/a1/controller"
)

// global synchronization variables
var c1, c2 = true, true

// Start starts the execution of EWD123b.
func Start() {
	go process1()
	go process2()
}

// process1 simulates the behaviour of the first process
func process1() {
	for {

		for c2 == false {
			log.Printf("Process 1 waiting\n")
			controller.OutsideCriticalSection(1, 100)
		}

		controller.EnterCriticalSection(1)
		controller.InsideCriticalSection(1, 50)
		controller.LeaveCriticalSection(1)

		c1 = false

		controller.OutsideCriticalSection(1, 100)

		c1 = true

		if controller.ProcessCrashed(0.1) {
			log.Printf("Process 1 crashed\n")
			return
		}

	}
}

// process2 simulates the behaviour of the second process
func process2() {
	for {

		for c1 == false {
			log.Printf("Process 1 waiting\n")
			controller.OutsideCriticalSection(1, 100)
		}

		c2 = false

		controller.EnterCriticalSection(2)
		controller.InsideCriticalSection(2, 50)
		controller.LeaveCriticalSection(2)

		c2 = true

		controller.OutsideCriticalSection(2, 100)

		if controller.ProcessCrashed(0.1) {
			log.Printf("Process 2 crashed\n")
			return
		}

	}
}
