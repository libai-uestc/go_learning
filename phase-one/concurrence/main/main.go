package main

import "libai/go/basic/phase-one/concurrence"

func main() {
	// concurrence.SimpleGoroutine()
	// concurrence.SubRoutine()
	// concurrence.WaitGroup()
	// concurrence.Atomic()
	// concurrence.Lock()
	// concurrence.ReentranceRLock(3)
	// concurrence.ReentranceWLock(3)
	// concurrence.WLockExclusion()
	// concurrence.RLockExclusion()
	// concurrence.LockQueue()
	// concurrence.ReadWriteRace()
	// concurrence.CollectionSafety()
	// concurrence.ServiceMain()
	// concurrence.TraverseChannel()
	// concurrence.Block()
	// concurrence.Broadcast()
	// concurrence.CondSignal()
	// concurrence.ChannelSignal()
	// concurrence.CondBroadcast()
	// concurrence.ChannelBroadcast()
	concurrence.DealMassFile("data/biz_log")
	concurrence.QpsLimit()
}

// go run .\phase-one\concurrence\main\main.go
