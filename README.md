# Repository for Parallel and Distributed Systems Assignments

This repository contains my solutions for the assignments in the course "Parallel and Distributed Systems." The course covered various aspects of mutual exclusion in Go, specifications in mCRL2, and concurrent algorithms in Go.

## Assignment 1: Mutual Exclusion in Go

This section includes implementations for the mutual exclusion problem using Go. It covers five solutions:
- `ewd123a`
- `ewd123b`
- `ewd123c`
- `ewd123d`
- `ewd123dekker`

Each solution demonstrates different aspects and challenges in achieving synchronization in critical sections. The lab reports provide detailed analyses of their behaviors and their theoretical underpinnings.

## Assignment 2: Specifications in mCRL2

### 2a: Vending Machine - Sequential Specifications

This part of the assignment involves extending a vending machine specification in mCRL2. The task includes defining equations for coin values, product prices, and the main vending machine process. The specification is both unbounded and bounded (limited credit acceptance).

### 2b: Traffic Lights - Parallel Specifications

This section develops a specification for a traffic light system at a four-way crossing. The task progresses through four stages, from independent light control to a distributed, synchronized system preventing unsafe states.

## Assignment 3: Concurrent Algorithms in Go

The final assignment is an implementation of the mCRL2 traffic light system specification using Go. This involves creating concurrent functions with goroutines and channels, emphasizing synchronization without shared memory or global variables.

## Repository Structure

- `a1/`: Contains Go files for each mutual exclusion solution and their respective lab reports.
- `a2/`: 
  - `VendingMachine/`: mCRL2 specifications for the vending machine.
  - `TrafficLights/`: mCRL2 specifications for the traffic light system.
- `a3/`: Go implementation of the traffic light system using concurrent programming techniques.
- `README.md`: This file, providing an overview of the repository contents.

