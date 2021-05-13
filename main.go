package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var (
	remoteIP   string
	remotePort string
)

func rawConnect(remoteIP string, remotePort string) {
	logger := logger()
	tStatrt := time.Now()
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(remoteIP, remotePort), timeout)
	tEnd := time.Now()
	duration := tEnd.Sub(tStatrt)
	if err != nil {
		logger.Fatalf("dstAddr: %s:%s, Status: %s, time: %s", remoteIP, remotePort, "connection refused", duration)
	}
	if conn != nil {
		defer conn.Close()
		logger.Printf("dstAddr: %s, Status: %s, time: %s", conn.RemoteAddr().String(), "ok", duration)
	}
}

func logger() (l *log.Logger) {
	writer, err := os.OpenFile("check.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("create file check.log failed: %v", err)
	}
	logger := log.New(io.MultiWriter(writer), "", log.Lshortfile|log.LstdFlags)
	return logger
}

func init() {
	flag.StringVar(&remoteIP, "remoteIP", "", "remote ip")
	flag.StringVar(&remotePort, "remotePort", "", "remote port")
}

func main() {
	flag.Parse()
	rawConnect(remoteIP, remotePort)
}
