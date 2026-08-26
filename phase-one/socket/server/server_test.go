package server_test

import (
	"libai/go/basic/phase-one/socket/server"
	"testing"
)

func TestTcpServer(t *testing.T) {
	server.TcpServer()
}

func TestUdpServer(t *testing.T) {
	server.UdpServer()
}

func TestTcpLongConnection(t *testing.T) {
	server.TcpLongConnection()
}

func TestUdpLongConnection(t *testing.T) {
	server.UdpLongConnection()
}

func TestTcpStick(t *testing.T) {
	server.TcpStick()
}

func TestUdpConnectionCurrent(t *testing.T) {
	server.UdpConnectionCurrent()
}

func TestUdpRpcServer(t *testing.T) {
	server.UdpRpcServer()
}

// go test -v .\phase-one\socket\server\ -run=^TestTcpServer$ -count=1
