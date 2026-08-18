package io_test

import (
	"libai/go/basic/phase-one/io"
	"testing"
)

func TestCompress(t *testing.T) {
	io.Compress("../img/迈克尔乔丹.png", "../img/迈克尔乔丹.png.gzip", io.GZIP)
	io.Decompress("../img/迈克尔乔丹.png.gzip", "../data/迈克尔乔丹.png", io.GZIP)
}

// go test -v ./phase-one/io -run=^TestCompress$ -count=1
