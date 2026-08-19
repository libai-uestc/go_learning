package io_test

import (
	"libai/go/basic/phase-one/io"
	"testing"
)

func TestLog(t *testing.T) {
	logger := io.NewLogger("../data/biz.log")
	io.Log(logger)
}

func TestSLog(t *testing.T) {
	logger := io.NewSLogger("../data/sbiz.log")
	io.SLog(logger)
}
