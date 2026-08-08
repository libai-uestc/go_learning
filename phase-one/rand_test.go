package main_test

import (
	"fmt"
	"math/rand"
	randv2 "math/rand/v2"
	"testing"
)

const MAX = 1e9

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(MAX)
	}
}

func BenchmarkRandV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randv2.IntN(MAX)
	}
}

func TestRand(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d  ", randv2.IntN(100))
	}
}

func TestRandSeed(t *testing.T) {
	source := randv2.NewPCG(123, 456)
	for i := 0; i < 5; i++ {
		// source.Seed(123,456)
		rander := randv2.New(source)
		fmt.Printf("%d  ", rander.IntN(100))
	}
}

// go test -v .\phase-one\ -run=^TestRand$ -count=1
// go test -v ./phase-one -run=^TestRandSeed$ -count=1
// go test ./phase-one -bench=Rand -run=^$ -count=1
/*
BenchmarkRand-32        154673538                7.658 ns/op
BenchmarkRandV2-32      263124138                4.848 ns/op
*/
