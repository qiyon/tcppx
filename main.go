package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"tcppx/servive"
	"time"
)

var (
	local  = flag.String("l", "", "Local Address, example: 0.0.0.0:18888")
	remote = flag.String("r", "", "Remote Address, example: 127.0.0.1:18889")
)

func main() {
	flag.Parse()
	if *local == "" || *remote == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	localAddr, err := net.ResolveTCPAddr("tcp", *local)
	if err != nil {
		fmt.Println("Err local addr: ", err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}
	remoteAddr, err := net.ResolveTCPAddr("tcp", *remote)
	if err != nil {
		fmt.Println("Err remote addr: ", err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Tcp Proxy from %s to %s\n", *local, *remote)

	listener, err := net.ListenTCP("tcp", localAddr)
	if err != nil {
		fmt.Println("Listen err: ", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept connection err: ", err.Error())
			continue
		}
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] Connection from: %s\n", timeStr, conn.RemoteAddr().String())
		px := servive.NewProxy(conn, localAddr, remoteAddr)
		go px.Run()
	}
}
