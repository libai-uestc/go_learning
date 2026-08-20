package io_test

import (
	"libai/go/basic/phase-one/io"
	"testing"
)

func TestRegex(t *testing.T) {
	io.UseRegex()
}

// go test -v ./phase-one/io -run=^TestRegex$ -count=1
