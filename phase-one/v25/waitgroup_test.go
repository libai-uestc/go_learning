package v25_test

import (
	v25 "libai/go/basic/phase-one/v25"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	v25.Sum()
}

// go test -v ./phase-one/v25 -run=^TestWaitGroup$ -count=1
