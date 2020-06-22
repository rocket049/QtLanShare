package main

import (
	"time"

	"gitee.com/rocket049/discover-go"
)

func runDiscover(ips []string) {
	server := discover.NewServer()
	for _, ip := range ips {
		server.Append("http", ip, 6868, "index", "FileServer", "Share Files")
	}

	go server.Serve(true)
	time.Sleep(time.Millisecond * 100)
}
