package main

import "libai/go/basic/phase-one/concurrence"

func main() {
	// concurrence.SimpleGoroutine()
	// concurrence.SubRoutine()
	// concurrence.WaitGroup()
	// concurrence.Atomic()
	// concurrence.Lock()
	// concurrence.ReentranceRLock(3)
	concurrence.ReentranceWLock(3)
}

// go run .\phase-one\concurrence\main\main.go
