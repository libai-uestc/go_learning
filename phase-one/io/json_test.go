package io_test

import (
	"libai/go/basic/phase-one/io"
	"testing"
)

func TestJson(t *testing.T) {
	io.JsonSerialize()
}

// go test -v ./phase-one/io -run=^TestJson$ -count=1
// 序列化成功 {"Name":"李白","Birthday":"2026-08-19T11:59:05.4875843+08:00","CreatedAt":"2026-08-19","gender":1,"Address":{"Province":"浙江","City":"杭州"}}
// 反序列化成功 {Name:李白 Age:0 height:0 Birthday:2026-08-19 11:59:05.4875843 +0800 CST CreatedAt:2026-08-19 Sex:1 Address:{Province:浙江 City:杭州}}
