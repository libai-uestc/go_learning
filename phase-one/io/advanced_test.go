package io_test

import (
	"libai/go/basic/phase-one/io"
	"testing"
)

// 测试 LimitReader
func TestLimitReader(t *testing.T) {
	io.LimitReader()
}

func TestMultiReader(t *testing.T) {
	io.MultiReader()
}

func TestMultiWriter(t *testing.T) {
	io.MultiWriter()
}

func TestTeeReader(t *testing.T) {
	io.TeeReader()
}

func TestPipeIO(t *testing.T) {
	io.PipeIO()
}

// go test -v ./phase-one/io -run=^TestLimitReader$ -count=1
// go test -v ./phase-one/io -run=^TestMultiReader$ -count=1
// go test -v ./phase-one/io -run=^TestMultiWriter$ -count=1
// go test -v ./phase-one/io -run=^TestTeeReader$ -count=1
// go test -v ./phase-one/io -run=^TestPipeIO$ -count=1
