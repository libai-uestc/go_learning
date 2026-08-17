package io_test

import (
	"libai/go/basic/phase-one/io"
	"testing"
)

func TestCreateFile(t *testing.T) {
	io.CreateFile("../data/poem.txt")
}
