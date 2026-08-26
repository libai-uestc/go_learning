package client_test

import (
	"libai/go/basic/phase-one/socket/client"
	"testing"
)

func TestTcpClient(t *testing.T) {
	client.TcpClient()
}

func TestUdpClient(t *testing.T) {
	client.UdpClient()
}

func TestTcpLongConnection(t *testing.T) {
	client.TcpLongConnection()
}

func TestUdpLongConnection(t *testing.T) {
	client.UdpLongConnection()
}

func TestTcpStick(t *testing.T) {
	client.TcpStick()
}

func TestUdpConnectionCurrent(t *testing.T) {
	client.UdpConnectionCurrent()
}

func TestUdpRpcClient(t *testing.T) {
	client.UdpRpcClient()
}

//  go test -v .\phase-one\socket\client\ -run=^TestTcpClient$ -count=1
