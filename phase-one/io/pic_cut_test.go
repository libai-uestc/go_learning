package io_test

import (
	"libai/go/basic/phase-one/io"
	"testing"
)

func TestSplitFile(t *testing.T) {
	imgFile := "../img/迈克尔乔丹.png"
	io.SplitFile(imgFile, "../img/图像分割", 4)
}

func TestMergeFile(t *testing.T) {
	io.MergeFile("../img/图像分割", "../img/图像合并.png")
}

// go test -v ./phase-one/io -run=^TestSplitFile$ -count=1
// go test -v ./phase-one/io -run=^TestMergeFile$ -count=1
