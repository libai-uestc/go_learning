package server

import (
	transport "libai/go/basic/phase-one/socket"
	"log"
	"net"
	"time"
)

func UdpServer() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5678")
	transport.CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr) // UDP不需要创建连接，所以不需要像TCP那样通过Accept()创建连接，这里的conn是个假连接，不需要阻塞
	transport.CheckError(err)
	log.Printf("return conn")
	conn.SetReadDeadline(time.Now().Add(30 * time.Second)) // 超时到来之前，client必须发来数据
}
