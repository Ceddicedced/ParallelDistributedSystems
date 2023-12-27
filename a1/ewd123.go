/*
 * ewd123.go
 *
 * A program to select one of the mutex strategies, as described in EWD123.
 *
 * Copyright (c) 2019-2022 HS Emden/Leer
 * All Rights Reserved.
 *
 * version 2.00 - 30 Oct 2022 - GJV - transform into workspace version
 * version 1.00 - 21 Oct 2019 - GJV - initial version
 *
 * author: Gert Veltink, gert.veltink@hs-emden-leer.de (GJV)
 */

package main

import (
	"fmt"
	"log"
	"os"

	"busacker.dev/a1/ewd123a"
	"busacker.dev/a1/ewd123b"
	"busacker.dev/a1/ewd123c"
	"busacker.dev/a1/ewd123d"
	"busacker.dev/a1/ewd123dekker"
)

// main is the entry point for execution.
func main() {
	var quitChannel = make(chan string) // channel to send a quit signal

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile) // set specific logging attributes

	var variant = ""
	if len(os.Args[1:]) > 0 { // do we have arguments?
		variant = os.Args[1]
	}

	log.Printf("*** Start %s ***\n", variant)

	switch variant {
	case "EWD123a":
		ewd123a.Start()
	case "EWD123b":
		ewd123b.Start()
	case "EWD123c":
		ewd123c.Start()
	case "EWD123d":
		ewd123d.Start()
	case "Dekker":
		ewd123dekker.Start()
	default:
		log.Printf("unknown variant: '%s'!", variant)
		fmt.Printf("Usage: %s <variant>, where variant is: ", os.Args[0])
		fmt.Printf("EWD123a, EWD123b, EWD123c, EWD123d, Dekker or empty for the Dekker solution\n")
		os.Exit(-1) // signal error code -1
	}

	log.Printf("*** Finish: %s\n", <-quitChannel) // wait for a quit signal
}
